```python
import base64, codecs, json, requests

REST_HOST = 'localhost:{{restPort}}'
{{#if requiresMacaroon}}
MACAROON_PATH = '{{macaroonPath}}'
{{/if}}
TLS_PATH = '{{upper daemonName}}_DIR/tls.cert'

url = f'https://{REST_HOST}{{method.restPath}}'
{{#if requiresMacaroon}}
macaroon = codecs.encode(open(MACAROON_PATH, 'rb').read(), 'hex')
headers = {'Grpc-Metadata-macaroon': macaroon}
{{/if}}
{{#if isRestPost}}
data = {
  {{#each requestFields}}
  '{{name}}': {{#if encoding}}base64.b64encode(<{{type}}>){{else}}<{{type}}>{{/if}},
  {{/each}}
}
{{/if}}
r = requests.{{lower method.restMethod}}(url, {{{pythonRestArgs}}}verify=TLS_PATH)
{{#if method.responseStreaming}}
for raw_response in r.iter_lines():
  json_response = json.loads(raw_response)
  print(json_response)
{{else}}
print(r.json())
{{/if}}
# {
  {{#each responseFields}}
#    "{{name}}": <{{type}}>,
  {{/each}}
# }
```
