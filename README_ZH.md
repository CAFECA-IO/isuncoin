# iSunCoin

iSunCoin 是相容以太坊的區塊鏈，基於以太坊 v1.16.0 為基礎，添加了零知識證明的技術，讓 iSunCoin 能在智能合約上實現隱私數據運算。

[Switch to English](/README.md)

## 部署 iSunCoin
last updated on 2024-06-18

## 部署環境
- Ubuntu 22.04

## 步驟
- [建立用戶](https://github.com/CAFECA-IO/KnowledgeManagement/blob/master/linux/create_sudoer_user_in_ubuntu.md)
- [設定 SWAP](https://github.com/CAFECA-IO/KnowledgeManagement/blob/master/linux/setup_swap.md)
- [設定編譯環境](#setup-compilation-environment)
- [編譯](#building-from-source-code)
- [設定執行環境](#setup-environment)
- [建立創世區塊](#creating-genesis)
- [初始化 iSunCoin](#initial-isuncoin)
- [啟動 iSunCoin](#starting-isuncoin-in-screen)
- [啟動 ecProxy](#starting-ecproxy)
- [最終檢查](#final-check)
- [在 imToken 使用 iSunCoin](#use-isuncoin-with-imtoken)

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
--datadir /workspace/chaindata \
--networkid 8017 \
--mine --miner.threads=1 --miner.etherbase 0xCAFECA05eB2686e2D7e78449F35d8F6D2Faee174 \
--http --http.api eth,net,web3 \
--http.port 8545 --port 30303 --authrpc.port 8551
```

### Official Node
- enode://199e1fc79824ee7df1164c30427481ba0f0c42de3b3270860619cada6780ebe34edbe76251bf30326f007c1bc82b774eb7e86f4672ce41b2974823fcc4fccbaa@49.0.255.11:30303


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
- 點擊「我」
- 點擊「區塊鏈網路」
- 點擊右上角的「+」
- 點擊「自定義」
```text
網路名稱:
iSunCoin

新增 RPC 網址:
https://isuncoin.baifa.io

Chain ID:
8017

此網路為測試網:
(關閉)

代幣符號:
ISC

區塊鏈瀏覽器:
https://baifa.io/app/chains/8017
```
- 點擊「保存」

## Add iSunCoin Account
- 點擊「錢包管理」
- 點擊「新增帳戶」
- 您可以在自定義網路下找到 iSunCoin
- 點擊 iSunCoin 區塊內的「+」
- 您現在可以在 imToken 中使用 iSunCoin