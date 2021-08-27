# brotli

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/flamego/brotli/Go?logo=github&style=for-the-badge)](https://github.com/flamego/brotli/actions?query=workflow%3AGo)
[![Codecov](https://img.shields.io/codecov/c/gh/flamego/brotli?logo=codecov&style=for-the-badge)](https://app.codecov.io/gh/flamego/brotli)
[![GoDoc](https://img.shields.io/badge/GoDoc-Reference-blue?style=for-the-badge&logo=go)](https://pkg.go.dev/github.com/flamego/brotli?tab=doc)
[![Sourcegraph](https://img.shields.io/badge/view%20on-Sourcegraph-brightgreen.svg?style=for-the-badge&logo=sourcegraph)](https://sourcegraph.com/github.com/flamego/brotli)

Package brotli is a middleware that provides brotli compression to responses for [Flamego](https://github.com/flamego/flamego).

## Installation

The minimum requirement of Go is **1.16**.

    go get github.com/flamego/brotli


## Getting started

```go
package main

import (
	"github.com/flamego/brotli"
	"github.com/flamego/flamego"
)

func main() {
	f := flamego.Classic()
	f.Use(brotli.Brotli())
	f.Get("/", func() string {
		return "ok"
	})
	f.Run()
}
```

## License

This project is under the MIT License. See the [LICENSE](LICENSE) file for the full license text.
