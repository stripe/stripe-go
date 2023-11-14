#!/usr/bin/env bash

# This script is used to expose the local server to the internet for Salesforce
# to be able to access it, in order to test changes locally before deploying.

function install_localtunnel() {
  EXEC_PATH="$1"
  NOT_THIS_NODE="$2"
  TEST="$3"
  if [ -z "$EXEC_PATH" ] || [ -n "$NOT_THIS_NODE" ]; then
    if [ -z "$TEST" ]; then
      npm install localtunnel -g
    else
      echo npm install localtunnel -g
    fi
    echo "LocalTunnel installed."
  elif [ -n "$TEST" ]; then
    echo "LocalTunnel already installed."
  fi
}

if [ "$1" == "test" ]; then
  echo "Should install..."
  install_localtunnel "" "" "test"
  echo "Should install..."
  install_localtunnel "/Users/jmather-c/.nodenv/shims/lt" "nodenv: lt: command not found" "test"
  echo "Should not install..."
  install_localtunnel "/Users/jmather-c/.nodenv/shims/lt" "" "test"
  exit 0
fi

LOCAL_TUNNEL_EXEC=$(which lt 2>/dev/null)
LOCAL_TUNNEL_EXEC_NOTFOUND=$(lt 2>&1 | grep "command not found")

install_localtunnel "$LOCAL_TUNNEL_EXEC" "$LOCAL_TUNNEL_EXEC_NOTFOUND"

HOST_TARGET="$1"
if [ -z "$HOST_TARGET" ]; then
  HOST_TARGET=$(whoami)
fi

lt -h "https://:11f79d6f2dc2dd51d4e98114e8c1eeb5@lt.jmather.com" -p 3100 -l 127.0.0.1 -s "$HOST_TARGET"