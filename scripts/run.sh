#!/bin/sh
# Build
# ./scripts/build.sh

# Free ports
killall -9 boletus

# Set environment variables
REV=$(eval git rev-parse HEAD)

# Service
export BLT_SVC_NAME="boletus"
export BLT_SVC_REVISION=$REV
export BLT_SVC_PINGPORT=8090

# Servers
export BLT_WEB_SERVER_PORT=8080
export BLT_JSONREST_SERVER_PORT=8081
export BLT_GRPC_SERVER_PORT=8082
export BLT_WEB_COOKIESTORE_KEY="iVuOOv4PNBnqTk2o13JsBMOPcPAe4p18"
export BLT_WEB_SECCOOKIE_HASH="iVuOOv4PNBnqTk2o13JsBMOPcPAe4p18"
export BLT_WEB_SECCOOKIE_BLOCK="iVuOOv4PNBnqTk2o"
export BLT_SITE_URL="localhost:8080"

# Postgres
export BLT_PG_SCHEMA="public"
export BLT_PG_DATABASE="boletus_dev"
export BLT_PG_HOST="localhost"
export BLT_PG_PORT="5432"
export BLT_PG_USER="boletus"
export BLT_PG_PASSWORD="boletus"
export BLT_PG_BACKOFF_MAXTRIES="3"

# Seeding
export BLT_SEEDING_FORCE="false"

# Scheduler
export BLT_SCHEDULER_ONE_MINUTES="1"
#export BLT_SCHEDULER_TWO_MINUTES="10"

# Scheduler
export BLT_RESERVATION_EXPIRE_MINUTES="60"

# Confirmation
# users/{slug}/{token}/confirm
export BLT_USER_CONFIRMATION_PATH="auth/%s/%s/confirm"
export BLT_USER_CONFIRMATION_SEND="false"
export BLT_USER_CONFIRMATION_DEBUG="true"

# Amazon SES MAiler
# Those are sample not usable keys
export AWS_ACCESS_KEY_ID=FIIAHI5FF3A2OG3MJEX5
export AWS_SECRET_KEY=9BiWmd5Hdgmk2rR4pwG332bHwvLGiJOoxLLtDy12

# Switches
export BLT_APP_USERNAME_UPDATABLE=false

# Client
export BLT_GRPC_CLIENT_HOST=8082

go build -o ./bin/boletus ./cmd/boletus.go
./bin/boletus
# go run -race main.go
