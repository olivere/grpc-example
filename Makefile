.PHONY: protos server client
PKG=github.com/olivere/grpc-example

default: server

protos:
	protoc -I. tasks.proto --go_out=plugins=grpc:tasks

server:
	go build $(PKG)/cmd/server

client:
	go build $(PKG)/cmd/client

serve: server
	./server -addr=":8000" -cert=./etc/star.go.crt -key=./etc/star.go.key

cli: client
	./client -addr=":8000" -cert=./etc/star.go.crt -name=rpc.go

