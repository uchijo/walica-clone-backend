# walica-clone-backend

## walica-cloneについて

[walica](https://walica.jp)という、旅行などの立替を記録しておくと、最終的に誰が誰にいくら払えばいいのかを計算してくれる素晴らしいサービスがあります。
これを再現してみようという試みです。

バックエンドはこのリポジトリ、フロントエンドは[walica-clone-web](https://github.com/uchijo/walica-clone-web)にあります。

実際に作ったものに関しては、[こちら(https://walica-clone.uchijo.com)](https://walica-clone.uchijo.com)で公開しており、実際に動作しているところを確認いただけます。

## 使用しているツール類（主要なもののみ）

- gorm
  - SQLite
- grpc-gateway

## 動かし方

動かす際に、.envファイルが必要になります。設定する必要のある環境変数は以下のとおりです

- `DB_FILE`: sqliteのファイル名を指定します。
- `CORS_ORIGIN`: CORSで許可すべきホスト名を指定します。

例:

```.env
DB_FILE=dev.db
CORS_ORIGIN=http://localhost:3000
```

ビルド前に以下のコマンドを使い、マイグレーションを行う必要があります。

```bash
go run data/migrate/migrate.go
```

これらの設定を行えば動くようになります。

```bash
go run main.go
```

```bash
go build
```

## その他注意点など

grpcのポートは8080, grpc gatewayのポートは8090となっています。

`api.proto`を更新した際は、`cd proto && buf generate`でスタブの生成を行えます。
