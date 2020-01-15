#!/bin/sh
# Build
# ./scripts/build.sh

# Free ports
killall -9 client

# Set environment variables
# Client
export BLT_GRPC_CLIENT_HOST="localhost"
export BLT_GRPC_CLIENT_PORT=8082

go build -o ./bin/client ./cmd/client/client.go
./bin/client
