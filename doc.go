/*
Package compress is a handler for the core (https://godoc.org/github.com/volatile/core).
If accepted by the client, it compress the server response with GZIP.

Unlike almost all other compressing packages, it takes care to not handle contents that are already compressed (like JPEG, MPEG or PDF).
Trying to GZIP them not only wastes CPU but can potentially increase file sizes.

Usage

Use adds the handler to the default handlers stack:

	compress.Use()

Make sure to include the handler above any other handler that alter the response body.
*/
package compress
