### {{fullName}}

{{#if fields}}
| Field | gRPC Type | REST Type |
| ----- | --------- | --------- |
{{/if}}
{{#fields}}
| <MessageField name="{{name}}">{{description}}</MessageField> | {{{linkedType}}} | `{{restType}}` |
{{/fields}}
{{^fields}}
:::note

This response has no parameters.

:::
{{/fields}}
