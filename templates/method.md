# {{.Name}}

{{if .IsDeprecated}}:::danger

This RPC is deprecated and will be removed in a future version.

:::
{{end}}
{{.Description}}
{{if (ne .Source "")}}
<small>

Source: [{{.Service.FileName}}]({{.Source}})

</small>{{end}}

### gRPC

{{if (ne .StreamingDirection "")}}:::info

This is a {{.StreamingDirection}}-streaming RPC

:::
{{end}}
```
rpc {{.Name}} ({{if .RequestStreaming}}stream {{end}}{{.RequestType}}) returns ({{if .ResponseStreaming}}stream {{end}}{{.ResponseType}});
```

{{if .HasRestMapping}}
### REST

| HTTP Method                   | Path           |
| ----------------------------- | -------------- |
| <Pill> {{.RestMethod}} </Pill> | `{{.RestPath}}` |

{{end}}
## Messages

{{template "request_message.md" multiArgs "Message" .Request "Source" .RequestTypeSource}}
{{template "message.md" multiArgs "Message" .Response "Source" .ResponseTypeSource}}

{{if .HasNestedMessages}}## Nested Messages


{{range .NestedMessages}}{{template "message.md" multiArgs "Message" .}}{{end}}{{end}}

{{if .HasNestedEnums}}## Enums


{{range .NestedEnums}}{{template "enum.md" .}}{{end}}{{end}}