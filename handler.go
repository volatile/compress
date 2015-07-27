package compress

import (
	"compress/gzip"
	"strings"

	"github.com/volatile/core"
	"github.com/volatile/core/coreutil"
)

// Use tells the core to compress all responses.
func Use() {
	core.Use(func(c *core.Context) {
		if strings.Contains(c.Request.Header.Get("Accept-Encoding"), "gzip") && len(c.Request.Header.Get("Sec-WebSocket-Key")) == 0 {
			c.ResponseWriter.Header().Set("Content-Encoding", "gzip")

			gzw := gzip.NewWriter(c.ResponseWriter)
			defer gzw.Close()

			// Set the new ResponseWriter.
			c.ResponseWriter = core.ResponseWriterBinder{
				Writer:         gzw,
				ResponseWriter: c.ResponseWriter,
				BeforeWrite: func(b []byte) {
					coreutil.SetDetectedContentType(c, b)
				},
			}
		}

		c.Next()
	})
}
