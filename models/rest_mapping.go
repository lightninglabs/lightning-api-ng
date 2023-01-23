package models

import (
	"fmt"
	"strings"

	"github.com/lightninglabs/lightning-api-ng/defs"
)

type RestMapping struct {
	Method string
	Path   string
	// Parameters maps from parameter name to its info. Keys could be:
	// 'sat_per_byte' or 'channel_point.funding_txid_str'.
	Parameters map[string]*RestParameter
	// HasBodyParams indicates if there is a generic `body` parameter which
	// holds the full request message.
	HasBodyParams bool
}

type RestParameter struct {
	Name   string
	Format string
	Type   string
	In     string
}

// NewRestMapping creates a new RestMapping from a RESTMapping definition.
func NewRestMapping(mappingDef defs.RESTMapping) *RestMapping {
	mapping := &RestMapping{}
	mapping.Method = mappingDef.Method
	mapping.Path = mappingDef.Path

	mapping.Parameters = make(map[string]*RestParameter)
	if mappingDef.Details != nil && mappingDef.Details.Parameters != nil {
		for _, paramDef := range mappingDef.Details.Parameters {
			if mapping.Method == "POST" &&
				paramDef.Name == "body" &&
				paramDef.Schema != nil &&
				paramDef.Schema.Ref != "" {

				// Indicate that all of the parameters are in
				// the body.
				mapping.HasBodyParams = true
			} else {
				param := &RestParameter{
					Name:   paramDef.Name,
					Format: paramDef.Format,
					Type:   paramDef.Type,
					In:     paramDef.In,
				}
				mapping.Parameters[param.Name] = param
			}
		}
	}
	return mapping
}

// UpdateMessage updates the REST type and placement of a message's fields.
func (r *RestMapping) UpdateMessage(msg *Message) {
	// Update REST type and placement from restMappings.
	for _, field := range msg.Fields {
		// skip fields that have already been set from the restTypes
		// parsing in `rest-types.go`.
		if field.RestType != "unknown" &&
			field.RestPlacement != "unknown" {

			continue
		}

		if strings.Contains(field.FullType, ".") {
			r.updateStructField(field)
		} else {
			r.updateScalarField(field)
		}

		// If no parameter was found to update the placement, this is
		// because the restMapping only has a `body` parameter for
		// POST requests.
		if r.HasBodyParams && field.RestPlacement == "unknown" {
			field.RestPlacement = "body"
		}
	}
}

// updateStructField updates the REST type and placement of a struct field.
func (r *RestMapping) updateStructField(field *Field) {
	// Find the parameters that match the field name.
	params := r.getParamsByPrefix(field.Name)

	// The REST type may have been set already from the rest_types parsing.
	if field.RestType == "unknown" {
		field.RestType = "object"
	}
	field.RestPlacement = getCombinedPlacement(params)

	if len(params) == 0 {
		return
	}

	// Update the field's description to include details about the placement
	// of each param.
	details := make([]string, len(params))
	for i, param := range params {
		details[i] = fmt.Sprintf("`%s`: `%s` in `%s`", param.Name,
			param.Type, param.In)
	}
	field.Description += "<h4>Nested REST Parameters</h4>"
	field.Description += fmt.Sprintf("<p>%s</p>",
		strings.Join(details, "<br />"))
}

// updateScalarField updates the REST type and placement of a scalar field.
func (r *RestMapping) updateScalarField(field *Field) {
	// Find the parameter that matches the field name.
	param := r.Parameters[field.Name]
	if param == nil {
		return
	}

	// Update the field with the parameter info.
	field.RestType = param.Type
	field.RestPlacement = param.In
}

// GetParamsByPrefix returns all parameters that have the given prefix. If
// the prefix is empty, all parameters are returned.
func (r *RestMapping) getParamsByPrefix(prefix string) []*RestParameter {
	var params []*RestParameter
	for _, param := range r.Parameters {
		if param.Name == prefix ||
			strings.HasPrefix(param.Name, prefix+".") {

			params = append(params, param)
		}
	}
	return params
}

// getCombinedPlacement returns the placement of the parameters. If all
// parameters are in the same placement, that placement is returned. If
// parameters are in different placements, "mixed" is returned. If there are
// no parameters, "unknown" is returned.
func getCombinedPlacement(params []*RestParameter) string {
	if len(params) == 0 {
		return "unknown"
	}
	placement := params[0].In
	for _, param := range params {
		if param.In != placement {
			return "mixed"
		}
	}
	return placement
}
