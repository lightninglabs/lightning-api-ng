---
sidebar_position: 1
hide_table_of_contents: true
---

# REST Endpoints

Here's a list of all the REST API endpoints available in {{.PascalName}}.

| REST Path | REST Method | RPC |
| --------- | ----------- | --- |{{range .RestEndpoints}}
| [{{.RestPath}}]({{.LinkUrl}}) | `{{.RestMethod}}` | {{.MethodName}} |{{end}}
