package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/encoding/protojson"
	"gopkg.in/yaml.v3"
)

func main() {
	app := os.Args[1]
	mainFile := fmt.Sprintf("./build/protos/%s/generated.json", app)
	template := &Template{}
	fmt.Printf("Reading template file %s\n", mainFile)
	tplBytes, err := ioutil.ReadFile(mainFile)
	if err != nil {
		fail(err)
	}

	err = json.Unmarshal(tplBytes, &template)
	if err != nil {
		fail(err)
	}

	fmt.Printf("Got template with %d files\n", len(template.Files))
	template.RESTTypes = make(map[string]interface{})

	for idx, file := range template.Files {
		if len(template.Files[idx].Services) == 0 {
			continue
		}

		baseName := strings.ReplaceAll(file.Name, ".proto", "")

		restFile := fmt.Sprintf(
			"./build/protos/%s/%s.swagger.json", app, baseName,
		)

		fmt.Printf("Reading REST file %s\n", restFile)
		restBytes, err := ioutil.ReadFile(restFile)
		if err != nil {
			fail(err)
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
		mappingBytes, err := ioutil.ReadFile(mappingFile)
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

	err = ioutil.WriteFile(finalFile, finalBytes, 0o644)
	if err != nil {
		fail(err)
	}
}

func fail(err error) {
	fmt.Printf("Error: %v\n", err)
	os.Exit(1)
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
