package parser

import (
	"fmt"

	"github.com/lightninglabs/lightning-api-ng/config"
	"github.com/lightninglabs/lightning-api-ng/defs"
)

func ExportMarkdown(conf *config.Config, apiSpec *defs.ApiSpec) error {
	fmt.Println("Generating Markdown files for Docusaurus")
	fmt.Printf("App: %s\n", conf.App)
	fmt.Printf("Output Dir: %s\n", conf.AppOutputDir)

	return nil
}
