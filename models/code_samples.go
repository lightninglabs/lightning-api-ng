package models

import (
	"fmt"
	"path"
	"strings"

	"golang.org/x/exp/slices"
)

// CodeSamples is a view model for the code samples markdown template. Using
// these struct methods simplifies the logic needed in the markdown templates
type CodeSamples struct {
	Method *Method
}

func NewCodeSamples(method *Method) *CodeSamples {
	return &CodeSamples{
		Method: method,
	}
}

func (c *CodeSamples) AppName() string {
	return c.Method.Service.Pkg.App.Name
}

func (c *CodeSamples) PackageName() string {
	return c.Method.Service.Pkg.Name
}

func (c *CodeSamples) ServiceName() string {
	return c.Method.Service.Name
}

func (c *CodeSamples) MethodName() string {
	return c.Method.Name
}

func (c *CodeSamples) ProtoFileName() string {
	return c.Method.Service.FileName
}

func (c *CodeSamples) LoaderFiles() string {
	if c.AppName() == "lnd" {
		if c.ProtoFileName() == "lightning.proto" {
			return "'lightning.proto'"
		}

		return fmt.Sprintf("['lightning.proto', '%s']",
			c.ProtoFileName())
	}
	return fmt.Sprintf("'%s'", c.ProtoFileName())
}

func (c *CodeSamples) MacaroonPath() string {
	switch c.AppName() {
	case "lnd":
		return "LND_DIR/data/chain/bitcoin/regtest/admin.macaroon"
	case "taproot-assets":
		return "TAPD_DIR/data/regtest/admin.macaroon"
	default:
		return fmt.Sprintf("%s_DIR/regtest/%s.macaroon",
			strings.ToUpper(c.AppName()), c.AppName())
	}
}

// TlsPath returns the TLS cert path with placeholder directory for use
// in Python/JS code samples.
func (c *CodeSamples) TlsPath() string {
	dirVar := fmt.Sprintf("%s_DIR", strings.ToUpper(c.AppName()))
	if c.AppName() == "taproot-assets" {
		dirVar = "TAPD_DIR"
	}
	switch c.AppName() {
	case "loop", "faraday", "pool":
		return fmt.Sprintf("%s/regtest/tls.cert", dirVar)
	default:
		return fmt.Sprintf("%s/tls.cert", dirVar)
	}
}

func (c *CodeSamples) RequiresMacaroon() bool {
	anonServices := []string{"lnrpc.WalletUnlocker", "lnrpc.State"}
	return !slices.Contains(anonServices, c.Method.Service.FullName)
}

func (c *CodeSamples) GrpcPort() uint16 {
	return c.Method.Service.Pkg.App.Config.GrpcPort
}

func (c *CodeSamples) RestPort() uint16 {
	return c.Method.Service.Pkg.App.Config.RESTPort
}

func (c *CodeSamples) RequestName() string {
	return c.Method.Request().Name
}

func (c *CodeSamples) RequestFields() []*Field {
	return c.Method.Request().Fields
}

func (c *CodeSamples) ResponseFields() []*Field {
	return c.Method.Response().Fields
}

func (c *CodeSamples) IsUnary() bool {
	return c.Method.StreamingDirection() == ""
}

func (c *CodeSamples) IsServerStreaming() bool {
	return c.Method.StreamingDirection() == "server"
}

func (c *CodeSamples) IsClientStreaming() bool {
	return c.Method.StreamingDirection() == "client"
}

func (c *CodeSamples) IsBidirectionalStreaming() bool {
	return c.Method.StreamingDirection() == "bidirectional"
}

func (c *CodeSamples) IsStreaming() bool {
	return !c.IsUnary()
}

func (c *CodeSamples) IsRestPost() bool {
	return c.Method.RestMethod() == "POST"
}

func (c *CodeSamples) StubFileName() string {
	return strings.ReplaceAll(path.Base(c.ProtoFileName()), ".proto", "")
}

func (c *CodeSamples) DefaultDir() string {
	switch c.AppName() {
	case "lnd":
		return "~/.lnd"
	case "taproot-assets":
		return "~/.tapd"
	default:
		return fmt.Sprintf("~/.%s", c.AppName())
	}
}

