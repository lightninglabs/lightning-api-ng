package markdown

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

const (
	// templatesDir is the directory where the markdown templates are
	// stored.
	templatesDir = "./templates"
)

var (
	// pathGlobs is a list of globs that are used to load all templates.
	pathGlobs = []string{"*.md", "apps/*.md"}
)

// LoadAllTemplates loads all templates from the templates directory.
func LoadAllTemplates() (*template.Template, error) {
	templates := template.New("").Funcs(getTemplateHelpers())
	var err error
	fmt.Printf("Loading templates from %s\n", templatesDir)
	for _, glob := range pathGlobs {
		fullGlob := fmt.Sprintf("%s/%s", templatesDir, glob)
		templates, err = templates.ParseGlob(fullGlob)
		if err != nil {
			return nil,
				fmt.Errorf("error loading templates: %v", err)
		}
	}
	for _, tpl := range templates.Templates() {
		fmt.Printf("Loaded template %s\n", tpl.Name())
	}
	return templates, nil
}

// ExecuteMethodTemplate executes the method template with the given data and
// writes the result to the output file.
func ExecuteMethodTemplate(templates *template.Template, data any,
	outputFilePath string) error {

	return execute(templates, "method.md", data, outputFilePath)
}

// ExecuteEndpointsTemplate executes the rest_endpoints template with the given
// data and writes the result to the output file.
func ExecuteEndpointsTemplate(templates *template.Template, data any,
	outputFilePath string) error {

	return execute(templates, "rest_endpoints.md", data, outputFilePath)
}

// ExecuteAppTemplate executes the app template with the given data and writes
// the result to the output file.
func ExecuteAppTemplate(templates *template.Template, appName string, data any,
	outputFilePath string) error {

	return execute(templates, fmt.Sprintf("%s.md", appName), data,
		outputFilePath)
}

// execute executes the template with the given data and writes the result to
// the output file.
func execute(templates *template.Template, fileName string, data any,
	outputFilePath string) error {
	// Create the output file.
	file, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// Execute the template on the supplied data.
	err = templates.ExecuteTemplate(writer, fileName, data)
	if err != nil {
		return err
	}

	// Flush the writer to ensure the file is written to disk.
	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

// getTemplateHelpers returns the shared helper functions that can be used in
// all templates.
func getTemplateHelpers() template.FuncMap {
	return template.FuncMap{
		"lower":     strings.ToLower,
		"upper":     strings.ToUpper,
		"pascal":    strcase.ToCamel,
		"camel":     strcase.ToLowerCamel,
		"snake":     strcase.ToSnake,
		"kebab":     ToKebabCase,
		"multiArgs": multiArgs,
	}
}

// multiArgs is a helper function that can be used in templates to create a
// map of key value pairs. It takes an even number of arguments and returns a
// map of the arguments. The first argument is the key and the second is the
// value. The third argument is the next key and so on. This is useful for
// passing multiple arguments to a nested template.
func multiArgs(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, fmt.Errorf("multiArgs must have an even number of " +
			"args")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, fmt.Errorf("multiArgs keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}
