#!/bin/sh

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

GOPATH_BIN="$(go env GOPATH)/bin"
if [[ ":$PATH:" != *":$GOPATH_BIN:"* ]]; then
    export PATH="$GOPATH_BIN:$PATH"
fi

if [ -x "$(command -v apk)" ];
then
    sudo apk add --no-cache protobuf
elif [ -x "$(command -v apt-get)" ];
then
    sudo apt-get install protobuf-compiler
#elif [ -x "$(command -v dnf)" ];
#then
#    sudo dnf install ""
#elif [ -x "$(command -v zypper)" ];
#then
#    sudo zypper install ""
else
    echo "FAILED TO INSTALL PACKAGE: Package manager not found. You must manually install: 'protobuf'">&2;
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
