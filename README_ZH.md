# iSunCoin

iSunCoin 是相容以太坊的區塊鏈，基於以太坊 v1.11.6 為基礎，添加了零知識證明的技術，讓 iSunCoin 能在智能合約上實現隱私數據運算。

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
make isuncoin
```

### Building Windows EXE
```shell
GOOS=windows GOARCH=amd64 go build -o isuncoin.exe ./cmd/isuncoin
```

### Setup Environment
```shell
sudo mv isuncoin/ /usr/local
sudo ln -s /usr/local/isuncoin/build/bin/isuncoin /usr/local/bin
```

### Start iSunCoin
```shell
isuncoin \
--datadir /workspace/isuncoin \
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