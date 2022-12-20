```javascript
const fs = require('fs');
const request = require('request');
{{#if requiresMacaroon}}
const macaroon = fs.readFileSync('{{macaroonPath}}').toString('hex');
{{/if}}
{{#if isRestPost}}
let requestBody = {
  {{#each requestFields}}
  {{name}}: <{{restType}}>, // <{{type}}> {{#if encoding}}({{encoding}} encoded){{/if}}
  {{/each}}
};
{{/if}}
let options = {
  url: 'https://localhost:{{ restPort }}{{ method.restPath }}',
  // Work-around for self-signed certificates.
  rejectUnauthorized: false,
  json: true,
  {{#if requiresMacaroon}}
  headers: {
    'Grpc-Metadata-macaroon': macaroon,
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
const macaroon = fs.readFileSync('{{macaroonPath}}').toString('hex');
let ws = new WebSocket('wss://localhost:{{ restPort }}{{ method.restPath }}?method={{ method.restMethod }}', {
  // Work-around for self-signed certificates.
  rejectUnauthorized: false,
  headers: {
    'Grpc-Metadata-Macaroon': macaroon,
  },
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
