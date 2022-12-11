# agricoladb-server

AgricolaDB (v2~) のサーバアプリケーション。GraphQL の API 提供と、データベースのスキーマ管理を行う。

## development server

```
cd /path/to/agricoladb-server
make docker-build # if needed
make docker-start
```

開発用サーバが http://localhost:8080/ で立ち上がる。データベースは localhost:3306 でアクセスできる。
ソースコード更新時に、自動でサーバが再起動する。
