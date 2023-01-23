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
	if c.AppName() == "lnd" {
		return "LND_DIR/data/chain/bitcoin/regtest/admin.macaroon"
	}

	return fmt.Sprintf("%s_DIR/regtest/%s.macaroon",
		strings.ToUpper(c.AppName()), c.AppName())
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
