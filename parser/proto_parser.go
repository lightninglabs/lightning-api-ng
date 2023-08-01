package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/lightninglabs/lightning-api-ng/config"
	"github.com/lightninglabs/lightning-api-ng/defs"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/encoding/protojson"
	"gopkg.in/yaml.v3"
)

var (
	// subCommandPattern is the pattern that matches against the CLI sub
	// command indication in a gRPC method's comment.
	subCommandPattern = regexp.MustCompile("(.*?): `(.*?)`")
)

func LoadApiSpec(config *config.Config) *defs.ApiSpec {
	mainFile := fmt.Sprintf("./build/protos/%s/generated.json", config.App)
	apiSpec := &defs.ApiSpec{}
	fmt.Printf("Reading API specification file %s\n", mainFile)
	tplBytes, err := os.ReadFile(mainFile)
	if err != nil {
		fail(err)
	}

	err = json.Unmarshal(tplBytes, &apiSpec)
	if err != nil {
		fail(err)
	}

	fmt.Printf("Got API spec with %d files\n", len(apiSpec.Files))
	apiSpec.RESTTypes = make(map[string]*defs.RESTType)

	for idx, file := range apiSpec.Files {
		if len(apiSpec.Files[idx].Services) == 0 {
			continue
		}

		baseName := strings.ReplaceAll(file.Name, ".proto", "")

		protoFile := fmt.Sprintf(
			"./build/%s/%s%s", config.App, config.ProtoSrcDir,
			file.Name,
		)
		externalLink := fmt.Sprintf(
			"%s/blob/%s/%s%s", config.RepoURL, config.Commit,
			config.ProtoSrcDir, file.Name,
		)
		fmt.Printf("Reading proto file %s with external link %s\n",
			protoFile, externalLink)
		protoSourceBytes, err := os.ReadFile(protoFile)
		if err != nil {
			fail(err)
		}

		fmt.Printf("Generating source file lookup links and CLI help " +
			"texts\n")
		for sIdx := range file.Services {
			for mIdx := range file.Services[sIdx].Methods {
				method := file.Services[sIdx].Methods[mIdx]
				assignSourceLinks(
					method, string(protoSourceBytes),
					externalLink,
				)

				parseMethodDescription(method)
			}
		}

		restFile := fmt.Sprintf(
			"./build/protos/%s/%s.swagger.json", config.App,
			baseName,
		)

		fmt.Printf("Reading REST file %s\n", restFile)
		restBytes, err := os.ReadFile(restFile)
		if err != nil {
			fmt.Printf("Skipping missing REST file\n")
			continue
		}

		rest := &defs.Swagger{}
		err = json.Unmarshal(restBytes, &rest)
		if err != nil {
			fail(err)
		}

		for name, definition := range rest.Definitions {
			apiSpec.RESTTypes[name] = definition
		}

		mappingFile := fmt.Sprintf(
			"./build/protos/%s/%s.yaml", config.App, baseName,
		)

		fmt.Printf("Reading http mapping file %s\n", mappingFile)
		mappingBytes, err := os.ReadFile(mappingFile)
		if err != nil {
			fail(err)
		}

		http, err := loadGrpcAPIServiceFromYAML(
			mappingBytes, mappingFile,
		)
		if err != nil {
			fail(err)
		}

		rules := http.Rules
		fmt.Printf("Got mapping with %d rules\n", len(rules))
		for _, rule := range rules {
			service := apiSpec.Files[idx].Services[0]
			ruleMethod := strings.ReplaceAll(
				rule.GetSelector(), service.FullName+".", "",
			)

			for idx := range service.Methods {
				method := service.Methods[idx]
				if method.Name == ruleMethod {
					mapping := defs.NewRESTMapping(
						rule, rest.Paths,
					)

					method.RESTMappings = append(
						method.RESTMappings,
						mapping,
					)

					break
				}
			}
		}
	}

	return apiSpec
}

func fail(err error) {
	fmt.Printf("Error: %v\n", err)
	os.Exit(1)
}

func assignSourceLinks(method *defs.ServiceMethod, source, baseLink string) {
	lines := strings.Split(strings.ReplaceAll(source, "\r\n", "\n"), "\n")

	lineSuffix := func(needle string) string {
		for idx, line := range lines {
			if strings.Contains(line, needle) {
				// GitHub starts counting at line 1...
				return fmt.Sprintf("#L%d", idx+1)
			}
		}

		fmt.Printf("WARN: Source for needle '%s' not found in %s\n",
			needle, baseLink)
		return ""
	}

	// Find the method definition in the proto file.
	suffix := lineSuffix(fmt.Sprintf("rpc %s", method.Name))
	method.Source = fmt.Sprintf("%s%s", baseLink, suffix)

	// Find the request type in the proto file.
	suffix = lineSuffix(fmt.Sprintf("message %s", method.RequestType))
	method.RequestTypeSource = fmt.Sprintf("%s%s", baseLink, suffix)

	// Find the response type in the proto file.
	suffix = lineSuffix(fmt.Sprintf("message %s", method.ResponseType))
	method.ResponseTypeSource = fmt.Sprintf("%s%s", baseLink, suffix)
}

func parseMethodDescription(method *defs.ServiceMethod) {
	description := method.Description
	if subCommandPattern.MatchString(description) {
		matches := subCommandPattern.FindStringSubmatch(description)

		// Run the command and capture its output.
		args := append(strings.Split(matches[2], " "), "--help")
		cmd := exec.Command(matches[1], args...)
		out, err := cmd.Output()
		if err != nil {
			fmt.Printf("error invoking %s: %s\n",
				method.CommandLine, err.Error())
			return
		}

		// At index 0 we have the full matched string, capture groups
		// start at index 1.
		method.CommandLine = fmt.Sprintf(
			"%s %s", matches[1], matches[2],
		)

		method.CommandLineHelp = string(out)
	}
}

func loadGrpcAPIServiceFromYAML(yamlFileContents []byte,
	yamlSourceLogName string) (*annotations.Http, error) {

	var yamlContents interface{}
	err := yaml.Unmarshal(yamlFileContents, &yamlContents)
	if err != nil {
		return nil, fmt.Errorf("failed to parse gRPC API "+
			"Configuration from YAML in '%v': %v",
			yamlSourceLogName, err)
	}

	parsedHttp := yamlContents.(map[string]interface{})["http"]

	jsonContents, err := json.Marshal(parsedHttp)
	if err != nil {
		return nil, err
	}

	// As our GrpcAPIService is incomplete, accept unknown fields.
	unmarshaler := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}

	http := &annotations.Http{}
	if err := unmarshaler.Unmarshal(jsonContents, http); err != nil {
		return nil, fmt.Errorf("failed to parse gRPC API "+
			"Configuration from YAML in '%v': %v",
			yamlSourceLogName, err)
	}

	return http, nil
}
