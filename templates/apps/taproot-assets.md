---
sidebar_position: 5
---

# Taproot Assets Protocol

Welcome to the API reference documentation for Taproot Assets Daemon.

The Taproot Assets Daemon `tapd` implements the [Taproot Assets Protocol](https://github.com/lightninglabs/taproot-assets/blob/master/docs/bip-tap.mediawiki) for issuing assets on the Bitcoin blockchain. Taproot Assets leverages Taproot transactions to commit to newly created assets and their transfers in an efficient and scalable manner. Multiple assets can be created and transferred in a single bitcoin UTXO, while witness data is transacted and kept off-chain.

Features:

- Mint assets
- Synchronize to universes
- Send and receive assets
- Export and import Taproot Asset proofs
- Create and manage CLI profiles

## Usage

Learn how to install, configure, and use Taproot Assets Daemon by viewing the documentation in the [Builder's Guide](https://docs.lightning.engineering/the-lightning-network/taproot-assets).

{{template "shared.md" .}}
