```javascript
const fs = require('fs');
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const loaderOptions = {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
};
const packageDefinition = protoLoader.loadSync({{{loaderFiles}}}, loaderOptions);
const {{ packageName }} = grpc.loadPackageDefinition(packageDefinition).{{ packageName }};
process.env.GRPC_SSL_CIPHER_SUITES = 'HIGH+ECDSA';
const tlsCert = fs.readFileSync('{{upper daemonName}}_DIR/tls.cert');
const sslCreds = grpc.credentials.createSsl(tlsCert);
{{#if requiresMacaroon}}
const macaroon = fs.readFileSync('{{macaroonPath}}').toString('hex');
const macaroonCreds = grpc.credentials.createFromMetadataGenerator(function(args, callback) {
  let metadata = new grpc.Metadata();
  metadata.add('macaroon', macaroon);
  callback(null, metadata);
});
let creds = grpc.credentials.combineChannelCredentials(sslCreds, macaroonCreds);
let client = new {{packageName}}.{{serviceName}}('localhost:{{grpcPort}}', creds);
{{else}}
let client = new {{packageName}}.{{serviceName}}('localhost:{{grpcPort}}', sslCreds);
{{/if}}
{{#if requestFields}}
let request = {
  {{#each requestFields}}
  {{name}}: <{{type}}>,
  {{/each}}
};
{{else}}
let request = {};
{{/if}}
{{#if isUnary}}
client.{{camel method.name}}(request, function(err, response) {
  console.log(response);
});
{{/if}}
{{#if isServerStreaming}}
let call = client.{{camel method.name}}(request);
call.on('data', function(response) {
  // A response was received from the server.
  console.log(response);
});
call.on('status', function(status) {
  // The current status of the stream.
});
call.on('end', function() {
  // The server has closed the stream.
});
{{/if}}
{{#if isBidirectionalStreaming}}
let call = client.{{camel method.name}}({});
call.on('data', function(response) {
  // A response was received from the server.
  console.log(response);
});
call.on('status', function(status) {
  // The current status of the stream.
});
call.on('end', function() {
  // The server has closed the stream.
});
call.write(request);
{{/if}}
// Console output:
//  {
{{#each responseFields}}
//    "{{name}}": <{{type}}>,
{{/each}}
//  }
```
