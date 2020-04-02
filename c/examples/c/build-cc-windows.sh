#!/bin/bash

x86_64-w64-mingw32-gcc fennel.c fennel.a libargp.a -lpthread -Wl,-rpath,. -l winmm -l ntdll -lws2_32 -o fennel-c
