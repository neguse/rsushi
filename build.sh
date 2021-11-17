#!/bin/sh

# install go
curl https://golang.org/dl/go1.17.3.linux-amd64.tar.gz -o go1.17.3.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.3.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# copy wasm_exec.js
cp $(go env GOROOT)/misc/wasm/wasm_exec.js web/

# build
export GOOS=js
export GOARCH=wasm
go build -o web/rsushi.wasm github.com/neguse/rsushi
