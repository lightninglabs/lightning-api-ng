#!/bin/bash
set -e

mkdir -p build
go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest

go build -o mdgen ./cmd/mdgen/


function initrepo() {
  PROTO_DIR=$PROTO_ROOT_DIR/$COMPONENT
  LOCAL_REPO_PATH=./build/$COMPONENT
  if [[ ! -d $LOCAL_REPO_PATH ]]; then
    git clone $REPO_URL $LOCAL_REPO_PATH
  fi

  # Update the repository to the respective CHECKOUT_COMMIT and install the binary.
  pushd $LOCAL_REPO_PATH
  COMMIT=$(git rev-parse HEAD)
  git reset --hard HEAD
  git pull
  git checkout $CHECKOUT_COMMIT
  COMMIT=$(git rev-parse HEAD)
  eval $INSTALL_CMD
  popd

  # Copy over all proto and json files from the checked out source directory.
  mkdir -p $PROTO_DIR
  rsync -a --prune-empty-dirs --include '*/' --include '*.proto' \
    --include '*.json' --include '*.yaml' --exclude '*' \
    $LOCAL_REPO_PATH/$PROTO_SRC_DIR/ $PROTO_DIR/

  # Copy any additional proto files from other projects.
  INCLUDE_FLAG=""
  if [[ ! -z $INCLUDE_PROTOS ]]; then
    echo "Including protos from $INCLUDE_PROTOS"
    INCLUDE_FLAG="-I$INCLUDE_PROTOS"
  fi

  pushd $PROTO_DIR
  proto_files=$(find . -name '*.proto' -not -path "$EXCLUDE_PROTOS")
  
  protoc -I. -I/usr/local/include $INCLUDE_FLAG \
    --doc_out=json,generated.json:. $proto_files
  popd
}

function compile() {
  echo "Using ${COMPONENT} repo URL ${REPO_URL} and commit ${CHECKOUT_COMMIT}"

  if [[ "${INIT_REPOS}" == "true" ]]; then
    initrepo
  else
    # If we're not initializing the repos, we still need to the COMMIT variable
    pushd ./build/$COMPONENT > /dev/null
    COMMIT=$(git rev-parse HEAD)
    popd > /dev/null
  fi
  
  export REPO_URL COMMIT PROTO_SRC_DIR EXPERIMENTAL_PACKAGES GRPC_PORT REST_PORT COMMAND DAEMON
  ./mdgen $COMPONENT
}

# Generic options.
WS_ENABLED="${WS_ENABLED:-true}"
LND_FORK="${LND_FORK:-lightningnetwork}"
LND_COMMIT="${LND_COMMIT:-master}"
LOOP_FORK="${LOOP_FORK:-lightninglabs}"
LOOP_COMMIT="${LOOP_COMMIT:-master}"
FARADAY_FORK="${FARADAY_FORK:-lightninglabs}"
FARADAY_COMMIT="${FARADAY_COMMIT:-master}"
POOL_FORK="${POOL_FORK:-lightninglabs}"
POOL_COMMIT="${POOL_COMMIT:-master}"
TAPD_FORK="${TAPD_FORK:-lightninglabs}"
TAPD_COMMIT="${TAPD_COMMIT:-main}"
LIT_FORK="${LIT_FORK:-lightninglabs}"
LIT_COMMIT="${LIT_COMMIT:-master}"
PROTO_ROOT_DIR="build/protos"

# Set to 'false' to skip cloning and building each repo 
INIT_REPOS="${INIT_REPOS:-true}"

# Remove previously generated templates.
if [[ "${INIT_REPOS}" == "true" ]]; then
  rm -rf $PROTO_ROOT_DIR
fi

########################
## Compile docs for lnd
########################
REPO_URL="https://github.com/${LND_FORK}/lnd"
CHECKOUT_COMMIT=$LND_COMMIT
COMPONENT=lnd
COMMAND=lncli
DAEMON=lnd
PROTO_SRC_DIR=lnrpc
EXCLUDE_PROTOS="none"
EXPERIMENTAL_PACKAGES="autopilotrpc signrpc walletrpc chainrpc invoicesrpc watchtowerrpc neutrinorpc monitoring peersrpc kvdb_postgres kvdb_etcd"
INSTALL_CMD="make clean && make install tags=\"$EXPERIMENTAL_PACKAGES\""
GRPC_PORT=10009
REST_PORT=8080
compile

########################
## Compile docs for loop
########################
REPO_URL="https://github.com/${LOOP_FORK}/loop"
CHECKOUT_COMMIT=$LOOP_COMMIT
COMPONENT=loop
COMMAND=loop
DAEMON=loopd
PROTO_SRC_DIR=""
EXCLUDE_PROTOS="./swapserverrpc/*"
EXPERIMENTAL_PACKAGES=""
INSTALL_CMD="make install"
GRPC_PORT=11010
REST_PORT=8081
compile

########################
## Compile docs for faraday
########################
REPO_URL="https://github.com/${FARADAY_FORK}/faraday"
CHECKOUT_COMMIT=$FARADAY_COMMIT
COMPONENT=faraday
COMMAND=frcli
DAEMON=faraday
PROTO_SRC_DIR=frdrpc
EXCLUDE_PROTOS="none"
EXPERIMENTAL_PACKAGES=""
INSTALL_CMD="make install"
GRPC_PORT=8465
REST_PORT=8082
compile

########################
## Compile docs for pool
########################
REPO_URL="https://github.com/${POOL_FORK}/pool"
CHECKOUT_COMMIT=$POOL_COMMIT
COMPONENT=pool
COMMAND=pool
DAEMON=poold
PROTO_SRC_DIR=""
EXCLUDE_PROTOS="none"
EXCLUDE_SERVICES="ChannelAuctioneer"
EXPERIMENTAL_PACKAGES=""
INSTALL_CMD="make install"
GRPC_PORT=12010
REST_PORT=8281
compile

########################
## Compile docs for lit
########################
REPO_URL="https://github.com/${LIT_FORK}/lightning-terminal"
CHECKOUT_COMMIT=$LIT_COMMIT
COMPONENT=lit
COMMAND=litcli
DAEMON=litd
PROTO_SRC_DIR="litrpc"
EXCLUDE_PROTOS="none"
EXCLUDE_SERVICES=""
EXPERIMENTAL_PACKAGES=""
INSTALL_CMD="make go-install"
GRPC_PORT=8443
REST_PORT=8443
compile

########################
## Compile docs for taproot-assets
########################

# This must come last, because some of the proto files reference litcli
# commands. So we need to have litcli installed first.

REPO_URL="https://github.com/${TAPD_FORK}/taproot-assets"
CHECKOUT_COMMIT=$TAPD_COMMIT
COMPONENT=taproot-assets
COMMAND=tapcli
DAEMON=tapd
PROTO_SRC_DIR="taprpc"
INCLUDE_PROTOS="../lnd"
EXCLUDE_PROTOS="none"
EXCLUDE_SERVICES=""
EXPERIMENTAL_PACKAGES=""
INSTALL_CMD="make install"
GRPC_PORT=10029
REST_PORT=8089
compile