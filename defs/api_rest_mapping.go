package defs

import (
	"strings"

	"google.golang.org/genproto/googleapis/api/annotations"
)

type RESTMapping struct {
	Method  string       `json:"method"`
	Path    string       `json:"path"`
	Details *RESTDetails `json:"details"`
}

type RESTDetails struct {
	Summary     string           `json:"summary"`
	OperationID string           `json:"operationId"`
	Parameters  []*RESTParameter `json:"parameters"`
	Tags        []string         `json:"tags"`
}

type RESTParameter struct {
	Name   string      `json:"name"`
	Format string      `json:"format"`
	Type   string      `json:"type"`
	In     string      `json:"in"`
	Schema *RESTSchema `json:"schema"`
}

type RESTSchema struct {
	Ref        string                   `json:"$ref"`
	Type       string                   `json:"type"`
	Properties map[string]*RESTProperty `json:"properties"`
}

type RESTType struct {
	Type       string                   `json:"type"`
	Properties map[string]*RESTProperty `json:"properties"`
	Enum       []string                 `json:"enum"`
}

type RESTProperty struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Format      string `json:"format"`
	Ref         string `json:"$ref"`
}

func NewRESTMapping(rule *annotations.HttpRule,
	restPaths map[string]map[string]*RESTDetails) *RESTMapping {

	var m *RESTMapping
	switch {
	case rule.GetGet() != "":
		m = &RESTMapping{
			Method: "GET",
			Path:   rule.GetGet(),
		}

	case rule.GetPost() != "":
		m = &RESTMapping{
			Method: "POST",
			Path:   rule.GetPost(),
		}

	case rule.GetDelete() != "":
		m = &RESTMapping{
			Method: "DELETE",
			Path:   rule.GetDelete(),
		}

	case rule.GetPatch() != "":
		m = &RESTMapping{
			Method: "PATCH",
			Path:   rule.GetPatch(),
		}

	case rule.GetPut() != "":
		m = &RESTMapping{
			Method: "PUT",
			Path:   rule.GetPut(),
		}

	case rule.GetCustom() != nil:
		m = &RESTMapping{
			Method: rule.GetCustom().Kind,
			Path:   rule.GetCustom().Path,
		}

	default:
		m = &RESTMapping{
			Method: "UNKNOWN",
			Path:   "",
		}
	}

	methods, ok := restPaths[m.Path]
	if !ok {
		return m
	}

	m.Details = methods[strings.ToLower(m.Method)]

	return m
}
