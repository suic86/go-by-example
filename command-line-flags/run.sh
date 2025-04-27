#!/usr/bin/env bash

# Build the binary

SOURCE=./command-line-flags.go
EXE=./command-line-flags
go build "$SOURCE"

# Run different configuration of command-line flags

# "-word=opt -numb=7 -fork -svar=flag" 
# gives values to all flags

# "-word=opt"
# if flags are omitted, their default values are taken

# "-word=opt a1 a2 a3"
# trailing positional arguments can be provided after any flags

# "-word=opt a1 a2 a3 -numb=7"
# `flag` package require all flags to appear before positional arguments
# (otherwise the flags will be interpreted as positional arguments.

# "-h" or "--help" automaticall generate help

# "-wat"
# if you provide a flag that wasn't specified to the `flag` package
# the program will print an error message and show the help text again.

FLAGS=("-word=opt -numb=7 -fork -svar=flag" \
       "-word=opt" \
       "-word=opt a1 a2 a3" \
       "-word=opt a1 a2 a3 -numb=7" \
       "-h" \
       "-wat" \
)

for fs in "${FLAGS[@]}"; do
    echo "\$ $EXE $fs" 
    "$EXE" $fs
    echo ""
done

# Cleanup: remove the binary
rm "$EXE"