package models

import (
	"fmt"
	"os"

	"github.com/iancoleman/strcase"
	"github.com/lightninglabs/lightning-api-ng/defs"
)

type Service struct {
	Pkg      *Package
	FileName string

	Name        string
	FullName    string
	Description string
}

// NewService creates a new Service model from a Service definition.
func NewService(svcDef *defs.Service, pkg *Package, fileName string) *Service {
	fmt.Printf("Creating service %s with %d methods\n", svcDef.Name,
		len(svcDef.Methods))

	svc := &Service{
		Pkg:         pkg,
		FileName:    fileName,
		FullName:    svcDef.FullName,
		Name:        svcDef.Name,
		Description: svcDef.Description,
	}
	return svc
}

// ExportMarkdown exports the service to a category JSON file and exports its
// methods to markdown.
func (s *Service) ExportMarkdown() error {
	fmt.Printf("Exporting service %s\n", s.Name)
	svcDirName := strcase.ToKebab(s.Name)
	servicePath := fmt.Sprintf("%s/%s", s.Pkg.App.Config.AppOutputDir,
		svcDirName)

	// Create the service dir if it doesn't exists.
	if _, err := os.Stat(servicePath); os.IsNotExist(err) {
		err := os.MkdirAll(servicePath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Save the category file.
	catFilePath := fmt.Sprintf("%s/%s", servicePath, CategoryFileName)
	err := WriteCategoryFile(
		catFilePath,
		fmt.Sprintf("%s Service", s.Name),
		fmt.Sprintf("Documentation for the %s Service", s.Name),
	)
	if err != nil {
		return err
	}
	return nil
}
