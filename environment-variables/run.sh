#!/usr/bin/env bash

SCRIPT="environment-variables.go"
go run "$SCRIPT"

echo "----------------------------------"

BAR=2 go run "$SCRIPT"