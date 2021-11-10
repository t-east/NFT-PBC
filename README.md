# PBC暗号によるストレージ監査システム
## 概要
ストレージへのアップロードデータとそのログをブロックチェーンに残して，それをPBC暗号を用いた第三者機関による監査システム

## 環境構築
### １，Solidity -> Golang
```
cd src/ethereum/
solc --optimize --abi ./solidities/Art.sol -o build --overwrite
solc --optimize --bin ./solidities/Art.sol -o build --overwrite
abigen --abi=./build/IndexTable.abi --bin=./build/IndexTable.bin --pkg=contracts --out=./contracts/art.go
cd ../
```
### ２，docker imageをビルド，コンテナを立ち上げる
```
docker-compose up -d
```
### ３．ganache-cliの秘密鍵情報を.envに書き込む
```
docker logs gana
```
でアドレス情報を参照できる
### ３，　１で作成したコントラクトのアドレスを取得
```
cd ../ethereum
go run create_contract_address.go
#出力結果を.envのCONTRACT_ADDRESSに追加
```

### 4, userコンテナ起動
```
docker exec -it user sh
cd src/User
go user.go
# userのginサーバが立ち上がる

#同様にsp, tpaもginサーバを立ち上げる
# 別ターミナル
docker exec -it sp sh
cd src/SP
go run sp.go

# 別ターミナル
docker exec -it tpa sh
cd src/TPA
go run tpa.go
```