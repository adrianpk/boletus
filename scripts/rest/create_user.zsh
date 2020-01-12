#!/bin/zsh

# Vars
HOST="localhost"
PORT="8081"
API_PATH="api"
API_VER="v1"
RES_PATH="users"


post () {
  echo "POST $1"
  /usr/bin/curl -X POST $1 --header 'Content-Type: application/json' -d @scripts/rest/create_user.json
}

# Request
post "http://$HOST:$PORT/$API_PATH/$API_VER/$RES_PATH"
