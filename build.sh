#!/bin/sh

cp $(go env GOROOT)/misc/wasm/wasm_exec.js web/

export GOOS=js
export GOARCH=wasm
go build -o web/rsushi.wasm github.com/neguse/rsushi
