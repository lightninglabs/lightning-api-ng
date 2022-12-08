{{! This request message has the additional REST Placement column }}

### {{fullName}}

| Field | gRPC Type | REST Type | REST Placement |
| ----- | --------- | --------- | -------------- |
{{#fields}}
| <MessageField name="{{name}}">{{description}}</MessageField> | {{{linkedType}}} | `TODO` | TODO |
{{/fields}}