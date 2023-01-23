<Tabs groupId="protocol">
<TabItem value="grpc" label="gRPC">

<Tabs groupId="code-samples">
<TabItem value="js" label="Javascript">

{{template "grpc_js.md" .}}

</TabItem>
<TabItem value="py" label="Python">

{{template "grpc_py.md" .}}

</TabItem>
</Tabs>

</TabItem>{{if (ne .Method.RestPath "")}}
<TabItem value="rest" label="REST">

<Tabs groupId="code-samples">
<TabItem value="js" label="Javascript">

{{template "rest_js.md" .}}

</TabItem>
<TabItem value="py" label="Python">

{{template "rest_py.md" .}}

</TabItem>
</Tabs>

</TabItem>{{end}}
<TabItem value="bash" label="Shell">

```bash{{if (ne .Method.CommandLine "")}}
$ {{.Method.CommandLine}} --help

{{.Method.CommandLineHelp}}{{else}}
# There is no CLI command for this RPC{{end}}
```

</TabItem>
</Tabs>