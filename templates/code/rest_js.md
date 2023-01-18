```javascript
const fs = require('fs');
const request = require('request');

const REST_HOST = 'localhost:{{.RestPort}}'{{if .RequiresMacaroon}}
const MACAROON_PATH = '{{.MacaroonPath}}'{{end}}
{{if .IsRestPost}}
let requestBody = {{"{"}}{{range .RequestFields}}
  {{.Name}}: <{{.RestType}}>, // <{{.Type}}> {{if .Encoding}}({{.Encoding}} encoded){{end}}{{end}}
};{{end}}
let options = {
  url: `https://${REST_HOST}{{.Method.RestPath }}`,
  // Work-around for self-signed certificates.
  rejectUnauthorized: false,
  json: true,{{if .RequiresMacaroon}}
  headers: {
    'Grpc-Metadata-macaroon': fs.readFileSync(MACAROON_PATH).toString('hex'),
  },{{end}}{{if .IsRestPost}}
  form: JSON.stringify(requestBody),{{end}}
}
request.{{lower .Method.RestMethod}}(options, function(error, response, body) {
  console.log(body);
});
// Console output:
//  {{"{"}}{{range .ResponseFields}}
//    "{{.Name}}": <{{.RestType}}>, // <{{.Type}}> {{end}}
//  }{{if .IsStreaming}}



// --------------------------
// Example with websockets:
// --------------------------
const WebSocket = require('ws');
const fs = require('fs');

const REST_HOST = 'localhost:{{.RestPort}}'{{if .RequiresMacaroon}}
const MACAROON_PATH = '{{.MacaroonPath}}'{{end}}

let ws = new WebSocket(`wss://${REST_HOST}{{ .Method.RestPath }}?method={{ .Method.RestMethod }}`, {
  // Work-around for self-signed certificates.
  rejectUnauthorized: false,{{if .RequiresMacaroon}}
  headers: {
    'Grpc-Metadata-Macaroon': fs.readFileSync(MACAROON_PATH).toString('hex'),
  },{{end}}
});
let requestBody = {{"{"}}{{range .RequestFields}}
  {{.Name}}: <{{.Type}}>, // <{{.Type}}> {{if .Encoding}}({{.Encoding}} encoded){{end}}{{end}}
};
ws.on('open', function() {
    ws.send(JSON.stringify(requestBody));
});
ws.on('error', function(err) {
    console.log('Error: ' + err);
});
ws.on('message', function(body) {
    console.log(body);
});
// Console output:
//  {{"{"}}{{range .ResponseFields}}
//    "{{.Name}}": <{{.RestType}}>, // <{{.Type}}>{{end}}
//  }{{end}}
```