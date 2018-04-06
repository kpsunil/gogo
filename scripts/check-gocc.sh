#!/usr/bin/env bash
# Script to generate temporary tokens 

SRC='./src'
tmp='tmp'

mkdir -p "$tmp"

if [ -z "$GOPATH" ]; then
    echo "GOPATH environment variable is not set"
    exit 1
else
    goccpath="$GOPATH"/bin/gocc
    if [ -f "$goccpath" ]; then
        "$goccpath" -a -zip -o "$tmp" "$SRC"/imports.bnf
    else
        echo "gocc is not properly installed"
        exit 1
    fi
fi

