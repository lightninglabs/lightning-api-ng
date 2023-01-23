```python
import codecs, grpc, os
# Generate the following 2 modules by compiling the {{.ProtoFileName}} with the grpcio-tools.
# See https://github.com/lightningnetwork/lnd/blob/master/docs/grpc/python.md for instructions.
import {{.StubFileName}}_pb2 as {{.PackageName}}, {{.StubFileName}}_pb2_grpc as {{.StubFileName}}stub

GRPC_HOST = 'localhost:{{.GrpcPort}}'{{if .RequiresMacaroon}}
MACAROON_PATH = '{{.MacaroonPath}}'{{end}}
TLS_PATH = '{{upper .AppName}}_DIR/tls.cert'
{{if .RequiresMacaroon}}
# create macaroon credentials
macaroon = codecs.encode(open(MACAROON_PATH, 'rb').read(), 'hex')
def metadata_callback(context, callback):
  callback([('macaroon', macaroon)], None)
auth_creds = grpc.metadata_call_credentials(metadata_callback)
# create SSL credentials
os.environ['GRPC_SSL_CIPHER_SUITES'] = 'HIGH+ECDSA'
cert = open(TLS_PATH, 'rb').read()
ssl_creds = grpc.ssl_channel_credentials(cert)
# combine macaroon and SSL credentials
combined_creds = grpc.composite_channel_credentials(ssl_creds, auth_creds)
# make the request
channel = grpc.secure_channel(GRPC_HOST, combined_creds){{else}}
# create SSL credentials
os.environ['GRPC_SSL_CIPHER_SUITES'] = 'HIGH+ECDSA'
cert = open(TLS_PATH, 'rb').read()
ssl_creds = grpc.ssl_channel_credentials(cert)
# make the request
channel = grpc.secure_channel(GRPC_HOST, ssl_creds){{end}}
stub = {{.StubFileName}}stub.{{.ServiceName}}Stub(channel){{if .Method.RequestStreaming}}
# Define a generator that returns an Iterable of {{.RequestName}} objects.
def request_generator():
    # Initialization code here.
    while True:
        # Parameters here can be set as arguments to the generator.{{if .RequestFields}}
        request = {{.PackageName}}.{{.RequestName}}({{range .RequestFields}}
            {{.Name}}=<{{.Type}}>,{{end}}
        ){{else}}
        request = {{.PackageName}}.{{.RequestName}}(){{end}}
        yield request
        # Do things between iterations here.
request_iterable = request_generator(){{else}}{{if .RequestFields}}
request = {{.PackageName}}.{{.RequestName}}({{range .RequestFields}}
  {{.Name}}=<{{.Type}}>,{{end}}
){{else}}
request = {{.PackageName}}.{{.RequestName}}(){{end}}{{end}}{{if .Method.ResponseStreaming}}
for response in stub.{{.MethodName}}(request{{if .Method.RequestStreaming}}_iterable{{end}}):
  print(response){{else}}
response = stub.{{.MethodName}}(request)
print(response){{end}}
# {{"{"}}{{range .ResponseFields}}
#    "{{.Name}}": <{{.Type}}>,{{end}}
# }
```