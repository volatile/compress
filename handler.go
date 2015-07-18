package compress

import (
	"compress/gzip"
	"strings"

	"github.com/volatile/core"
	"github.com/volatile/core/coreutil"
)

// Use tells the core to use this handler.
func Use() {
	core.Use(func(c *core.Context) {
		gzw, c := setWriter(c)
		c.Next()
		gzw.Close()
	})
}

// LocalUse allows to compress locally, inside a single handler.
func LocalUse(c *core.Context, handler func(c *core.Context)) {
	gzw, c := setWriter(c)
	handler(c)
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
