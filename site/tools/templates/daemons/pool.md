---
sidebar_position: 4
---

# Pool

Welcome to the API reference documentation for Lightning Pool.

Lightning Pool is a non-custodial batched uniform clearing-price auction for
Lightning Channel Lease (LCL). A LCL packages up inbound (or outbound!) channel
liquidity (ability to send/receive funds) as a fixed incoming asset (earning
interest over time) with a maturity date expressed in blocks. The maturity date
of each of the channels is enforced by Bitcoin contracts, ensuring that the
funds of the maker (the party that sold the channel) can't be swept until the
maturity height. All cleared orders (purchased channels) are cleared in a
single batched on-chain transaction.

This repository is home to the Pool client and depends on the Lightning Network
daemon lnd. All of lnd’s supported chain backends are fully supported when
using the Pool client: Neutrino, Bitcoin Core, and btcd.

The service can be used in various situations:

- Bootstrapping new users with side car channels
- Bootstrapping new services to Lightning
- Demand fueled routing node channel selection
- Allowing users to instantly receive with a wallet

## Usage

Learn how to install, configure, and use Pool by viewing the documentation in the [Builder's Guide](https://docs.lightning.engineering/lightning-network-tools/pool/overview).

{{> partial_daemons_shared summary=true}}
