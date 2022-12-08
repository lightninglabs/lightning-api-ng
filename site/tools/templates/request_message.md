{{! This request message has the additional REST Placement column }}

### {{fullName}}

{{#if fields}}
| Field | gRPC Type | REST Type | REST Placement |
| ----- | --------- | --------- | -------------- |
{{/if}}
{{#fields}}
| <MessageField name="{{name}}">{{description}}</MessageField> | {{{linkedType}}} | `TODO` | TODO |
{{/fields}}
{{^fields}}
:::note

This request has no parameters.

:::
{{/fields}}