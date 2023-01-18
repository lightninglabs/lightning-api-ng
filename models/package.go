package models

import (
	"fmt"

	"github.com/lightninglabs/lightning-api-ng/defs"
)

type Package struct {
	Name         string
	Description  string
	Experimental bool

	App *App
}

// NewPackage creates a new package model from a package definition
func NewPackage(name string, app *App) *Package {
	fmt.Printf("Creating package %s\n", name)
	return &Package{
		Name: name,
		App:  app,
	}
}

// AddProtoFile adds the messages, enums, and services from the given proto
// file definition to the package.
func (p *Package) addProtoFile(file *defs.File) {
	fmt.Printf("Adding proto file %s to package %s\n", file.Name,
		file.Package)
}

// ExportMarkdown exports the services in this package to markdown.
func (p *Package) ExportMarkdown() error {
	fmt.Printf("Exporting package %s\n", p.Name)
	return nil
}
