```python
import base64, codecs, json, requests

REST_HOST = 'localhost:{{.RestPort}}'{{if .RequiresMacaroon}}
MACAROON_PATH = '{{.MacaroonPath}}'{{end}}
TLS_PATH = '{{upper .AppName}}_DIR/tls.cert'

url = f'https://{REST_HOST}{{.Method.RestPath}}'{{if .RequiresMacaroon}}
macaroon = codecs.encode(open(MACAROON_PATH, 'rb').read(), 'hex')
headers = {'Grpc-Metadata-macaroon': macaroon}{{end}}{{if .IsRestPost}}
data = {{"{"}}{{range .RequestFields}}
  '{{.Name}}': {{if .Encoding}}base64.b64encode(<{{.Type}}>){{else}}<{{.Type}}>{{end}},{{end}}
}{{end}}
r = requests.{{lower .Method.RestMethod}}(url, {{.PythonRestArgs}}verify=TLS_PATH){{if .Method.ResponseStreaming}}
for raw_response in r.iter_lines():
  json_response = json.loads(raw_response)
  print(json_response){{else}}
print(r.json()){{end}}
# {{"{"}}{{range .ResponseFields}}
#    "{{.Name}}": <{{.Type}}>,{{end}}
# }
```