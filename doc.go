/*
Package compress is a handler used by the core package in the handlers stack.
If accepted by the client, it compress the server response with GZIP.

Be sure to put it on the first position, above all other handlers!

Example

Here is the classic "Hello, World!" example:

	package main

	import (
		"net/http"

		"github.com/volatile/compress"
		"github.com/volatile/core"
	)

	func main() {
		core.Use(compress.Handler)

		core.Use(func(c *core.Context) {
			c.Response = []byte("Hello, World!")
		})

		core.Run()
	}
*/
package compress
