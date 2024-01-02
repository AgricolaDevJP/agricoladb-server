<div align="center">
    <img alt="AgricolaDB Server" src="docs/img/agricoladb-server-logo.png" />
</div>

<br />

![release](https://img.shields.io/github/v/release/AgricolaDevJP/agricoladb-server)
![go-version](https://img.shields.io/github/go-mod/go-version/AgricolaDevJP/agricoladb-server?filename=go.mod)
[![on-push-main](https://github.com/AgricolaDevJP/agricoladb-server/actions/workflows/on-push-main.yml/badge.svg)](https://github.com/AgricolaDevJP/agricoladb-server/actions/workflows/on-push-main.yml)

<br />

# agricoladb-server

AgricolaDB(v2~)のサーバアプリケーション。GraphQL APIの提供と、データベースのスキーマ管理を行う。

## コード生成

entとGraphQLインテグレーションを用いて、スキーマや実装をコード生成している。

基本、編集するファイルは ent/schema/*.goだけで済むはずである。

```sh
make generate
```

上記コマンドを実行すると、以下の順でコード生成される：

1. entによるORMclient/migrationなどのコード生成
2. entgqlによるGraphQLスキーマgraph/ent.graphqlsの生成
3. gqlgenによるGraphQLサーバのGo実装の生成

その後、必要に応じて graph/ent.resolvers.goなどに含まれる未実装の部分をコーディングする。

## LICENSE

### Source Code

MIT License

Copyright (c) 2022 ASAKURA Kazuki

### Information of "Agricola"

The information provided in this API is quoted or translated from the products of "Agricola" and is copyrighted by the respective companies:

Copyright (c) 2007 Lookout Games

Copyright (c) 2008 Z-Man Games, Inc.

Copyright (c) 2008 Hobby Japan Co., Ltd.
