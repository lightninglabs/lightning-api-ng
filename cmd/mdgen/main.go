package main

import (
	"fmt"
	"os"

	"github.com/lightninglabs/lightning-api-ng/config"
	"github.com/lightninglabs/lightning-api-ng/parser"
)

func main() {
	app := os.Args[1]
	cfg, err := config.NewConfig(app)
	if err != nil {
		fail(err)
	}
	fmt.Printf("\nLoaded config: %+v\n", cfg)

	apiSpec := parser.LoadApiSpec(cfg)

	err = parser.ExportMarkdown(cfg, apiSpec)
	if err != nil {
		fail(err)
	}
}

func fail(err error) {
	fmt.Printf("Error: %v\n", err)
	os.Exit(1)
}
