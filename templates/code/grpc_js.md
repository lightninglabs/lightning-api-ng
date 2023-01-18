```javascript
const fs = require('fs');
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

const GRPC_HOST = 'localhost:{{.GrpcPort}}'{{if .RequiresMacaroon}}
const MACAROON_PATH = '{{.MacaroonPath}}'{{end}}
const TLS_PATH = '{{upper .AppName}}_DIR/tls.cert'

const loaderOptions = {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
};
const packageDefinition = protoLoader.loadSync({{.LoaderFiles}}, loaderOptions);
const {{ .PackageName }} = grpc.loadPackageDefinition(packageDefinition).{{ .PackageName }};
process.env.GRPC_SSL_CIPHER_SUITES = 'HIGH+ECDSA';
const tlsCert = fs.readFileSync(TLS_PATH);
const sslCreds = grpc.credentials.createSsl(tlsCert);{{if .RequiresMacaroon}}
const macaroon = fs.readFileSync(MACAROON_PATH).toString('hex');
const macaroonCreds = grpc.credentials.createFromMetadataGenerator(function(args, callback) {
  let metadata = new grpc.Metadata();
  metadata.add('macaroon', macaroon);
  callback(null, metadata);
});
let creds = grpc.credentials.combineChannelCredentials(sslCreds, macaroonCreds);
let client = new {{.PackageName}}.{{.ServiceName}}(GRPC_HOST, creds);{{else}}
let client = new {{.PackageName}}.{{.ServiceName}}(GRPC_HOST, sslCreds);{{end}}{{if .RequestFields}}
let request = {{"{"}}{{range .RequestFields}}
  {{.Name}}: <{{.Type}}>,{{end}}
};{{else}}
let request = {};{{end}}{{if .IsUnary}}
client.{{camel .MethodName}}(request, function(err, response) {
  console.log(response);
});{{end}}{{if .IsServerStreaming}}
let call = client.{{camel .MethodName}}(request);
call.on('data', function(response) {
  // A response was received from the server.
  console.log(response);
});
call.on('status', function(status) {
  // The current status of the stream.
});
call.on('end', function() {
  // The server has closed the stream.
});{{end}}{{if .IsBidirectionalStreaming}}
let call = client.{{camel .MethodName}}({});
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
call.write(request);{{end}}
// Console output:
//  {
{{range .ResponseFields}}//    "{{.Name}}": <{{.Type}}>,
{{end}}//  }
```