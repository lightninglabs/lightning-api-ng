package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/encoding/protojson"
	"gopkg.in/yaml.v3"
)

var (
	// subCommandPattern is the pattern that matches against the CLI sub
	// command indication in a gRPC method's comment.
	subCommandPattern = regexp.MustCompile("(.*?): `(.*?)`")
)

func main() {
	app := os.Args[1]
	repoURL := os.Args[2]
	protoSrcDir := os.Args[3]
	srcCommit := os.Args[4]
	mainFile := fmt.Sprintf("./build/protos/%s/generated.json", app)
	template := &Template{}
	fmt.Printf("Reading template file %s\n", mainFile)
	tplBytes, err := os.ReadFile(mainFile)
	if err != nil {
		fail(err)
	}

	err = json.Unmarshal(tplBytes, &template)
	if err != nil {
		fail(err)
	}

	fmt.Printf("Got template with %d files\n", len(template.Files))
	template.RESTTypes = make(map[string]interface{})

	// If the proto source dir is not empty, make sure we can use
	// it directly by appending a tailing path separator.
	if protoSrcDir != "" {
		protoSrcDir = fmt.Sprintf("%s/", protoSrcDir)
	}

	for idx, file := range template.Files {
		if len(template.Files[idx].Services) == 0 {
			continue
		}

		baseName := strings.ReplaceAll(file.Name, ".proto", "")

		protoFile := fmt.Sprintf(
			"./build/%s/%s%s", app, protoSrcDir, file.Name,
		)
		externalLink := fmt.Sprintf(
			"%s/blob/%s/%s%s", repoURL, srcCommit, protoSrcDir,
			file.Name,
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

				err := parseMethodDescription(method)
				if err != nil {
					fail(err)
				}
			}
		}

		restFile := fmt.Sprintf(
			"./build/protos/%s/%s.swagger.json", app, baseName,
		)

		fmt.Printf("Reading REST file %s\n", restFile)
		restBytes, err := os.ReadFile(restFile)
		if err != nil {
			fmt.Printf("Skipping missing REST file\n")
			continue
		}

		rest := &Swagger{}
		err = json.Unmarshal(restBytes, &rest)
		if err != nil {
			fail(err)
		}

		for name, definition := range rest.Definitions {
			template.RESTTypes[name] = definition
		}

		mappingFile := fmt.Sprintf(
			"./build/protos/%s/%s.yaml", app, baseName,
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
			service := template.Files[idx].Services[0]
			ruleMethod := strings.ReplaceAll(
				rule.GetSelector(), service.FullName+".", "",
			)

			for idx := range service.Methods {
				method := service.Methods[idx]
				if method.Name == ruleMethod {
					mapping := NewRESTMapping(
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

	finalFile := fmt.Sprintf("./build/%s.json", app)
	finalBytes, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		fail(err)
	}

	err = os.WriteFile(finalFile, finalBytes, 0o644)
	if err != nil {
		fail(err)
	}
	fmt.Printf("Saved output to: %s\n", finalFile)
}

func fail(err error) {
	fmt.Printf("Error: %v\n", err)
	os.Exit(1)
}

func assignSourceLinks(method *ServiceMethod, source, baseLink string) {
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

func parseMethodDescription(method *ServiceMethod) error {
	description := method.Description
	if subCommandPattern.MatchString(description) {
		matches := subCommandPattern.FindStringSubmatch(description)

		// At index 0 we have the full matched string, capture groups
		// start at index 1.
		method.CommandLine = fmt.Sprintf(
			"%s %s", matches[1], matches[2],
		)

		// Run the command and capture its output.
		args := append(strings.Split(matches[2], " "), "--help")
		cmd := exec.Command(matches[1], args...)
		out, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("error invoking %s: %w",
				method.CommandLine, err)
		}

		method.CommandLineHelp = string(out)
	}

	return nil
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
