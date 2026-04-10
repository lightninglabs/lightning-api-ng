<Tabs groupId="protocol">
<TabItem value="grpc" label="gRPC">

<Tabs groupId="code-samples">
<TabItem value="js" label="Javascript">

{{template "grpc_js.md" .}}

</TabItem>
<TabItem value="py" label="Python">

{{template "grpc_py.md" .}}

</TabItem>
<TabItem value="shell" label="grpcurl">

{{template "grpc_grpcurl.md" .}}

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
<TabItem value="shell" label="curl">

{{template "rest_curl.md" .}}

</TabItem>
</Tabs>

</TabItem>{{end}}{{if (ne .Method.CommandLine "")}}
<TabItem value="bash" label="{{.CliCommand}}">

```bash
$ {{.Method.CommandLine}} --help

{{.Method.CommandLineHelp}}
```

</TabItem>{{end}}
</Tabs>
