# fennel cshared examples

# how do i use these

- run `go build -buildmode=c-shared` with `cgocheck` disabled
- move `cshared` into the `examples/` directory and rename it to `fennel.so`
- move `cshared.h` into `examples/c/` and rename it to `fennel.h`
- copy the `keypair/` directory at the root of this repository into every subfolder of the `examples/` directory
- follow the common running instructions for the programming language of every example that you would wish to run, with the exeption of the c example. there's a buildscript for c

