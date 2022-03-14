// Copyright 2021 Flamego. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package brotli

import (
	"bufio"
	"bytes"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andybalholm/brotli"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flamego/flamego"
)

func TestBrotli(t *testing.T) {
	calledBefore := false

	f := flamego.NewWithLogger(&bytes.Buffer{})
	f.Use(Brotli(Options{-10}))
	f.Use(func(r http.ResponseWriter) {
		r.(flamego.ResponseWriter).Before(func(rw flamego.ResponseWriter) {
			calledBefore = true
		})
	})
	f.Get("/", func() string { return "hello world!" })

	// Not accepting brotli
	resp := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)

	f.ServeHTTP(resp, req)

	ce := resp.Header().Get(headerContentEncoding)
	assert.NotEqual(t, "br", ce)

	// Accepting brotli
	resp = httptest.NewRecorder()
	req.Header.Set(headerAcceptEncoding, "br")
	f.ServeHTTP(resp, req)

	ce = resp.Header().Get(headerContentEncoding)
	assert.Equal(t, "br", ce)

	body, err := io.ReadAll(brotli.NewReader(resp.Body))
	require.NoError(t, err)
	assert.Equal(t, "hello world!", string(body))

	assert.True(t, calledBefore, "calledBefore")
}

type hijackableResponse struct {
	Hijacked bool
	header   http.Header
}

func newHijackableResponse() *hijackableResponse {
	return &hijackableResponse{header: make(http.Header)}
}

func (h *hijackableResponse) Header() http.Header       { return h.header }
func (h *hijackableResponse) Write([]byte) (int, error) { return 0, nil }
func (h *hijackableResponse) WriteHeader(int)           {}
func (h *hijackableResponse) Flush()                    {}
func (h *hijackableResponse) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h.Hijacked = true
	return nil, nil, nil
}

func TestResponseWriterHijack(t *testing.T) {
	hijackable := newHijackableResponse()

	f := flamego.NewWithLogger(&bytes.Buffer{})
	f.Use(Brotli())
	f.Use(func(rw http.ResponseWriter) {
		hj, ok := rw.(http.Hijacker)
		require.True(t, ok)

		_, _, err := hj.Hijack()
		assert.Nil(t, err)
	})

	r, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)

	r.Header.Set(headerAcceptEncoding, "br")
	f.ServeHTTP(hijackable, r)

	assert.True(t, hijackable.Hijacked)
}
