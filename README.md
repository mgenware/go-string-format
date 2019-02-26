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