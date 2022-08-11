# LND API Reference

:::info

This doc was copied from the existing API docs and needs to be tweaked

:::

## gRPC

Welcome to the gRPC API reference documentation for LND, the Lightning Network
Daemon.

This site features the API documentation for lncli (CLI), [Python](https:///dev.lightning.community/guides/python-grpc/),
and [JavaScript](https://dev.lightning.community/guides/javascript-grpc/) in
order to communicate with a local `lnd` instance through gRPC. It is intended
for those who already understand how to work with LND. If this is your first
time or you need a refresher, you may consider perusing our LND developer site
featuring a tutorial, resources and guides at [dev.lightning.community](https://dev.lightning.community).

The examples to the right assume that the there is a local `lnd` instance
running and listening for gRPC connections on port 10009. `LND_DIR` will be used
as a placeholder to denote the base directory of the `lnd` instance. By default,
this is `~/.lnd` on Linux and `~/Library/Application Support/Lnd` on macOS.

At the time of writing this documentation, two things are needed in order to
make a gRPC request to an `lnd` instance: a TLS/SSL connection and a macaroon
used for RPC authentication. The examples to the right will show how these can
be used in order to make a successful, secure, and authenticated gRPC request.

The original `*.proto` files from which the gRPC documentation was generated
can be found here:

- [`autopilotrpc/autopilot.proto`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/autopilotrpc/autopilot.proto)
- [`chainrpc/chainnotifier.proto`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/chainrpc/chainnotifier.proto)
- [`devrpc/dev.proto`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/devrpc/dev.proto)
- [`invoicesrpc/invoices.proto`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/invoicesrpc/invoices.proto)
- [`lightning.proto`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/lightning.proto)
- [`lnclipb/lncli.proto`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/lnclipb/lncli.proto)
- [`neutrinorpc/neutrino.proto`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/neutrinorpc/neutrino.proto)
- [`peersrpc/peers.proto`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/peersrpc/peers.proto)
- [`routerrpc/router.proto`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/routerrpc/router.proto)
- [`signrpc/signer.proto`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/signrpc/signer.proto)
- [`stateservice.proto`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/stateservice.proto)
- [`verrpc/verrpc.proto`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/verrpc/verrpc.proto)
- [`walletrpc/walletkit.proto`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/walletrpc/walletkit.proto)
- [`walletunlocker.proto`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/walletunlocker.proto)
- [`watchtowerrpc/watchtower.proto`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/watchtowerrpc/watchtower.proto)
- [`wtclientrpc/wtclient.proto`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/wtclientrpc/wtclient.proto)

This is the reference for the **gRPC API**. Alternatively, there is also a [REST
API which is documented here](#lnd-rest-api-reference).

<small>This documentation was
[generated automatically](https://github.com/lightninglabs/lightning-api) against commit
[`115b0418747ba941875bd9f2be1bb18754f6f636`](https://github.com/lightningnetwork/lnd/tree/115b0418747ba941875bd9f2be1bb18754f6f636).</small>

### Experimental services

The following RPCs/services are currently considered to be experimental. This means
they are subject to change in the future. They also need to be enabled with a
compile-time flag to be active (they are active in the official release binaries).

- [Service _ChainNotifier_](#service-chainnotifier) (file `chainrpc/chainnotifier.proto`)
- [Service _Invoices_](#service-invoices) (file `invoicesrpc/invoices.proto`)
- [Service _Signer_](#service-signer) (file `signrpc/signer.proto`)
- [Service _WalletKit_](#service-walletkit) (file `walletrpc/walletkit.proto`)
- [Service _Watchtower_](#service-watchtower) (file `watchtowerrpc/watchtower.proto`)

## REST

Welcome to the REST API reference documentation for LND, the Lightning Network
Daemon.

This site features the API documentation for Python and JavaScript, along with
barebones examples using `curl`, for HTTP requests. It is intended for those who
already understand how to work with LND. If this is your first time or you need
a refresher, you may consider perusing our LND developer site featuring a
tutorial, resources and guides at [dev.lightning.community](https://dev.lightning.community).

The examples to the right assume that the there is a local `lnd` instance
running and listening for REST connections on port 8080. `LND_DIR` will be used
as a placeholder to denote the base directory of the `lnd` instance. By default,
this is `~/.lnd` on Linux and `~/Library/Application Support/Lnd` on macOS.

At the time of writing this documentation, two things are needed in order to
make an HTTP request to an `lnd` instance: a TLS/SSL connection and a macaroon
used for RPC authentication. The examples to the right will show how these can
be used in order to make a successful, secure, and authenticated HTTP request.

The original `*.swagger.js` files from which the gRPC documentation was generated
can be found here:

- [`lightning.swagger.json`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/lightning.swagger.json)
- [`stateservice.swagger.json`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/stateservice.swagger.json)
- [`walletunlocker.swagger.json`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/walletunlocker.swagger.json)
- [`autopilotrpc/autopilot.swagger.json`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/autopilotrpc/autopilot.swagger.json)
- [`chainrpc/chainnotifier.swagger.json`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/chainrpc/chainnotifier.swagger.json)
- [`devrpc/dev.swagger.json`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/devrpc/dev.swagger.json)
- [`invoicesrpc/invoices.swagger.json`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/invoicesrpc/invoices.swagger.json)
- [`lnclipb/lncli.swagger.json`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/lnclipb/lncli.swagger.json)
- [`neutrinorpc/neutrino.swagger.json`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/neutrinorpc/neutrino.swagger.json)
- [`peersrpc/peers.swagger.json`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/peersrpc/peers.swagger.json)
- [`routerrpc/router.swagger.json`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/routerrpc/router.swagger.json)
- [`signrpc/signer.swagger.json`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/signrpc/signer.swagger.json)
- [`verrpc/verrpc.swagger.json`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/verrpc/verrpc.swagger.json)
- [`walletrpc/walletkit.swagger.json`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/walletrpc/walletkit.swagger.json)
- [`watchtowerrpc/watchtower.swagger.json`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/watchtowerrpc/watchtower.swagger.json)
- [`wtclientrpc/wtclient.swagger.json`](https://github.com/lightningnetwork/lnd/blob/115b0418747ba941875bd9f2be1bb18754f6f636/lnrpc/wtclientrpc/wtclient.swagger.json)

**NOTE**: The `byte` field type must be set as the base64 encoded string
representation of a raw byte array. Also, any time this must be used in a URL path
(ie. `/v1/abc/xyz/{payment_hash}`) the base64 string must be encoded using a
[URL and Filename Safe Alphabet](https://tools.ietf.org/html/rfc4648#section-5). This means you must replace `+` with `-`,
`/` with `_`, and keep the trailing `=` as is. Url encoding (ie. `%2F`) will not work.

This is the reference for the **REST API**. Alternatively, there is also a [gRPC
API which is documented here](#lnd-grpc-api-reference).

<small>This documentation was
[generated automatically](https://github.com/lightninglabs/lightning-api) against commit
[`115b0418747ba941875bd9f2be1bb18754f6f636`](https://github.com/lightningnetwork/lnd/tree/115b0418747ba941875bd9f2be1bb18754f6f636).</small>
