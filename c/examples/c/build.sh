#!/bin/bash

clang fennel.c fennel.a -lpthread -Wl,-rpath,. -o fennel-c
