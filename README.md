# PBC暗号によるストレージ監査システム
## 概要
ストレージへのアップロードデータとそのログをブロックチェーンに残して，それをPBC暗号を用いた第三者機関による監査システム

## 環境構築
### Solidity -> Golang
```
cd ethereum/
solc --optimize --abi ./solidities/FIT.sol -o build --overwrite
solc --optimize --bin ./solidities/FIT.sol -o build --overwrite
abigen --abi=./build/FileIndexTable.abi --bin=./build/FileIndexTable.bin --pkg=contracts --out=./contracts/fit.go
go run create_contract_address.go
```