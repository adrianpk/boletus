protoc -I internal/app/grpc/api/v1 -I --go_out=plugins=grpc:api internal/app/grpc/api/v1/api.proto

protoc -I/usr/local/include -I. \
  -I$GOPATH_GENERAL/src \
  -I$GOPATH_GENERAL/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  internal/app/grpc/api/v1/api.proto
