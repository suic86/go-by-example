#!/usr/bin/env bash

# Build the binary

SOURCE=./command-line-subcommands.go
EXE=./command-line-subcommands
go build "$SOURCE"

FLAGS=("foo -enable -name=joe a1 a2"
    "bar -level 8 a1"
    "bar -enable a1"
)

for fs in "${FLAGS[@]}"; do
    echo "\$ $EXE $fs"
    "$EXE" $fs
    echo ""
done

# Cleanup: remove the binary
rm "$EXE"
