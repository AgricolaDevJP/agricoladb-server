# agricoladb-server

AgricolaDB (v2~) のサーバアプリケーション。GraphQL の API 提供と、データベースのスキーマ管理を行う。

## 開発用サーバ

```sh
cd /path/to/agricoladb-server
make docker-build # if needed
make docker-start
```

開発用サーバが http://localhost:8080/ で立ち上がる。データベースは localhost:3306 でアクセスできる。
ソースコード更新時に、自動でサーバが再起動する。
（依存する外部パッケージが増えたときには、 `make docker-restart-server` で手動再起動する必要がある。）

サーバを終了するときには以下の操作を行う：

```sh
make docker-stop
```

## コード生成

ent と GraphQL インテグレーションを用いて、スキーマや実装をコード生成している。

基本、編集するファイルは ent/schema/*.go だけで済むはずである。

```sh
make generate
```

上記コマンドを実行すると、以下の順でコード生成される：

1. ent による ORM client / migration などのコード生成
2. entgql による GraphQL スキーマ graph/ent.graphqls の生成
3. gqlgen による GraphQL サーバの Go 実装の生成

その後、必要に応じて graph/ent.resolvers.go などに含まれる未実装の部分をコーディングする。
