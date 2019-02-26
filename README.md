# go-string-format

[![Build Status](https://travis-ci.org/mgenware/go-string-format.svg?branch=master)](http://travis-ci.org/mgenware/go-string-format)

C# `String.Format` in Golang

## Usage
```go
import (
    strf "github.com/mgenware/go-string-format"
)

strf.Format("🐆{0} {1} {0}", "leopard", -3)
// returns "🐆leopard -3 leopard" 
```

## Benchmarks
```
goos: darwin
goarch: amd64
pkg: github.com/mgenware/go-string-format
BenchmarkSprintf-8              10000000               173 ns/op              48 B/op          1 allocs/op
BenchmarkFormat-8                1000000              1058 ns/op             186 B/op          7 allocs/op
BenchmarkFormatCore-8            5000000               285 ns/op              80 B/op          3 allocs/op
```