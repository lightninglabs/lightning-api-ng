<Tabs groupId="protocol">
<TabItem value="grpc" label="gRPC">

<Tabs groupId="code-samples">
<TabItem value="js" label="Javascript">

{{> partial_code_grpc_js}}

</TabItem>
<TabItem value="py" label="Python">

{{> partial_code_grpc_py}}

</TabItem>
</Tabs>

</TabItem>
{{#if method.restPath}}
<TabItem value="rest" label="REST">

<Tabs groupId="code-samples">
<TabItem value="js" label="Javascript">

{{> partial_code_rest_js}}

</TabItem>
<TabItem value="py" label="Python">

{{> partial_code_rest_py}}

</TabItem>
</Tabs>

</TabItem>
{{/if}}
<TabItem value="bash" label="Shell">

```bash
{{#if method.commandLine}}
$ {{method.commandLine}} --help

{{method.commandLineHelp}}
{{else}}
# There is no CLI command for this RPC
{{/if}}
```

</TabItem>
</Tabs>
