# protobuf-sample

WebサーバーとAndroidアプリによるProtocol Buffer通信のサンプルコード

> **NOTE** gRPCではなく、サーバー側はREST HTTPサーバー。Content-TypeをJSONではなくProtocol Bufferでのやりとりにするサンプル

## TODO

- [ ] Client (Android)
- [x] Server (Go)

## Client

- Android
- 書籍データをWebサーバーにリクエストし、画面に表示する

### Library

- Retrofit
- Dagger2
- Epoxy
- マルチモジュール

## Server

- Go
- テストデータとして固定の書籍データリストを返却するだけのWebサーバー

## Library

- echo

## LICENCE

MIT LICENCE

## Author

yasukotelin
