# compiles libninty.c
clang libninty.c ../libninty.so -Wl,-rpath,. -o libninty-c
