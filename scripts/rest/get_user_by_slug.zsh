#!/bin/zsh

# Vars
HOST="localhost"
PORT="8081"
API_PATH="api"
API_VER="v1"
RES_PATH="users"
USER_SLUG="username1-129a82a252c2"


get () {
  echo "POST $1"
  /usr/bin/curl -X GET $1
}

# Request
get "http://$HOST:$PORT/$API_PATH/$API_VER/$RES_PATH/$USER_SLUG"
