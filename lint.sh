#!/bin/bash
set -e

# Build the Docker image for the linter tool
docker build -q -t api-ng-tools ./tools

# Run the linter tool
docker run -v $(pwd):/build api-ng-tools golangci-lint run -v 