func (c *CodeSamples) ShellMacaroonPath() string {
	switch c.AppName() {
	case "lnd":
		return fmt.Sprintf("$%s/data/chain/bitcoin/$NETWORK/admin.macaroon",
			c.DirVarName())
	case "taproot-assets":
		return fmt.Sprintf("$%s/data/$NETWORK/admin.macaroon",
			c.DirVarName())
	default:
		appName := c.AppName()
		return fmt.Sprintf("$%s/$NETWORK/%s.macaroon",
			c.DirVarName(), appName)
	}
}

func (c *CodeSamples) ShellTlsPath() string {
	switch c.AppName() {
	case "loop", "faraday", "pool":
		return fmt.Sprintf("$%s/$NETWORK/tls.cert", c.DirVarName())
	default:
		return fmt.Sprintf("$%s/tls.cert", c.DirVarName())
	}
}

func (c *CodeSamples) DirVarName() string {
	if c.AppName() == "taproot-assets" {
		return "TAPD_DIR"
	}
	return fmt.Sprintf("%s_DIR", strings.ToUpper(c.AppName()))
}

func (c *CodeSamples) RepoURL() string {
	return c.Method.Service.Pkg.App.Config.RepoURL
}

func (c *CodeSamples) ProtoSrcDir() string {
	return c.Method.Service.Pkg.App.Config.ProtoSrcDir
}

func (c *CodeSamples) SourceVarName() string {
	if c.AppName() == "taproot-assets" {
		return "TAPD_SOURCE"
	}
	return fmt.Sprintf("%s_SOURCE", strings.ToUpper(c.AppName()))
}

func (c *CodeSamples) SourceRepoName() string {
	// Extract repo name from URL like "https://github.com/lightningnetwork/lnd"
	parts := strings.Split(c.RepoURL(), "/")
	return parts[len(parts)-1]
}

// NeedsLndProtos returns true if this service's proto file imports LND
// proto types from a separate repository.
func (c *CodeSamples) NeedsLndProtos() bool {
	if c.AppName() == "lnd" {
		return false
	}
	// Loop's SwapClient imports lnrpc/lightning.proto
	if c.AppName() == "loop" && c.ServiceName() == "SwapClient" {
		return true
	}
	// Taproot-Assets' TaprootAssetChannels imports lightning.proto
	// and routerrpc/router.proto
	if c.AppName() == "taproot-assets" &&
		c.ServiceName() == "TaprootAssetChannels" {
		return true
	}
	return false
}

// GrpcurlImportPaths returns the -import-path flags for grpcurl.
func (c *CodeSamples) GrpcurlImportPaths() string {
	srcVar := "$" + c.SourceVarName()
	protoSrc := c.ProtoSrcDir()

	var mainPath string
	if protoSrc != "" {
		// e.g. $LND_SOURCE/lnrpc/
		mainPath = srcVar + "/" + protoSrc
	} else {
		mainPath = srcVar
	}

	result := fmt.Sprintf("-import-path %s", mainPath)

	if c.NeedsLndProtos() {
		if c.AppName() == "loop" {
			// loop imports "lnrpc/lightning.proto", so we need
			// the LND repo root.
			result += " \\\n    -import-path $LND_SOURCE"
		} else if c.AppName() == "taproot-assets" {
			// tapd imports "lightning.proto" and
			// "routerrpc/router.proto" directly, so we need
			// the lnrpc/ subdir.
			result += " \\\n    -import-path $LND_SOURCE/lnrpc"
		}
	}

	return result
}

// ShellRequestData returns the JSON request body for curl/grpcurl -d flag,
// with proper comma separation (no trailing comma).
func (c *CodeSamples) ShellRequestData() string {
	fields := c.RequestFields()
	if len(fields) == 0 {
		return ""
	}
	parts := make([]string, len(fields))
	for i, f := range fields {
		val := fmt.Sprintf("<%s>", f.Type)
		if f.Encoding() != "" {
			val = "BASE64_ENCODED_VALUE"
		}
		parts[i] = fmt.Sprintf("\"%s\": %s", f.Name, val)
	}
	return "{ " + strings.Join(parts, ", ") + " }"
}

func (c *CodeSamples) PythonRestArgs() string {
	args := make([]string, 0)
	if c.RequiresMacaroon() {
		args = append(args, "headers=headers, ")
	}
	if c.Method.ResponseStreaming {
		args = append(args, "stream=True, ")
	}
	if c.IsRestPost() {
		args = append(args, "data=json.dumps(data), ")
	}
	return strings.Join(args, "")
}
