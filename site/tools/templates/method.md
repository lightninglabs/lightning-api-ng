# {{name}}

{{{description}}}

### gRPC

{{#if streamingDirection}}
:::info

This is a {{streamingDirection}}-streaming RPC

:::
{{/if}}

```
rpc {{name}} ({{#if requestStreaming}}stream {{/if}}{{requestType}}) returns ({{#if responseStreaming}}stream {{/if}}{{responseType}});
```

### REST

| HTTP Method       | Path   |
| ----------------- | ------ |
| <Pill>TODO</Pill> | `TODO` |

## Messages

{{#request}}
{{> partial_request_message}}
{{/request}}

{{#response}}
{{> partial_message}}
{{/response}}

## Nested Messages

{{#nestedMessages}}
{{> partial_message}}
{{/nestedMessages}}
