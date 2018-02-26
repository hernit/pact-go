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

#bundle exec rackup -p 9393 provider-interface.ru &
pushd ../../../; go build -o examples/messages/provider/pact-go; popd
./pact-go daemon -v -l DEBUG &
pid=$!
# bundle exec pact-provider-verifier message-pact.json --provider-base-url=http://localhost:9393
# Run testss
go test -v .
