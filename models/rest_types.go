package models

import (
	"fmt"
	"os"
	"strings"

	"github.com/lightninglabs/lightning-api-ng/defs"
)

type RestTypes struct {
	// FieldTypes is a mapping from a message's full field name to the REST
	// type.
	// ex: frdrpcCloseRecommendationsResponse.considered_channels -> integer
	FieldTypes map[string]string
	// EnumTypes is a mapping from an enum's name to the REST type
	// ex: ForceClosedChannelAnchorState -> string
	EnumTypes map[string]string
}

// NewRestTypes creates a new RestTypes from a map of RESTType definitions.
func NewRestTypes(restTypesDef map[string]*defs.RESTType) *RestTypes {
	fmt.Printf("Creating %d RestTypes\n", len(restTypesDef))
	r := &RestTypes{}
	r.FieldTypes = make(map[string]string)
	r.EnumTypes = make(map[string]string)

	// First parse all enums so we can reference them when parsing objects
	for name, restType := range restTypesDef {
		if restType.Type != "object" {
			r.parseEnum(name, restType)
		}
	}

	// Next, parse all objects.
	for name, restType := range restTypesDef {
		if restType.Type == "object" {
			r.parseObject(name, restType)
		}
	}
	return r
}

// parseEnum adds a RESTType definition to the EnumTypes map
func (r *RestTypes) parseEnum(typeName string, restType *defs.RESTType) {
	r.EnumTypes[typeName] = restType.Type
}

// parseObject parses a RESTType definition and adds it to the FieldTypes map
func (r *RestTypes) parseObject(typeName string, restType *defs.RESTType) {
	// skip if there are no properties defined
	if restType.Properties == nil {
		return
	}
	for fieldName, details := range restType.Properties {
		fullName := typeName + "." + fieldName
		r.FieldTypes[fullName] = details.Type
		fieldType := "unknown"
		if details.Ref != "" {
			// $ref looks like: '#/definitions/InvoiceInvoiceState'
			// check if it is an enum or fallback to 'object'
			index := strings.LastIndex(details.Ref, "/") + 1
			refName := details.Ref[index:]

			if enumType, ok := r.EnumTypes[refName]; ok {
				fieldType = enumType
			} else {
				fieldType = "object"
			}
		} else if details.Type != "" {
			fieldType = details.Type
		} else {
			fmt.Printf("Could not determine the field type "+
				"for %s\n", fullName)
			os.Exit(1)
		}
		r.FieldTypes[fullName] = fieldType
	}
}

// UpdateMessage updates a Message's fields with the REST type
func (r *RestTypes) UpdateMessage(msg *Message) {
	msgRestType := convertGrpcTypeToRest(msg.FullName)

	// update REST type and placement from restTypes
	for _, field := range msg.Fields {
		restTypeName := msgRestType + "." + field.Name
		if restType, ok := r.FieldTypes[restTypeName]; ok {
			field.RestType = restType
		}
	}
}

// convertGrpcTypeToRest converts a gRPC type to a REST type
// ex: lnrpc.AddInvoiceResponse -> lnrpcAddInvoiceResponse
func convertGrpcTypeToRest(typeName string) string {
	parts := strings.Split(typeName, ".")
	return parts[0] + parts[len(parts)-1]
}
