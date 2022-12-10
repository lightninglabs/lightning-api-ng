# Taro

Welcome to the API reference documentation for Taro.

The Taro Daemon `tarod` implements the [Taro protocol](https://github.com/Roasbeef/bips/blob/bip-taro/bip-taro.mediawiki) for issuing assets on the Bitcoin blockchain. Taro leverages Taproot transactions to commit to newly created assets and their transfers in an efficient and scalable manner. Multiple assets can be created and transferred in a single bitcoin UTXO, while witness data is transacted and kept off-chain.

Features:

- Mint assets
- Send and receive assets
- Export and import Taro proofs
- Create and manage profiles

{{> partial_daemons_shared summary=true}}
