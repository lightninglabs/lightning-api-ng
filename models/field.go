package models

import (
	"fmt"
	"strings"

	"github.com/lightninglabs/lightning-api-ng/markdown"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
)

type Field struct {
	Message *Message

	Name        string
	Description string
	Label       string
	Type        string
	FullType    string

	// These values are set by the REST mapping
	RestType      string
	RestPlacement string
}

// NewField creates a new Field model from a MessageField definition
func NewField(fieldDef gendoc.MessageField, message *Message) *Field {
	return &Field{
		Message: message,
		Name:    fieldDef.Name,
		Description: markdown.CleanDescription(
			fieldDef.Description, true),
		Label:    fieldDef.Label,
		Type:     fieldDef.Type,
		FullType: fieldDef.FullType,

		RestType:      "unknown",
		RestPlacement: "unknown",
	}
}

// GrpcType returns the type of the field with a special case for arrays
func (f *Field) GrpcType() string {
	if f.Label == "repeated" {
		return fmt.Sprintf("%s[]", f.Type)
	}
	return f.Type
}

// LinkedType returns the field type in markdown link format if the type
// is not a scalar
func (f *Field) LinkedType() string {
	if strings.Contains(f.FullType, ".") {
		hash := strings.ReplaceAll(strings.ToLower(f.FullType), ".", "")
		return fmt.Sprintf("[`%s`](#%s)", f.GrpcType(), hash)
	} else {
		return fmt.Sprintf("`%s`", f.GrpcType())
	}
}

// Encoding returns the encoding of the field if it is a byte array depending
// on the REST placement
func (f *Field) Encoding() string {
	if f.Type == "bytes" {
		switch f.RestPlacement {
		case "query":
			return "base64 & URL"
		case "body":
			return "base64"
		}
	}
	return ""
}

// EncodingTip returns a tip for the encoding of the field if it is a byte array
func (f *Field) EncodingTip() string {
	if f.Encoding() != "" {
		return strings.Join([]string{
			"<Tip>",
			fmt.Sprintf("Use %s encoding in the %s.", f.Encoding(),
				f.RestPlacement),
			"<br />",
			fmt.Sprintf("See [REST Encoding]"+
				"(/api/%s/#rest-encoding).",
				f.Message.Package.App.Name),
			"</Tip>",
		}, "")
	}
	return ""
}
