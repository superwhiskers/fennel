#!/bin/bash
# compiles libninty.c
clang fennel.c ../fennel.so -Wl,-rpath,. -o fennel-c
