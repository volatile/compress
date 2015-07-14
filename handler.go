package compress

import (
	"compress/gzip"
	"net/http"
	"strings"

	"github.com/volatile/core"
)

// Use tells the core to use this handler.
func Use() {
	core.Use(func(c *core.Context) {
		if strings.Contains(c.Request.Header.Get("Accept-Encoding"), "gzip") && len(c.Request.Header.Get("Sec-WebSocket-Key")) == 0 {
			c.ResponseWriter.Header().Set("Content-Encoding", "gzip")

			gzw := gzip.NewWriter(c.ResponseWriter)
			defer gzw.Close()

			// Pass a new ResponseWriter
			c.NextWriter(core.ResponseWriterBinder{
				Writer:         gzw,
				ResponseWriter: c.ResponseWriter,
				BeforeWrite: func(b []byte) {
					if len(c.ResponseWriter.Header().Get("Content-Type")) == 0 {
						c.ResponseWriter.Header().Set("Content-Type", http.DetectContentType(b))
					}
				},
			})
		} else {
			c.Next()
		}
	})
}
