#!/usr/bin/env bash

mkdir -p folder

### create the files to embed
echo "hello go" >folder/single_file.txt
echo "123" >folder/file1.hash
echo "456" >folder/file2.hash

# don't put the files in folder to git
echo "*" >folder/.gitignore

go run embed-directive.go
