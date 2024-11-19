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
	Name      string
	Packages  map[string]*Package
	RestTypes *RestTypes

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
		RestTypes:   NewRestTypes(apiSpec.RESTTypes),
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

// RestEndpoints returns a list of all REST endpoints in Packages.
func (a *App) RestEndpoints() []*RestEndpoint {
	// Collect all of the services for all packages.
	services := make([]*Service, 0)
	for _, pkg := range a.Packages {
		services = append(services, pkg.Services...)
	}

	// Collect all of the methods for all services.
	methods := make([]*Method, 0)
	for _, service := range services {
		methods = append(methods, service.Methods...)
	}

	// Create a list of REST endpoints with a valid path.
	endpoints := make([]*RestEndpoint, 0)
	for _, method := range methods {
		if method.RestMapping == nil ||
			method.RestMapping.Path == "" {
			continue
		}
		ep := &RestEndpoint{
			RestPath:   method.RestMapping.Path,
			RestMethod: method.RestMapping.Method,
			LinkUrl: fmt.Sprintf(
				"../%s/%s/",
				markdown.ToKebabCase(method.Service.Name),
				markdown.ToKebabCase(method.Name),
			),
			MethodName: fmt.Sprintf("%s.%s",
				method.Service.Pkg.Name, method.Name),
		}
		endpoints = append(endpoints, ep)
	}

	// Sort the REST endpoints by path.
	sort.Slice(endpoints, func(i, j int) bool {
		return endpoints[i].RestPath < endpoints[j].RestPath
	})

	return endpoints
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
			a.RestTypes.UpdateMessage(msg)
			return msg, nil
		}
	}

	return nil, fmt.Errorf("cannot find message %s for %s in the %s "+
		"package", msgType, fullType, pkgName)
}

// GetNestedMessages recursively updates a map of all nested messages for the
// given message. The depth parameter is used to limit the recursion.
func (a *App) GetNestedMessages(message *Message,
	allMessages map[string]*Message, depth uint8) {

	for _, field := range message.Fields {
		// Only include the non-native field types (ex: lnrpc.OutPoint)
		if !strings.Contains(field.FullType, ".") {
			continue
		}

		msg, _ := a.GetMessage(field.FullType)

		// Add the message to the map if it was found.
		if msg != nil && depth > 0 {
			allMessages[field.FullType] = msg
			a.GetNestedMessages(msg, allMessages, depth-1)
		}
	}
}

// GetEnum returns the enum with the given full type name.
func (a *App) GetEnum(fullType string) (*Enum, error) {
	// Split "lnrpc.ChannelCloseSummary.ClosureType" into "lnrpc" and
	// "ChannelCloseSummary.ClosureType"
	period := strings.Index(fullType, ".")
	pkgName := fullType[:period]
	enumType := fullType[period+1:]

	if pkg, ok := a.Packages[pkgName]; ok {
		if enum, ok := pkg.Enums[enumType]; ok {
			return enum, nil
		}
	}

	return nil, fmt.Errorf("cannot find enum %s for %s in the %s package",
		enumType, fullType, pkgName)
}

// GetNestedEnums recursively updates a map of all nested enums for the given
// message. The depth parameter is used to limit the recursion.
func (a *App) GetNestedEnums(message *Message, allEnums map[string]*Enum,
	depth uint8) {

	for _, field := range message.Fields {
		// Only include the non-native field types (ex: lnrpc.OutPoint)
		if !strings.Contains(field.FullType, ".") {
			continue
		}

		enum, _ := a.GetEnum(field.FullType)

		// Add the enum to the map if it was found.
		if enum != nil {
			allEnums[field.FullType] = enum
			continue
		}

		// If the enum wasn't found, look for a nested message which
		// may have enum fields.
		msg, _ := a.GetMessage(field.FullType)
		if msg != nil && depth > 0 {
			// Search the nested messages for more enums.
			a.GetNestedEnums(msg, allEnums, depth-1)
		}
	}
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

	// Export a doc with the list of all rest paths.
	filePath := fmt.Sprintf("%s/rest-endpoints.md", a.Config.AppOutputDir)
	err := markdown.ExecuteEndpointsTemplate(a.Templates, a, filePath)
	if err != nil {
		return err
	}

	// Export the app index doc.
	filePath = fmt.Sprintf("%s/index.md", a.Config.AppOutputDir)
	fmt.Printf("Executing template %s for app %s\n", filePath, a.Name)
	err = markdown.ExecuteAppTemplate(a.Templates, a.Name, a, filePath)
	if err != nil {
		return err
	}

	return nil
}
