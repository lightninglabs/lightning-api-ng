### {{.Message.FullName}}
{{if .Source}}
<small>

Source: [{{.Message.FileName}}]({{.Source}})

</small>
{{end}}{{if .Message.HasFields}}
| Field | gRPC Type | REST Type |
| ----- | --------- | --------- |
{{range .Message.Fields}}| <MessageField name="{{.Name}}">{{.Description}}</MessageField> | {{.LinkedType}} | `{{.RestType}}` |
{{end}}{{else}}
:::note

This response has no parameters.

:::
{{end}}