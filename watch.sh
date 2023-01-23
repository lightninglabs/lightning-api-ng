#!/bin/bash
set -e

# Perform the initial generation which clones the repos, builds the binaries,
# runs protoc, and builds/runs the merger tool
# ./generate.sh

# 
INIT_REPOS=false nodemon --exec "./generate.sh || exit 1" --signal SIGTERM --watch . --ext go,md --ignore ./build --ignore ./site 