---
sidebar_position: 1
---

# LND

Welcome to the API reference documentation for LND, the Lightning Network
Daemon.

The Lightning Network Daemon (`lnd`) - is a complete implementation of a
[Lightning Network](https://lightning.network) node. `lnd` has several pluggable back-end
chain services including [`btcd`](https://github.com/btcsuite/btcd) (a
full-node), [`bitcoind`](https://github.com/bitcoin/bitcoin), and
[`neutrino`](https://github.com/lightninglabs/neutrino) (a new experimental light client). The project's codebase uses the
[btcsuite](https://github.com/btcsuite/) set of Bitcoin libraries, and also
exports a large set of isolated re-usable Lightning Network related libraries
within it. In the current state `lnd` is capable of:

- Creating channels.
- Closing channels.
- Completely managing all channel states (including the exceptional ones!).
- Maintaining a fully authenticated+validated channel graph.
- Performing path finding within the network, passively forwarding incoming payments.
- Sending outgoing [onion-encrypted payments](https://github.com/lightningnetwork/lightning-onion)
  through the network.
- Updating advertised fee schedules.
- Automatic channel management ([`autopilot`](https://github.com/lightningnetwork/lnd/tree/master/autopilot)).

## Usage

Learn how to install, configure, and use LND by viewing the documentation in the [Builder's Guide](https://docs.lightning.engineering/lightning-network-tools/lnd/run-lnd).

## Summary

This site features the API documentation for `{{cliCmd}}` (CLI), [Python](https:///dev.lightning.community/guides/python-grpc/),
and [JavaScript](https://dev.lightning.community/guides/javascript-grpc/) in
order to communicate with a local `{{daemonCmd}}` instance through gRPC and REST. It is intended
for those who already understand how to work with LND. If this is your first
time or you need a refresher, you may consider perusing our LND developer site
featuring a tutorial, resources and guides at [dev.lightning.community](https://dev.lightning.community).

{{> partial_daemons_shared}}

## Experimental services

The following RPCs/services are currently considered to be experimental. This means
they are subject to change in the future. They also need to be enabled with a
compile-time flag to be active (they are active in the official release binaries).

{{#each experimentalServices}}
- [Service _{{name}}_](../../category/{{lowerName}}-service) (file: `{{file}}`)
{{/each}}
