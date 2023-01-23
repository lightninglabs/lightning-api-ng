package models

import (
	"fmt"

	"github.com/lightninglabs/lightning-api-ng/markdown"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
)

type Enum struct {
	pkg      *Package
	FileName string

	Name        string
	LongName    string
	FullName    string
	Description string
	Values      []*EnumValue
}

type EnumValue struct {
	Name        string
	Number      string
	Description string
}

// NewEnum creates a new Enum model from an Enum definition
func NewEnum(enumDef *gendoc.Enum, pkg *Package, fileName string) *Enum {
	fmt.Printf("Creating enum %s with %d values\n", enumDef.FullName,
		len(enumDef.Values))

	enum := &Enum{
		pkg:      pkg,
		FileName: fileName,
		Name:     enumDef.Name,
		LongName: enumDef.LongName,
		FullName: enumDef.FullName,
		Description: markdown.CleanDescription(
			enumDef.Description, true),

		Values: make([]*EnumValue, 0),
	}
	for _, valueDef := range enumDef.Values {
		value := &EnumValue{
			Name:   valueDef.Name,
			Number: valueDef.Number,
			Description: markdown.CleanDescription(
				valueDef.Description, true),
		}
		enum.Values = append(enum.Values, value)
	}

	return enum
}
