grpc-helloworld
===

```
protoc --go_out=plugins=grpc:. proto/helloworld.proto
go run server.go

go run client.go
```
