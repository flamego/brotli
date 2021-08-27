// Copyright 2021 Flamego. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package brotli

import (
	"strings"

	"github.com/andybalholm/brotli"

	"github.com/flamego/flamego"
)

const (
	headerAcceptEncoding  = "Accept-Encoding"
	headerContentEncoding = "Content-Encoding"
	headerVary            = "Vary"
)

// Options represents a struct for specifying configuration options for the Brotli middleware.
type Options struct {
	// Compression level. Can be DefaultCompression(6) or any integer value
	// between BestSpeed(0) and BestCompression(11) inclusive.
	CompressionLevel int
}

func isCompressionLevelValid(level int) bool {
	return level == brotli.DefaultCompression ||
		(level >= brotli.BestSpeed && level <= brotli.BestCompression)
}

func prepareOptions(options []Options) Options {
	var opt Options
	if len(options) > 0 {
		opt = options[0]
	}

	if !isCompressionLevelValid(opt.CompressionLevel) {
		// For web content, level 5 seems to be a sweet spot.
		opt.CompressionLevel = 5
	}
	return opt
}

// Brotli returns a Handler that adds brotli compression to all requests.
// Make sure to include the Brotli middleware above other middleware
// that alter the response body (like the render middleware).
func Brotli(options ...Options) flamego.Handler {
	opt := prepareOptions(options)

	return flamego.ContextInvoker(func(ctx flamego.Context) {
		if !strings.Contains(ctx.Request().Header.Get(headerAcceptEncoding), "br") {
			return
		}

		headers := ctx.ResponseWriter().Header()
		headers.Set(headerContentEncoding, "br")
		headers.Set(headerVary, headerAcceptEncoding)

		// We've made sure compression level is valid in prepareOptions,
		// no need to check same error again.
		br := brotli.NewWriterLevel(ctx.ResponseWriter(), opt.CompressionLevel)

		defer br.Close()

		ctx.Next()

		// Delete content length after we know we have been written to
		ctx.ResponseWriter().Header().Del("Content-Length")
	})
}
