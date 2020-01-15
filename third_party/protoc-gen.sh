#!/bin/sh
protoc --proto_path=pkg/grpc/api/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/grpc/api/v1 api.proto
