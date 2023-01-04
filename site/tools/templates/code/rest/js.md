```javascript
const fs = require('fs');
const request = require('request');

const REST_HOST = 'localhost:{{restPort}}'
{{#if requiresMacaroon}}
const MACAROON_PATH = '{{macaroonPath}}'
{{/if}}

{{#if isRestPost}}
let requestBody = {
  {{#each requestFields}}
  {{name}}: <{{restType}}>, // <{{type}}> {{#if encoding}}({{encoding}} encoded){{/if}}
  {{/each}}
};
{{/if}}
let options = {
  url: `https://${REST_HOST}{{ method.restPath }}`,
  // Work-around for self-signed certificates.
  rejectUnauthorized: false,
  json: true,
  {{#if requiresMacaroon}}
  headers: {
    'Grpc-Metadata-macaroon': fs.readFileSync(MACAROON_PATH).toString('hex'),
  },
  {{/if}}
  {{#if isRestPost}}
  form: JSON.stringify(requestBody),
  {{/if}}
}
request.{{lower method.restMethod}}(options, function(error, response, body) {
  console.log(body);
});
// Console output:
//  {
{{#each responseFields}}
//    "{{name}}": <{{restType}}>, // <{{type}}> {{#if encoding}}({{encoding}} encoded){{/if}}
{{/each}}
//  }
{{#if isStreaming}}



// --------------------------
// Example with websockets:
// --------------------------
const WebSocket = require('ws');
const fs = require('fs');

const REST_HOST = 'localhost:{{restPort}}'
{{#if requiresMacaroon}}
const MACAROON_PATH = '{{macaroonPath}}'
{{/if}}

let ws = new WebSocket(`wss://${REST_HOST}{{ method.restPath }}?method={{ method.restMethod }}`, {
  // Work-around for self-signed certificates.
  rejectUnauthorized: false,
  {{#if requiresMacaroon}}
  headers: {
    'Grpc-Metadata-Macaroon': fs.readFileSync(MACAROON_PATH).toString('hex'),
  },
  {{/if}}
});
let requestBody = {
  {{#each requestFields}}
  {{name}}: <{{type}}>, // <{{type}}> {{#if encoding}}({{encoding}} encoded){{/if}}
  {{/each}}
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
//  {
{{#each responseFields}}
//    "{{name}}": <{{restType}}>, // <{{type}}>
{{/each}}
//  }
{{/if}}
```
