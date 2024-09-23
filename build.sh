#!/bin/sh -eux

GOVERSION=1.23.1

# install go
curl -sL https://golang.org/dl/go${GOVERSION}.linux-amd64.tar.gz -o go${GOVERSION}.linux-amd64.tar.gz
tar -C ./ -xzf go${GOVERSION}.linux-amd64.tar.gz
export PATH=${PWD}/go/bin:$PATH
export GOPROXY=
export GOROOT=${PWD}/go

# copy wasm_exec.js
cp $(go env GOROOT)/misc/wasm/wasm_exec.js web/

# build
export GOOS=js
export GOARCH=wasm
go build -o web/rsushi.wasm github.com/neguse/rsushi
