### {{.FullName}}

| Name | Number |
| ---- | ------ |
{{range .Values}}| <MessageField name="{{.Name}}">{{.Description}}</MessageField> | `{{.Number}}` |
{{end}}