package models

import (
	"github.com/lightninglabs/lightning-api-ng/defs"
)

type RestMapping struct {
	Method string
	Path   string
}

// NewRestMapping creates a new RestMapping from a RESTMapping definition.
func NewRestMapping(mappingDef defs.RESTMapping) *RestMapping {
	mapping := &RestMapping{
		Method: mappingDef.Method,
		Path:   mappingDef.Path,
	}
	return mapping
}
