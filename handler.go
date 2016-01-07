package compress

import (
	"compress/gzip"
	"net/http"
	"strings"
	"sync"

	"github.com/volatile/core"
	"github.com/volatile/core/httputil"
)

type compressWriter struct {
	http.ResponseWriter
}

// compressors is a pool containing previously used writers.
// It creates new ones if run out.
var compressors = sync.Pool{New: func() interface{} {
	return gzip.NewWriter(nil)
}}

// Use adds a handler that compress all the compressible responses.
func Use() {
	core.Use(func(c *core.Context) {
		if strings.Contains(c.Request.Header.Get("Accept-Encoding"), "gzip") && len(c.Request.Header.Get("Sec-WebSocket-Key")) == 0 {
			originalRW := c.ResponseWriter                   // Keep the original ResponseWriter.
			c.ResponseWriter = compressWriter{originalRW}    // Set the a ResponseWriter.
			defer func() { c.ResponseWriter = originalRW }() // Put back the original ResponseWriter for upstream handlers and core.PanicHandler.
		}
		c.Next()
	})
}

// WriteHeader sets the compressing headers and writes into the GZIP, but only if the Content-Type header defines a compressible format.
// Otherwise, it calls the original Write method.
func (cw compressWriter) Write(b []byte) (int, error) {
	if compressibleContentType(httputil.SetDetectedContentType(cw.ResponseWriter, b)) {
		setGZIPHeaders(cw.ResponseWriter) // If WriteHeader has already been called, this line has no effect. But most of the time, it's not the case.

		gzw := compressors.Get().(*gzip.Writer) // Get a writer from the pool.
		defer compressors.Put(gzw)              // When done, put the writer back in to the pool.

		gzw.Reset(cw.ResponseWriter)
		defer gzw.Close()

		return gzw.Write(b)
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
