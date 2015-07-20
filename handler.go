package compress

import (
	"compress/gzip"
	"net/http"
	"strings"

	"github.com/volatile/core"
	"github.com/volatile/core/coreutil"
)

// Use tells the core to use this handler.
func Use() {
	core.Use(func(c *core.Context) {
		var gzw *gzip.Writer
		gzw, c = setWriter(c)
		c.Next()
		gzw.Close()
	})
}

// LocalUse allows to compress locally, inside a single handler.
// Only the ResponseWriter is transmitted to avoid errors like calling c.Next() and risk to use multiple compressors over the response.
func LocalUse(c *core.Context, wf func(http.ResponseWriter)) {
	var gzw *gzip.Writer
	gzw, c = setWriter(c)
	wf(c.ResponseWriter)
	gzw.Close()
}

// setWriter create new ResponseWriter that will write to GZIP and returns it for the next handlers.
// The GZIP writer is also returned to close if ony after the handler has been exexuted.
func setWriter(c *core.Context) (*gzip.Writer, *core.Context) {
	var gzw *gzip.Writer

	if strings.Contains(c.Request.Header.Get("Accept-Encoding"), "gzip") && len(c.Request.Header.Get("Sec-WebSocket-Key")) == 0 {
		c.ResponseWriter.Header().Set("Content-Encoding", "gzip")

		gzw = gzip.NewWriter(c.ResponseWriter)
		// defer gzw.Close()

		// Set the new ResponseWriter.
		c.ResponseWriter = core.ResponseWriterBinder{
			Writer:         gzw,
			ResponseWriter: c.ResponseWriter,
			BeforeWrite: func(b []byte) {
				coreutil.SetDetectedContentType(c.ResponseWriter, b)
			},
		}
	}

	return gzw, c
}
