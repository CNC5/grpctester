#!/usr/bin/bash

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
GOPATH_BIN="$(go env GOPATH)/bin"
if [[ ":$PATH:" != *":$GOPATH_BIN:"* ]]; then
    export PATH="$GOPATH_BIN:$PATH"
fi

mkdir -p gen

protoc -I proto \
    --go_out=gen --go_opt=paths=source_relative \
    --go-grpc_out=gen --go-grpc_opt=paths=source_relative \
    proto/helloworld.proto && \
echo "proto generated"

go build client.go && \
echo "client built"

go build server.go && \
echo "server built"
