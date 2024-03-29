package defs

import (
	gendoc "github.com/pseudomuto/protoc-gen-doc"
)

// ApiSpec is a type for encapsulating all the parsed files, messages, fields,
// enums, services, extensions, etc. into an object that will be supplied to
// the markdown exporter.
type ApiSpec struct {
	// The files that were parsed
	Files []*File `json:"files"`

	// Details about the scalar values and their respective types in
	// supported languages.
	Scalars []*ScalarValue `json:"scalarValueTypes"`

	RESTTypes map[string]*RESTType `json:"restTypes"`
}

// File wraps all the relevant parsed info about a proto file. File objects
// guarantee that their top-level enums, extensions, messages, and services are
// sorted alphabetically based on their "long name". Other values (enum values,
// fields, service methods) will be in the order that they're defined within
// their respective proto files.
//
// In the case of proto3 files, HasExtensions will always be false, and
// Extensions will be empty.
type File struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Package     string `json:"package"`

	HasEnums      bool `json:"hasEnums"`
	HasExtensions bool `json:"hasExtensions"`
	HasMessages   bool `json:"hasMessages"`
	HasServices   bool `json:"hasServices"`

	Enums      []*gendoc.Enum          `json:"enums"`
	Extensions []*gendoc.FileExtension `json:"extensions"`
	Messages   []*gendoc.Message       `json:"messages"`
	Services   []*Service              `json:"services"`

	Options map[string]interface{} `json:"options,omitempty"`
}

// Service contains details about a service definition within a proto file.
type Service struct {
	Name        string           `json:"name"`
	LongName    string           `json:"longName"`
	FullName    string           `json:"fullName"`
	Description string           `json:"description"`
	Methods     []*ServiceMethod `json:"methods"`

	Options map[string]interface{} `json:"options,omitempty"`
}

// ServiceMethod contains details about an individual method within a service.
type ServiceMethod struct {
	Name               string         `json:"name"`
	Description        string         `json:"description"`
	Source             string         `json:"source"`
	CommandLine        string         `json:"commandLine"`
	CommandLineHelp    string         `json:"commandLineHelp"`
	RequestType        string         `json:"requestType"`
	RequestLongType    string         `json:"requestLongType"`
	RequestFullType    string         `json:"requestFullType"`
	RequestTypeSource  string         `json:"requestTypeSource"`
	RequestStreaming   bool           `json:"requestStreaming"`
	ResponseType       string         `json:"responseType"`
	ResponseLongType   string         `json:"responseLongType"`
	ResponseFullType   string         `json:"responseFullType"`
	ResponseTypeSource string         `json:"responseTypeSource"`
	ResponseStreaming  bool           `json:"responseStreaming"`
	RESTMappings       []*RESTMapping `json:"restMappings"`

	Options map[string]interface{} `json:"options,omitempty"`
}

// ScalarValue contains information about scalar value types in protobuf. The
// common use case for this type is to know which language specific type maps to
// the protobuf type.
//
// For example, the protobuf type `int64` maps to `long` in C#, and `Bignum` in
// Ruby. For the full list, take a look at
// https://developers.google.com/protocol-buffers/docs/proto3#scalar
type ScalarValue struct {
	ProtoType  string `json:"protoType"`
	Notes      string `json:"notes"`
	CppType    string `json:"cppType"`
	CSharp     string `json:"csType"`
	GoType     string `json:"goType"`
	JavaType   string `json:"javaType"`
	PhpType    string `json:"phpType"`
	PythonType string `json:"pythonType"`
	RubyType   string `json:"rubyType"`
}

type Swagger struct {
	Paths       map[string]map[string]*RESTDetails `json:"paths"`
	Definitions map[string]*RESTType               `json:"definitions"`
}
