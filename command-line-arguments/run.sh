#!/usr/bin/env bash

SOURCE=./command-line-arguments.go
EXE=./command-line-arguments

go build "$SOURCE"
"$EXE" a b c d
rm "$EXE"
