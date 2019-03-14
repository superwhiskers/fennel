#!/bin/bash

GODEBUG=cgocheck=0 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 go build -buildmode=c-archive
mv c.a examples/c/fennel.a
mv c.h examples/c/fennel.h
cp -r ../keypair/ examples/c/
cp types.h examples/c/

