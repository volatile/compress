package compress

import "net/http"

var uncompresibleTypes = map[string]struct{}{
	"":                             {},
	"application/pdf":              {},
	"application/x-gzip":           {},
	"application/x-rar-compressed": {},
	"application/zip":              {},
	"image/gif":                    {},
	"image/jpeg":                   {},
	"image/png":                    {},
	"video/mpeg":                   {},
	"video/mp4":                    {},
	"video/x-flv":                  {},
	"video/webm":                   {},
	"audio/webm":                   {},
	"audio/aac":                    {},
	"audio/mp4":                    {},
	"video/H264":                   {},
	"audio/mpeg":                   {},
	"audio/wav":                    {},
}

// compressibleContentType tells if the content type must be compressed.
// This avoids compressing already compressed content, or content without a declared type.
func compressibleContentType(w http.ResponseWriter) bool {
	if _, ok := uncompresibleTypes[w.Header().Get("Content-Type")]; ok {
		return false
	}
	return true
}
