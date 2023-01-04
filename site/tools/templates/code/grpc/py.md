```python
import codecs, grpc, os
# Generate the following 2 modules by compiling the {{protoFileName}} with the grpcio-tools.
# See https://github.com/lightningnetwork/lnd/blob/master/docs/grpc/python.md for instructions.
import {{stubFileName}}_pb2 as {{packageName}}, {{stubFileName}}_pb2_grpc as {{stubFileName}}stub

GRPC_HOST = 'localhost:{{grpcPort}}'
{{#if requiresMacaroon}}
MACAROON_PATH = '{{macaroonPath}}'
{{/if}}
TLS_PATH = '{{upper daemonName}}_DIR/tls.cert'

{{#if requiresMacaroon}}
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
channel = grpc.secure_channel(GRPC_HOST, combined_creds)
{{else}}
# create SSL credentials
os.environ['GRPC_SSL_CIPHER_SUITES'] = 'HIGH+ECDSA'
cert = open(TLS_PATH, 'rb').read()
ssl_creds = grpc.ssl_channel_credentials(cert)
# make the request
channel = grpc.secure_channel(GRPC_HOST, ssl_creds)
{{/if}}
stub = {{stubFileName}}stub.{{serviceName}}Stub(channel)
{{#if method.requestStreaming}}
# Define a generator that returns an Iterable of {{requestName}} objects.
def request_generator():
    # Initialization code here.
    while True:
        # Parameters here can be set as arguments to the generator.
            {{#if requestFields}}
        request = {{packageName}}.{{requestName}}(
                {{#each requestFields}}
            {{name}}=<{{type}}>,
                {{/each}}
        )
            {{else}}
        request = {{packageName}}.{{requestName}}()
            {{/if}}
        yield request
        # Do things between iterations here.
request_iterable = request_generator()
{{else}}
  {{#if requestFields}}
request = {{packageName}}.{{requestName}}(
  {{#each requestFields}}
  {{name}}=<{{type}}>,
  {{/each}}
)
  {{else}}
request = {{packageName}}.{{requestName}}()
  {{/if}}
{{/if}}
{{#if method.responseStreaming}}
for response in stub.{{methodName}}(request{{#if method.requestStreaming}}_iterable{{/if}}):
  print(response)
{{else}}
response = stub.{{methodName}}(request)
print(response)
{{/if}}
# {
  {{#each responseFields}}
#    "{{name}}": <{{type}}>,
  {{/each}}
# }
```
