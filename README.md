# PyGo

Experiment to build statically compiled Python runtime in Go via CGO on Mac.

## Install

To date I had no luck getting brews pkg-config to report the correct Python paths.
Currently, the linker and include paths are hard coded...

## Static Linking

```
go build -ldflags "-linkmode 'external' -extldflags '-static'" main.go
```

results in:
```
ld: library not found for -lcrt0.o
```

which is apparently due to crt0 not being build statically n Mac OS X

## Conclusion

Not easily achievable on OSX.
