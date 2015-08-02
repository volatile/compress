package compress

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/volatile/core"
	"github.com/volatile/core/coreutil"
)

type compressWriter struct {
	io.Writer
	http.ResponseWriter
}

// Use adds a handler that compress all the compressible responses.
func Use() {
	core.Use(func(c *core.Context) {
		if strings.Contains(c.Request.Header.Get("Accept-Encoding"), "gzip") && len(c.Request.Header.Get("Sec-WebSocket-Key")) == 0 {
			gzw := gzip.NewWriter(c.ResponseWriter)
			defer gzw.Close()
			c.ResponseWriter = compressWriter{gzw, c.ResponseWriter} // Set the new ResponseWriter.
		}
		c.Next()
	})
}

func (cw compressWriter) Write(b []byte) (int, error) {
	coreutil.SetDetectedContentType(cw.ResponseWriter, b) // If WriteHeader has already been called, this line has no effect. But most of the time, it's not the case.

	if compressibleContentType(cw.ResponseWriter.Header().Get("Content-Type")) {
		setGZIPHeaders(cw.ResponseWriter) // If WriteHeader has already been called, this line has no effect. But most of the time, it's not the case.
		return cw.Writer.Write(b)
	}

	return cw.ResponseWriter.Write(b)
}

// WriteHeader sets the compressing headers, but only if the Content-Type header defines a compressible format.
// After that, it calls the real WriteHeader.
func (cw compressWriter) WriteHeader(status int) {
	if compressibleContentType(cw.ResponseWriter.Header().Get("Content-Type")) {
		setGZIPHeaders(cw.ResponseWriter)
	}
	cw.ResponseWriter.WriteHeader(status)
}

// setGZIPHeaders sets the Content-Encoding header.
// Because the compressed content will have a new size, it also removes the Content-Length header as it could have been set by another handler downstream.
func setGZIPHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Del("Content-Length")
}
