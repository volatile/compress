package compress

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/volatile/core"
	"github.com/volatile/core/httputil"
)

type compressWriter struct {
	io.Writer
	http.ResponseWriter
}

// compressors is a pool containing previously used writers.
// It creates new ones if run out.
var compressors = sync.Pool{New: func() interface{} {
	return gzip.NewWriter(nil)
}}

// Use adds the handler to the default handlers stack.
// It compress all the compressible responses.
func Use() {
	core.Use(func(c *core.Context) {
		if strings.Contains(c.Request.Header.Get("Accept-Encoding"), "gzip") && len(c.Request.Header.Get("Sec-WebSocket-Key")) == 0 {
			gzw := compressors.Get().(*gzip.Writer) // Get a writer from the pool.
			defer compressors.Put(gzw)              // When done, put the writer back in to the pool.

			gzw.Reset(c.ResponseWriter)
			defer gzw.Close()

			c.ResponseWriter = compressWriter{gzw, c.ResponseWriter}
			defer func() { c.ResponseWriter.Write(nil) }() // Make sure to always use the GZIP writer.

			defer c.Recover()
		}

		c.Next()
	})
}

// WriteHeader sets the compressing headers and writes into the GZIP, but only if the Content-Type header defines a compressible format.
// Otherwise, it calls the original Write method.
func (cw compressWriter) Write(b []byte) (int, error) {
	if compressibleContentType(httputil.SetDetectedContentType(cw.ResponseWriter, b)) {
		setGZIPHeaders(cw.ResponseWriter) // If WriteHeader has already been called, this line has no effect. But most of the time, it's not the case.
		return cw.Writer.Write(b)
	}
	return cw.ResponseWriter.Write(b)
}

// WriteHeader sets the compressing headers, but only if the Content-Type header defines a compressible format.
// Finally, it calls the original WriteHeader method.
func (cw compressWriter) WriteHeader(status int) {
	if compressibleContentType(cw.ResponseWriter.Header().Get("Content-Type")) {
		setGZIPHeaders(cw.ResponseWriter)
	}
	cw.ResponseWriter.WriteHeader(status)
}

// setGZIPHeaders sets the Content-Encoding header.
// Because the compressed content will have a new size, it also removes the Content-Length header as it could have been set downstream by another handler.
func setGZIPHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Del("Content-Length")
}
