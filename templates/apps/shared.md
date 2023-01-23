{{if .ShowSummary -}}
## Summary

This site features the documentation for `{{.Config.CliCmd}}` (CLI), and the API documentation
for Python and JavaScript clients in order to communicate with a local `{{.Config.DaemonCli}}`
instance through gRPC.
{{- end}}

## gRPC

The code samples assume that the there is a local `{{.Config.DaemonCli}}` instance
running and listening for gRPC connections on port `{{.Config.GrpcPort}}`. `{{upper .Name}}_DIR` will be used
as a placeholder to denote the base directory of the `{{.Config.DaemonCli}}` instance. By default,
this is `~/.{{lower .Name}}` on Linux and `~/Library/Application Support/{{pascal .Name}}` on macOS.

At the time of writing this documentation, two things are needed in order to
make a gRPC request to an `{{.Config.DaemonCli}}` instance: a TLS/SSL connection and a macaroon
used for RPC authentication. The code samples will show how these can
be used in order to make a successful, secure, and authenticated gRPC request.

The original `*.proto` files from which the gRPC documentation was generated
can be found here:

{{range .FileRepoUrls}}
- [`{{.Name}}.proto`]({{.GrpcUrl}})
{{- end}}

## REST

View a listing of all REST URLs on the [REST Endpoints](rest-endpoints) page. 

The code samples assume that the there is a local `{{.Config.DaemonCli}}` instance
running and listening for REST connections on port `{{.Config.RESTPort}}`. `{{upper .Name}}_DIR` will be used
as a placeholder to denote the base directory of the `{{.Config.DaemonCli}}` instance. By default,
this is `~/.{{lower .Name}}` on Linux and `~/Library/Application Support/{{pascal .Name}}` on macOS.

At the time of writing this documentation, two things are needed in order to
make an HTTP request to an `{{.Config.DaemonCli}}` instance: a TLS/SSL connection and a macaroon
used for RPC authentication. The code samples will show how these can
be used in order to make a successful, secure, and authenticated HTTP request.

The original `*.swagger.json` files from which the gRPC documentation was generated
can be found here:

{{range .FileRepoUrls}}
- [`{{.Name}}.swagger.json`]({{.RestUrl}})
{{- end}}

### REST Encoding

**NOTE**: The `byte` field type must be set as the base64 encoded string
representation of a raw byte array. Also, any time this must be used in a URL path
(ie. `/v1/abc/xyz/{payment_hash}`) the base64 string must be encoded using a
[URL and Filename Safe Alphabet](https://tools.ietf.org/html/rfc4648#section-5). This means you must replace `+` with `-`,
`/` with `_`, and keep the trailing `=` as is. Url encoding (ie. `%2F`) will not work.

<small>

This documentation was
[generated automatically](https://github.com/lightninglabs/lightning-api-ng) against commit
[`{{ .Config.Commit }}`]({{ .Config.RepoURL }}/tree/{{ .Config.Commit }}).

</small>