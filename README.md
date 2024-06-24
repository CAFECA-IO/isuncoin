# iSunCoin

This repository is a fork of Ethereum's go-ethereum repository. We expand on v1.16.0 by adding functionalities that enable smart contracts to compute on encrypted data.

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
make geth
```

### Setup Environment
```shell
sudo mv isuncoin/ /usr/local
sudo ln -s /usr/local/isuncoin/build/bin/geth /usr/local/bin
```

### Creating Genesis
```shell
mkdir ~/isuncoin
cd ~/isuncoin
vi genesis.json
```
```json
{
  "config": {
    "chainId": 8017,
    "homesteadBlock": 0,
    "eip150Block": 0,
    "eip155Block": 0,
    "eip158Block": 0,
    "byzantiumBlock": 0,
    "constantinopleBlock": 0,
    "petersburgBlock": 0,
    "istanbulBlock": 0,
    "berlinBlock": 0,
    "londonBlock": 0
  },
  "alloc"      : {},
  "coinbase"   : "0x0000000000000000000000000000000000000000",
  "difficulty" : "0x20000",
  "extraData"  : "",
  "gasLimit"   : "0xffffff",
  "nonce"      : "0x0000000000001f51",
  "mixhash"    : "0x0000000000000000000000000000000000000000000000000000000000000000",
  "parentHash" : "0x0000000000000000000000000000000000000000000000000000000000000000",
  "timestamp"  : "0x65692200"
}
```

### Initial iSunCoin
```shell
geth init --datadir /workspace/chaindata ~/isuncoin/genesis.json
```

### Starting iSunCoin in Screen
- Will Single Command
```shell
geth \
--datadir /workspace/ \
--networkid 8017 \
--mine --miner.threads=1 --miner.etherbase 0xCAFECA05eB2686e2D7e78449F35d8F6D2Faee174 \
--http --http.api eth,net,web3 \
--http.port 8545 --port 30303 --authrpc.port 8551
```

### Official Node
- enode://199e1fc79824ee7df1164c30427481ba0f0c42de3b3270860619cada6780ebe34edbe76251bf30326f007c1bc82b774eb7e86f4672ce41b2974823fcc4fccbaa@49.0.255.11:30303
```shell
geth --bootnodes "enode://199e1fc79824ee7df1164c30427481ba0f0c42de3b3270860619cada6780ebe34edbe76251bf30326f007c1bc82b774eb7e86f4672ce41b2974823fcc4fccbaa@49.0.255.11:30303"
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
	"id":83
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