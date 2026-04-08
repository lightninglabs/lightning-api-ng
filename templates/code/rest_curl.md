```bash
REST_HOST=localhost:{{.RestPort}}
{{.DirVarName}}={{.DefaultDir}}{{if .RequiresMacaroon}}
NETWORK=mainnet
MACAROON_PATH="{{.ShellMacaroonPath}}"{{end}}
TLS_PATH="{{.ShellTlsPath}}"

curl {{if (ne .Method.RestMethod "GET")}}-X {{.Method.RestMethod}} {{end}}\
    --cacert $TLS_PATH{{if .RequiresMacaroon}} \
    -H "Grpc-Metadata-macaroon: $(xxd -ps -u -c 1000 $MACAROON_PATH)"{{end}}{{if .IsRestPost}} \
    -d '{{.ShellRequestData}}'{{end}} \
    https://$REST_HOST{{.Method.RestPath}}
```
