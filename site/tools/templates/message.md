### {{name}}

{{#if fields}}
| Field | gRPC Type | REST Type |
| ----- | --------- | --------- |
{{/if}}
{{#fields}}
| <MessageField name="{{name}}">{{description}}</MessageField> | `{{grpcType}}` | `TODO` |
{{/fields}}
{{^fields}}
:::note

This response has no parameters.

:::
{{/fields}}
