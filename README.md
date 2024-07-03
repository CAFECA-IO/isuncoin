# iSunCoin

This repository is a fork of Ethereum's go-ethereum repository. We expand on v1.11.6 by adding functionalities that enable smart contracts to compute on encrypted data.

[切換至中文](/README_ZH.md)

## Deploy iSunCoin
last updated on 2024-06-18

## Environment
- Ubuntu 22.04

## Step
- [Setup User](https://github.com/CAFECA-IO/KnowledgeManagement/blob/master/linux/create_sudoer_user_in_ubuntu.md)
- [Setup SWAP](https://github.com/CAFECA-IO/KnowledgeManagement/blob/master/linux/setup_swap.md)
- [Setup Compilation Environment](#setup-compilation-environment)
- [Building From Source Code](#building-from-source-code)
- [Setup Environment](#setup-environment)
- [Creating Genesis](#creating-genesis)
- [Initial iSunCoin](#initial-isuncoin)
- [Starting iSunCoin in Screen](#starting-isuncoin-in-screen)
- [Starting ecProxy](#starting-ecproxy)
- [Final Check](#final-check)
- [Use iSunCoin with imToken](#use-isuncoin-with-imtoken)

### Setup Compilation Environment
- Git
- [Install Golang](https://github.com/CAFECA-IO/KnowledgeManagement/blob/master/linux/install_golang.md)

### Building From Source Code
```shell
cd /workspace
git clone https://github.com/CAFECA-IO/isuncoin
cd isuncoin
make isuncoin
```

### Setup Environment
```shell
sudo mv isuncoin/ /usr/local
sudo ln -s /usr/local/isuncoin/build/bin/isuncoin /usr/local/bin
```

### Start iSunCoin
```shell
isuncoin \
--datadir ~/isuncoin \
--mine --miner.threads=1 --miner.etherbase 0xCAFECA05eB2686e2D7e78449F35d8F6D2Faee174 \
--http --http.api eth,net,web3 \
--http.port 8545 --port 30303 --authrpc.port 8551
```

### Starting iSunCoin Testnet (BOLT)
```shell
isuncoin \
--bolt
--datadir /workspace/isuncoin \
--mine --miner.threads=1 --miner.etherbase 0xCAFECA05eB2686e2D7e78449F35d8F6D2Faee174 \
--http --http.api eth,net,web3 \
--http.port 8545 --port 30303 --authrpc.port 8551
```

### Starting ecProxy
```shell
bash <(curl https://raw.githubusercontent.com/Luphia/ecProxy/master/shell/install-lite.sh -kL)
```

### Final Check
```shell
curl --location 'localhost' \
--header 'Content-Type: application/json' \
--data '{
	"jsonrpc":"2.0",
	"method":"eth_blockNumber",
	"params":[],
	"id":1
}'
```

### Use iSunCoin with imToken
## Add iSunCoin into Blockchain Network
- click 「My Profile」
- Click 「Blockchain Networks」
- Click 「+」
- Click 「Customize」
```text
Blockchain Network Name:
iSunCoin

Add RPC URL:
https://isuncoin.baifa.io

Chain ID:
8017

Use as a testnet:
(disable)

Native token symbol:
ISC

Block Explorer URL:
https://baifa.io/app/chains/8017
```
- Click 「Save」

## Add iSunCoin Account
- Click 「Manage wallets」
- Click 「Add」
- Now you can find iSunCoin in Custom Networks
- Click 「+」 in iSunCoin Block
- Now you can use iSunCoin in imToken