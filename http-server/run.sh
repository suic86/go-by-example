#!/usr/bin/env bash

# terminates `go run http-server.go` after running the main script
# see: https://stackoverflow.com/a/2173421
trap "trap - SIGTERM && kill -- -$$" SIGINT SIGTERM EXIT

go run http-server.go &

for path in "hello" "headers";do
    curl "localhost:8090/$path"
done

