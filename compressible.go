package compress

import "strings"

var incompressibleTypes = map[string]struct{}{
	"": {},
	"application/octet-stream":     {},
	"application/pdf":              {},
	"application/x-font-woff":      {},
	"application/x-gzip":           {},
	"application/x-rar-compressed": {},
	"application/zip":              {},
	"audio/aac":                    {},
	"audio/mp4":                    {},
	"audio/mpeg":                   {},
	"audio/wav":                    {},
	"audio/webm":                   {},
	"image/gif":                    {},
	"image/jpeg":                   {},
	"image/png":                    {},
	"video/h264":                   {},
	"video/mp4":                    {},
	"video/mpeg":                   {},
	"video/webm":                   {},
	"video/x-flv":                  {},
}

// compressibleContentType indicates whether the content of ct type can be compressed.
func compressibleContentType(ct string) bool {
	_, ok := incompressibleTypes[strings.ToLower(ct)]
	return !ok
}
