package models

import (
	"fmt"

	"github.com/lightninglabs/lightning-api-ng/defs"
)

type Package struct {
	Name         string
	Description  string
	Experimental bool
	Messages     map[string]*Message
	Services     []*Service

	App *App
}

// NewPackage creates a new package model from a package definition
func NewPackage(name string, app *App) *Package {
	fmt.Printf("Creating package %s\n", name)
	return &Package{
		Name:     name,
		App:      app,
		Messages: make(map[string]*Message),
		Services: make([]*Service, 0),
	}
}

// AddProtoFile adds the messages, enums, and services from the given proto
// file definition to the package.
func (p *Package) addProtoFile(file *defs.File) {
	fmt.Printf("Adding proto file %s to package %s\n", file.Name,
		file.Package)

	fmt.Printf("Adding %d messages to package %s\n", len(file.Messages),
		file.Package)
	for _, msgDef := range file.Messages {
		msg := NewMessage(msgDef, p, file.Name)
		p.Messages[msg.LongName] = msg
	}

	fmt.Printf("Adding %d services to package %s\n", len(file.Services),
		file.Package)
	for _, svcDef := range file.Services {
		svc := NewService(svcDef, p, file.Name)
		p.Services = append(p.Services, svc)
	}
}

// ExportMarkdown exports the services in this package to markdown.
func (p *Package) ExportMarkdown() error {
	fmt.Printf("Exporting package %s\n", p.Name)
	for _, svc := range p.Services {
		err := svc.ExportMarkdown()
		if err != nil {
			return err
		}
	}
	return nil
}
