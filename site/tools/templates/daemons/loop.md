---
sidebar_position: 2
---

# Loop

Welcome to the API reference documentation for Lightning Loop.

Lightning Loop is a non-custodial service offered by Lightning Labs to bridge
on-chain and off-chain Bitcoin using submarine swaps. This repository is home to
the Loop client and depends on the Lightning Network daemon loop. All of loop’s
supported chain backends are fully supported when using the Loop client:
Neutrino, Bitcoin Core, and btcd.

The service can be used in various situations:

- Acquiring inbound channel liquidity from arbitrary nodes on the Lightning
  network
- Depositing funds to a Bitcoin on-chain address without closing active
  channels
- Paying to on-chain fallback addresses in the case of insufficient route
  liquidity
- Refilling depleted channels with funds from cold-wallets or exchange
  withdrawals
- Servicing off-chain Lightning withdrawals using on-chain payments, with no
  funds in channels required
- As a failsafe payment method that can be used when channel liquidity along a
  route is insufficient

## Usage

Learn how to install, configure, and use Loop by viewing the documentation in the [Builder's Guide](https://docs.lightning.engineering/lightning-network-tools/loop/get-started).

{{> partial_daemons_shared summary=true}}
