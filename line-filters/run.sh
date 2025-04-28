#!/usr/bin/env bash

f="/tmp/lines"

echo 'hello' >"$f"
echo 'filter' >>"$f"

go run line-filters.go <"$f"
