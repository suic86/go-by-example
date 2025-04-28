#!/usr/bin/env bash

fname="/tmp/dat"

echo "hello" >"$fname"
echo "go" >>"$fname"
go run reading-files.go "$fname"
