#!/bin/zsh

# Vars
HOST="localhost"
PORT="8081"
API_PATH="api"
API_VER="v1"
RES_PATH="users"
USER_ID="username"


get () {
  echo "GET $1"
  /usr/bin/curl -X GET $1
}

# Request
get "http://$HOST:$PORT/$API_PATH/$API_VER/$RES_PATH/$USER_ID"
