# Lightning API Docs

This repo contains the source code used to generate the API docs websites for LND,
Loop, Pool, Faraday, and Taproot Assets Protocol: https://lightning.engineering/api-docs/

## Architecture

The websites are automatically generated from the protobuf definition files
(`*.proto`) and Swagger/OpenAPI files (`*.swagger.json`) found in the Github
repositories of each application. These files are sourced from the following
repos:

- LND: https://github.com/lightningnetwork/lnd
- Loop: https://github.com/lightninglabs/loop
- Pool: https://github.com/lightninglabs/pool
- Faraday: https://github.com/lightninglabs/faraday
- Taproot Assets Protocol: https://github.com/lightninglabs/taproot-assets

There is a multi-stage pipeline to consume the protobuf & Swagger files and
convert them into an HTML static website. The process is as follows:

```
+ ./generate.sh
  |
  | - downloads the master branch of each repo
  | - uses protoc to convert *.proto to generated.json
  | - compiles `cmd/mdgen/main.go` into ./mdgen console app
  |
+ ./mdgen
  |
  | - outputs markdown files into site/docs/
  |
+ yarn build
  |
  | - builds the Docusaurus site
  | - converts the markdown pages into HTML
  |
+ Output Static Website
```

This process is run every day at midnight UTC by a cron job. After the build,
the static sites are deployed.

## Development

Follow these steps to run the pipeline locally. You can modify the source
code for the `mdgen` tool and the markdown templates (in `templates/`), to
alter the generated static pages.

**System Requirements**

- [Go](https://golang.org) v1.18+
- [protoc](https://github.com/protocolbuffers/protobuf/releases/tag/v3.6.0) v3.6.0 (newer versions may not work as expected)
- [NodeJS](https://nodejs.org) v16.x or v18.x

#### Initial Setup

1. Clone this repo
   ```bash
   $ git clone https://github.com/lightninglabs/lightning-api-ng.git
   $ cd lightning-api-ng
   ```
1. Install NodeJS dependencies
   ```bash
   $ cd site/
   $ yarn
   ```

#### Active development

1. Run the script to generate JSON files for each daemon from their proto files
   ```bash
   $ ./generate.sh
   ```
1. In one terminal, start the script to watch for changes of source
   files and automatically regenerate the markdown files
   ```bash
   $ ./watch.sh
   ```
1. In a second terminal, start the Docusaurus site
   ```bash
   $ cd site/
   $ yarn start
   ```
1. A new browser window should automatically open http://localhost:3000/api-docs/

#### Building the static website for deployment

1. Run the script to generate JSON files for each daemon from their proto files
   ```bash
   $ ./generate.sh
   ```
1. Build the static site
   ```bash
   $ cd site/
   $ yarn build
   ```
