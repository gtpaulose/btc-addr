# BTC Address Generator

CLI client created using `go-cobra` to generate BIP32 mnemonics, associated addresses (BIP44, BIP49, BIP84) and multi-sig addresses.
Test Coverage: 78%

### Project Structure
```
.
├── README.md
├── cli
│   ├── generate.go
│   ├── list.go
│   ├── mnemonic.go
│   ├── multisig.go
│   └── root.go
├── cmd
│   └── btc-addr
│       └── main.go
├── go.mod
├── go.sum
└── internal
    ├── keystore
    │   ├── address.go
    │   ├── key.go
    │   ├── keystore.go
    │   └── keystore_test.go
    └── multisig
        ├── multisig.go
        └── multisig_test.go

6 directories, 15 files
```
## Installation

```bash
go get github.com/gtpaulose/btc-addr/cmd/btc-addr
```

## Usage

### Generate mnemonics
```bash
$ btc-addr generate mnemonic -p secret
Mnemonic:  over virus glove social tube throw mandate among arch copper include evoke
PassPhrase: secret
Seed:      28fb9ba65d3af8732214cb4e8a6bf5d671c7b2c7fd87decfa777d4f5c58d3ed8cd856621dae448f831cefabbf975678f1069d890aa61c81f28ff8cbd59e68151
Root Key (BIP32): xprv9s21ZrQH143K4HJJpqWizA8Lvt9pchcid8HuDjNdquQFErF8c55vU6nLLMpSgbLVfWffYzLb1oQRMbJijGg6d6LmnJPxNX9yZwDCqpKHnJD
```

### List addresses
```bash
$ btc-addr list -m "over virus glove social tube throw mandate among arch copper include evoke" -n 5     
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Path               Address                            Public Key                                                         Private Key
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------
m/44'/0'/0'/0/0    15KpzgkdS2eDzdbc7piU54eKR6pZL9dakj 031b62689ae7ef38bb0b2e5f5e6536b209b3e524600e50804727674c1516f064e9 Kxqnj8oqfE1WUhHtQqSQxgRusFdBuP57KN8rHLdnHQoHqWaW8G3i
m/44'/0'/0'/0/1    158CYWVfHJd2WnS5AACvS1U7aAMozEaCu2 02d6ff01563d129a292775bafac94fed9f73cf93d5cf8c38a4d855677c81f40c84 Ky96EeFSYW1X15VbXQzN1pKk9KkvRCCut7Pjtz7f24SzrajyqVCb
m/44'/0'/0'/0/2    1Ei2i8fZVx1Y7Qeh5vJbVCUNruTwvbWZQC 02c990ee100d613129e80feb726a3e29cc5e850541a45552ea19cc1f187913c9b0 KwGhXAdFRbL4mLk2X6U2SCAeJQ1zkJs3PnWNBBFaEDJdfLUHABUd
m/44'/0'/0'/0/3    1LDnbN8wTy2KQRnmEUczvu6Cwgpt2EMQeK 03d27a8469a75061f60f80fd9a27c6e523ae982f339b0d9f934b3c7b5e4f275946 KxKTz8tu6ASC8FXSVDUzx6V4wNdXhQAag9NrNmdWFBHXgEb5Rts4
m/44'/0'/0'/0/4    1KCgDANXryS8RQ1jW42xR2fTEtQBVRcnKr 0265b73d04e45a94b0303abba8838795adbe4b61290129b6bd0f0b053ae976184c L1ULx9faVXtvJZ4GaMxtsjLSVyJiHzFyX9nKC39FZwxBSN4mUv23
```

### List addresses for given path and format
```bash
$ btc-addr list -m "over virus glove social tube throw mandate among arch copper include evoke" -f bip49 -p "m/49'/0'/0'/0/6"
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Path               Address                            Public Key                                                         Private Key
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------
m/49'/0'/0'/0/6    36ZVcExdYMjoPrm8r1USHEsCfGDyxNpryN 03515588f90b2ebb8a4755597e82bafdd12148d384eba8bd152da423e23dc6a58e KwM5GGaFi1NVdm5ViGHFKvXLf8PbESyeLZusH77eCDoGfk4CGzXb
```

**Note:** Addresses corresponding to `-f bip49` and `-f bip84` are segwit nested and segwit(bech32) respectively

### Create N-of-M multi-sig address
```bash
$ btc-addr generate multisig --n 2 --m 3 020f8796e0f870a9a3b269be3b1e78e380c9b569885f0de98a9ff061c4a66e79d2 02dfa8990f3f015ff20e9b31b85ea36d47470220615fb2ac1597e20fc830727b25 03fbfbdc5df9c60e4b747805552686199e85299a5e87804dbb66a14597ddabcf29

Address:  38GxnyunFjXkHQMUEU5m4gqhD1fmKcRuEd
```

For more information and examples use `btc-addr --help` or `btc-addr [COMMAND] --help`

## Third Party Packages
1. [Cobra CLI](https://github.com/spf13/cobra) - CLI commander written in Go
2. [github.com/btcsuite/btcd](https://github.com/btcsuite/btcd/tree/master/btcec) - for BTC mainet config and script building utilities
3. [github.com/spf13/cast](https://github.com/spf13/cast) - safely cast primitive types
4. [github.com/stretchr/testify](https://github.com/stretchr/testify) - test assertion library
5. [github.com/tyler-smith/go-bip32](https://github.com/tyler-smith/go-bip32) & [github.com/tyler-smith/go-bip39](https://github.com/tyler-smith/go-bip39) - libraries used to generate mnemonics and bip32 keys

## Notes
All addresses and mnemonics created by the application can be tested using [https://iancoleman.io/bip39/](https://iancoleman.io/bip39/). When passing unverifiable inputs for mnemonics and paths, the client may generate inaccurate addresses. Make sure all mnemonics passed to the client are either generated using the application or conform to the BIP39 standards. 

The client also runs on your local machine and doesn't save any addresses to a state. Hence `stdout` will be the only place where your keys and mnemonics will be seen. The client runs offline. 