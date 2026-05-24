#!/bin/bash

protoc \
    --proto_path=../proto \
    --go_out=./pkg/helloworld \
    --go_opt=paths=source_relative \
    --go_opt=Mservice.proto=/helloworld \
    --go-grpc_out=./pkg/helloworld \
    --go-grpc_opt=paths=source_relative \
    --go-grpc_opt=Mservice.proto=/helloworld \
    service.proto
