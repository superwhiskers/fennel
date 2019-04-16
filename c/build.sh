#!/bin/bash

GODEBUG="cgocheck=0" go build -buildmode=c-archive
mv c.a examples/c/fennel.a
mv c.h examples/c/fennel.h
cp -r ../keypair/ examples/c/
cp types.h examples/c/

