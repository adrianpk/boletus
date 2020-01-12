#!/bin/zsh

# Vars
HOST="localhost"
PORT="8081"
API_PATH="api"
API_VER="v1"
RES_PATH="users"
USER_USERNAME="username"

put () {
  echo "PUT $1"
  /usr/bin/curl -X PUT $1 --header 'Content-Type: application/json' -d @scripts/rest/update_user.json
}

# Request
put "http://$HOST:$PORT/$API_PATH/$API_VER/$RES_PATH/$USER_USERNAME"
