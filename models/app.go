package models

import (
	"fmt"
	"sort"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/lightninglabs/lightning-api-ng/config"
	"github.com/lightninglabs/lightning-api-ng/defs"
	"github.com/lightninglabs/lightning-api-ng/markdown"
	"golang.org/x/exp/slices"
)

type App struct {
	Name     string
	Packages map[string]*Package

	ShowSummary  bool
	FileRepoUrls []FileRepoUrl
	Config       *config.Config
	Templates    *template.Template
}

type FileRepoUrl struct {
	Name    string
	GrpcUrl string
	RestUrl string
}

type RestEndpoint struct {
	RestPath   string
	RestMethod string
	LinkUrl    string
	MethodName string
}

// NewApp creates a new app model from the api spec containing all the proto
// and swagger definitions.
func NewApp(config *config.Config, apiSpec *defs.ApiSpec,
	templates *template.Template) *App {

	app := &App{
		Name:        config.App,
		ShowSummary: config.App != "lnd",
		Config:      config,
		Templates:   templates,

		Packages:     make(map[string]*Package),
		FileRepoUrls: make([]FileRepoUrl, 0),
	}
	for _, file := range apiSpec.Files {
		pkg, ok := app.Packages[file.Package]
		if !ok {
			pkg = NewPackage(file.Package, app)
			app.Packages[file.Package] = pkg
		}
		pkg.addProtoFile(file)

		if slices.Contains(config.ExperimentalPackages, pkg.Name) {
			pkg.Experimental = true
		}

		baseName := strings.ReplaceAll(file.Name, ".proto", "")
		baseUrl := fmt.Sprintf("%s/blob/%s/%s", config.RepoURL,
			config.Commit, config.ProtoSrcDir)
		urls := FileRepoUrl{
			Name:    strings.Clone(baseName),
			GrpcUrl: fmt.Sprintf("%s%s.proto", baseUrl, baseName),
			RestUrl: fmt.Sprintf("%s%s.swagger.json", baseUrl,
				baseName),
		}
		app.FileRepoUrls = append(app.FileRepoUrls, urls)
	}

	// Sort the file repo urls by name.
	sort.Slice(app.FileRepoUrls, func(i, j int) bool {
		return app.FileRepoUrls[i].Name < app.FileRepoUrls[j].Name
	})

	return app
}

// CamelName returns the app name in camel case or in all caps if the app is
// lnd.
func (a *App) PascalName() string {
	if a.Name == "lnd" {
		return "LND"
	}
	return strcase.ToCamel(a.Name)
}

type ExperimentalService struct {
	Name      string
	LowerName string
	File      string
}

// ExperimentalServices returns a list of all services in Packages that are
// marked as experimental.
func (a *App) ExperimentalServices() []*ExperimentalService {
	services := make([]*ExperimentalService, 0)
	for _, pkg := range a.Packages {
		if !pkg.Experimental {
			continue
		}

		fmt.Printf("Experimental package %s\n", pkg.Name)
		for _, service := range pkg.Services {
			services = append(services, &ExperimentalService{
				Name:      service.Name,
				LowerName: strings.ToLower(service.Name),
				File:      service.FileName,
			})
		}
	}

	// sort the services
	sort.Slice(services, func(i, j int) bool {
		return services[i].Name < services[j].Name
	})
	return services
}

// GetMessage returns the message with the given full type name.
func (a *App) GetMessage(fullType string) (*Message, error) {
	// Split "lnrpc.Invoice.InvoiceState" into "lnrpc" and
	// "Invoice.InvoiceState"
	period := strings.Index(fullType, ".")
	pkgName := fullType[:period]
	msgType := fullType[period+1:]

	if pkg, ok := a.Packages[pkgName]; ok {
		if msg, ok := pkg.Messages[msgType]; ok {
			return msg, nil
		}
	}

	return nil, fmt.Errorf("cannot find message %s for %s in the %s "+
		"package", msgType, fullType, pkgName)
}

// ExportMarkdown exports the app as markdown.
func (a *App) ExportMarkdown() error {
	fmt.Printf("Exporting app %s\n", a.Name)

	for _, pkg := range a.Packages {
		err := pkg.ExportMarkdown()
		if err != nil {
			return err
		}
	}

	// Export the app index doc.
	filePath := fmt.Sprintf("%s/index.md", a.Config.AppOutputDir)
	fmt.Printf("Executing template %s for app %s\n", filePath, a.Name)
	err := markdown.ExecuteAppTemplate(a.Templates, a.Name, a, filePath)
	if err != nil {
		return err
	}

	return nil
}
