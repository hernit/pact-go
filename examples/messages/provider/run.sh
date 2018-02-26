#!/bin/bash

pid=0
dir=$(pwd)
function cleanup {
  if [ $pid -gt 0 ]; then
    echo "--> Shutting down running processes"
    kill -9 $pid
  fi
  cd $dir
}

trap cleanup INT TERM EXIT

bundle install --binstubs
pushd ../../../; go build -o examples/messages/provider/pact-go; popd
./pact-go daemon -v -l DEBUG &
pid=$!
go test -v .
