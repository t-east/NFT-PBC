# PBC暗号によるストレージ監査システム
## 概要
ストレージへのアップロードデータとそのログをブロックチェーンに残して，それをPBC暗号を用いた第三者機関による監査システム

## 環境構築
### Solidity -> Golang
```
cd ethereum/
solc --optimize --abi ./solidities/Art.sol -o build --overwrite
solc --optimize --bin ./solidities/Art.sol -o build --overwrite
abigen --abi=./build/IndexTable.abi --bin=./build/IndexTable.bin --pkg=contracts --out=./contracts/art.go
go run create_contract_address.go
```