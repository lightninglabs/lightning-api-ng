```bash
# grpcurl docs: https://github.com/fullstorydev/grpcurl
# Proto source: {{.RepoURL}}{{if .NeedsLndProtos}}
# Also requires: https://github.com/lightningnetwork/lnd{{end}}
GRPC_HOST=localhost:{{.GrpcPort}}
{{.DirVarName}}={{.DefaultDir}}
{{.SourceVarName}}=path/to/{{.SourceRepoName}}{{if .NeedsLndProtos}}
LND_SOURCE=path/to/lnd{{end}}{{if .RequiresMacaroon}}
NETWORK=mainnet
MACAROON_PATH="{{.ShellMacaroonPath}}"{{end}}
TLS_PATH="{{.ShellTlsPath}}"
{{if or .IsClientStreaming .IsBidirectionalStreaming}}
# Client streaming or bidirectional streaming RPCs are not easily supported
# by grpcurl. Please use a programming language client for this RPC.{{else}}
grpcurl \
    {{.GrpcurlImportPaths}} \
    -proto {{.ProtoFileName}} \
    -cacert $TLS_PATH{{if .RequiresMacaroon}} \
    -H "macaroon: $(xxd -ps -u -c 1000 $MACAROON_PATH)"{{end}}{{if .RequestFields}} \
    -d '{{.ShellRequestData}}'{{end}} \
    $GRPC_HOST \
    {{.Method.Service.FullName}}/{{.MethodName}}{{end}}
```
