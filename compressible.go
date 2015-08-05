package compress

import "strings"

var incompressibleTypes = map[string]struct{}{
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
	"video/h264":                   {},
	"audio/mpeg":                   {},
	"audio/wav":                    {},
}

// compressibleContentType indicates whether the content of ct type can be compressed.
func compressibleContentType(ct string) bool {
	_, ok := incompressibleTypes[strings.ToLower(ct)]
	return !ok
}
