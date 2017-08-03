# Delve

Go works with GDB out of the box, but a popular alternative debugging tool is Delve. Delve's
aim is to be a full featured debugger that is much easier to use and integrate with. With this
goal in mind, there are many editor/IDE front-ends as well a standalone front-end.

We're gonna create a sample program to step through, but first let's install some tools.


## Install

Since we have a working Go installation, the simplist way to install is to use `go get`.

```bash
go get github.com/dereparker/delve/cmd/dlv
```

### Note to OSX Users

Installing with `go get` will do some compiling. For this you need to ensure you have the
proper install toolchain. Run

```bash
xcode-select --install
```
