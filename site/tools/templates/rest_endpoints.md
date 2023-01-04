---
sidebar_position: 1
hide_table_of_contents: true
---

# REST Endpoints

Here's a list of all the REST API endpoints available in {{pascalName}}.

| REST Path | REST Method | RPC |
| --------- | ----------- | --- |
{{#each restEndpoints}}
| [{{restPath}}]({{linkUrl}}) | `{{restMethod}}` | {{methodName}} |
{{/each}}
