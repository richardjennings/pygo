# PyGo

Experiment to build statically compiled Python runtime in Go via Cgo on Mac.

## Install

Requires python3-embed resolvable via pkg-config

```
brew install python@3.12
```

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

A fully static binary does not look easily achievable on OSX.

### And then

Can we include python-embed statically ? 

