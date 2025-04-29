#!/usr/bin/env bash

go run exit.go

go build exit.go
./exit
echo "$?"
rm ./exit