package parser

import (
	"errors"
	"fmt"
	"os"

	"github.com/lightninglabs/lightning-api-ng/config"
	"github.com/lightninglabs/lightning-api-ng/defs"
	"github.com/lightninglabs/lightning-api-ng/markdown"
	"github.com/lightninglabs/lightning-api-ng/models"
)

func ExportMarkdown(conf *config.Config, apiSpec *defs.ApiSpec) error {
	fmt.Println("Generating Markdown files for Docusaurus")
	fmt.Printf("App: %s\n", conf.App)
	fmt.Printf("Output Dir: %s\n", conf.AppOutputDir)

	// Remove the output dir if it already exists.
	if _, err := os.Stat(conf.AppOutputDir); !os.IsNotExist(err) {
		err := os.RemoveAll(conf.AppOutputDir)
		if err != nil {
			return err
		}
	}

	// Create an empty output dir.
	err := os.MkdirAll(conf.AppOutputDir, os.ModePerm)
	if err != nil {
		return err
	}

	templates, err := markdown.LoadAllTemplates()
	if err != nil {
		return err
	}

	fmt.Println("Exporting markdown files...")
	app := models.NewApp(conf, apiSpec, templates)
	err = app.ExportMarkdown()
	if err != nil {
		return err
	}

	// Save the root category file if it doesn't already exist.
	catFilePath := fmt.Sprintf("%s/%s", config.BaseOutputDir,
		models.CategoryFileName)
	if _, err := os.Stat(catFilePath); errors.Is(err, os.ErrNotExist) {
		err = models.WriteCategoryFile(
			catFilePath,
			"API Reference",
			"Documentation for the Lightning APIs",
		)
		if err != nil {
			return err
		}
	}

	return nil
}
