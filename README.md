# brotli

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/flamego/brotli/Go?logo=github&style=for-the-badge)](https://github.com/flamego/brotli/actions?query=workflow%3AGo)
[![GoDoc](https://img.shields.io/badge/GoDoc-Reference-blue?style=for-the-badge&logo=go)](https://pkg.go.dev/github.com/flamego/brotli?tab=doc)

Package brotli is a middleware that provides brotli compression to responses for [Flamego](https://github.com/flamego/flamego).

## Installation

```zsh
go get github.com/flamego/brotli
```

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

## Getting help

- Read [documentation and examples](https://flamego.dev/middleware/brotli.html).
- Please [file an issue](https://github.com/flamego/flamego/issues) or [start a discussion](https://github.com/flamego/flamego/discussions) on the [flamego/flamego](https://github.com/flamego/flamego) repository.

## License

This project is under the MIT License. See the [LICENSE](LICENSE) file for the full license text.
