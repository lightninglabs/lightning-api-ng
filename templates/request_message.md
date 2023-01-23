{{/* This request message has the additional REST Placement column */}}
### {{.Message.FullName}}

{{if (ne .Source "")}}<small>

Source: [{{.Message.FileName}}]({{.Source}})

</small>
{{end}}{{if .Message.HasFields}}
| Field | gRPC Type | REST Type | REST Placement |
| ----- | --------- | --------- | -------------- |
{{range .Message.Fields}}| <MessageField name="{{.Name}}">{{.Description}}</MessageField> | {{.LinkedType}} | `{{.RestType}}` | `{{.RestPlacement}}` {{.EncodingTip}} |
{{end}}{{else}}
:::note

This request has no parameters.

:::
{{end}}