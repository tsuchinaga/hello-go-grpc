# hello go grpc

久々にgRPC触ろうとしてたらAPIv2とやらになっていて、使い心地が変わっていたので確認。

## 前準備

* protocの入手
* `$ go get google.golang.org/protobuf/cmd/protoc-gen-go`
* `$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc`

## build

`$ protoc -I=./hellopb --go_out=./hellopb --go-grpc_out=./hellopb hellopb/hello.proto`

* `-I`: 出力先
* `--go_out`: protobufのmessageとかenumとかが吐かれる先
* `--go-grpc_out`: protobufのserviceから作られるserverとかclientが吐かれる先
* 最後にprotoファイルのパス

## プロジェクトで必要なパッケージ

`$ go get google.golang.org/grpc`

## 動かし方

* server
    * `go run server/main.go`
* client
    * `go run client/main.go -port {serverが表示したポート} -name {適当な文字列}`

### 例

`$ go run server/main.go`
> 2021/02/05 02:40:51 address is [::]:51010

`$ go run client/main.go -port 51010 -name World`
> 2021/02/05 02:44:17 message:"Hello, World"

`$ go run client/main.go -port 51010 -name Tsuchinaga`
> 2021/02/05 02:44:22 message:"Hello, Tsuchinaga"

## 参考

* [Protocol Buffer Basics: Go | Protocol Buffers | Google Developers](https://developers.google.com/protocol-buffers/docs/gotutorial)
* [Quick start | Go | gRPC](https://grpc.io/docs/languages/go/quickstart/)
