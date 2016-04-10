/*
Package compress is a handler for the core (https://godoc.org/github.com/volatile/core).
It provides a clever gzip compressing handler.

It takes care to not handle small contents, or contents that are already compressed (like JPEG, MPEG or PDF).
Trying to gzip them not only wastes CPU but can potentially increase the response size.

Usage

Use adds the handler to the default handlers stack:
	compress.Use()

Make sure to include this handler above any other handler that alter the response body.
*/
package compress
