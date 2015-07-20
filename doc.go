/*
Package compress is a handler for the Core.
If accepted by the client, it compress the server response with GZIP.

Make sure to include the handler above any other handler that alter the response body.

Usage

	package main

	import (
		"fmt"

		"github.com/volatile/compress"
		"github.com/volatile/core"
	)

	func main() {
		compress.Use()

		core.Use(func(c *core.Context) {
			fmt.Fprint(c.ResponseWriter, "Hello, World!")
		})

		core.Run()
	}
*/
package compress
