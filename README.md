# Lightning API Docs

This repo contains the source code used to generate the API docs websites for:

- LND: https://api.lightning.community/
- Loop, Pool, Faraday, and Taro: https://lightning.engineering/api-docs/

## Architecture

The websites are automatically generated from the protobuf definition files
(`*.proto`) and Swagger/OpenAPI files (`*.swagger.json`) found in the Github
repositories of each application. These files are sourced from the following
repos:

- LND: https://github.com/lightningnetwork/lnd
- Loop: https://github.com/lightninglabs/loop
- Pool: https://github.com/lightninglabs/pool
- Faraday: https://github.com/lightninglabs/faraday
- Taro: https://github.com/lightninglabs/taro

There is a multi-stage pipeline to process the protobuf & Swagger files and
convert them into an HTML static website. The process is as follows:

```
+ ./generate.sh
  |
  | - downloads the master branch of each repo
  | - uses protoc to convert *.proto to generated.json
  | - compiles main.go into ./merger app
  |
+ ./merger
  |
  | - merges *.swagger.json and generated.json into build/app.json
  |
+ yarn generate-docs
  |
  | - consumes build/*.json and site/tools/templates/*.md
  | - outputs markdown files into site/docs/
  | - outputs site config into site/build.config.json
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
code for the `merger` tool (golang), markdown generator (Typescript), and
the markdown templates (in `site/rc/templates`), to alter the generated static
pages.

**System Requirements**

- [Go](https://golang.org) v1.18+
- [protoc](https://github.com/protocolbuffers/protobuf/releases/tag/v3.6.0) v3.6.0 (newer versions may not work as expected)
- [NodeJS](https://nodejs.org) v16.x or v18.x

1. Clone this repo
   ```bash
   $ git clone https://github.com/lightninglabs/lightning-api-ng.git
   $ cd lightning-api-ng
   ```
1. Run the script to generate JSON files for each daemon from their proto files
   ```bash
   $ ./generate.sh
   ```
1. Install NodeJS dependencies
   ```bash
   $ cd site/
   $ yarn
   ```
1. Run the website locally
   1. In one terminal, start the Docusaurus site
      ```bash
      $ yarn start
      ```
   1. In a second terminal, start the script to watch for changes of source
      files and automatically regenerate the markdown files
      ```bash
      # Run one of the commands below.
      # For LND:
      $ yarn watch-lnd
      # For Loop, Pool, Faraday, Taro:
      $ yarn watch-labs
      ```

## Building the static website

1. Run the script to generate the markdown pages from the JSON files
   ```bash
   # Run one of the commands below.
   # For LND:
   $ BUILD_CONFIG=lnd yarn generate-docs
   # For Loop, Pool, Faraday, Taro:
   $ BUILD_CONFIG=labs yarn generate-docs
   ```
1. Build the static site
   ```bash
   $ yarn build
   ```
