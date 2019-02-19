# fennel c bindings

# note

currently, only the example program in c works

also, `build.sh` files that automate the build process are provided
for both the c bindings and the c example program for \*nix users

# how do i use this

- run `go build -buildmode=c-archive` with `cgocheck` disabled
- move `c.a` into the `examples/c/` directory and rename it to `fennel.a`
- move `c.h` into `examples/c/` and rename it to `fennel.h`
- copy `types.h` into `examples/c/`
- copy the `keypair/` directory at the root of this repository into `examples/c/`
