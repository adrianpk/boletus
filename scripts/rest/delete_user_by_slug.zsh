#!/bin/zsh

# Vars
HOST="localhost"
PORT="8080"
API_PATH="api"
API_VER="v1"
RES_PATH="users"
USER_SLUG="username1-129a82a252c2"


delete () {
  echo "DELETE $1"
  /usr/bin/curl -X DELETE $1
}

# Request
delete "http://$HOST:$PORT/$API_PATH/$API_VER/$RES_PATH/$USER_SLUG"
