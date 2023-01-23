package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	// BaseOutputDir is the base directory to output the generated markdown
	// files.
	BaseOutputDir = "./site/docs/api"
)

type Config struct {
	App                  string
	RepoURL              string
	Commit               string
	ProtoSrcDir          string
	ExperimentalPackages []string
	GrpcPort             uint16
	RESTPort             uint16
	CliCmd               string
	DaemonCli            string

	// AppOutputDir is the directory to output the generated markdown files.
	AppOutputDir string
}

func NewConfig(app string) (*Config, error) {
	config := &Config{
		App:         app,
		RepoURL:     os.Getenv("REPO_URL"),
		Commit:      os.Getenv("COMMIT"),
		ProtoSrcDir: os.Getenv("PROTO_SRC_DIR"),
		ExperimentalPackages: strings.Split(
			os.Getenv("EXPERIMENTAL_PACKAGES"), " ",
		),
		CliCmd:       os.Getenv("COMMAND"),
		DaemonCli:    os.Getenv("DAEMON"),
		AppOutputDir: fmt.Sprintf("%s/%s", BaseOutputDir, app),
	}

	// If the proto source dir is not empty, make sure we can use it
	// directly by appending a tailing path separator.
	if config.ProtoSrcDir != "" {
		config.ProtoSrcDir = fmt.Sprintf("%s/", config.ProtoSrcDir)
	}
	port, err := strconv.ParseUint(os.Getenv("GRPC_PORT"), 10, 16)
	if err != nil {
		return nil, err
	}
	config.GrpcPort = uint16(port)
	port, err = strconv.ParseUint(os.Getenv("REST_PORT"), 10, 16)
	if err != nil {
		return nil, err
	}
	config.RESTPort = uint16(port)

	return config, nil
}
