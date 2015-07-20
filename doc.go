/*
Package compress is a handler for the Core.
If accepted by the client, it compress the server response with GZIP.

Make sure to include the handler above any other handler that alter the response body.

Global usage

compress.Use() adds a handler to the Core so all the responses are compressed.

Make sure to include the handler above any other handler that alter the response body.

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

Local usage

compress.LocalUse(*core.Context, func(http.ResponseWriter)) can be used to compress the response inside a specific handler.
Only the ResponseWriter is transmitted to avoid errors like calling c.Next() and risk to use multiple compressors over the response.

Make sure to not use a local compress if the global handler is set.

	package main

	import (
		"fmt"

		"github.com/volatile/compress"
		"github.com/volatile/core"
	)

	func main() {
		compress.LocalUse(c, func(c *core.Context) {
			fmt.Fprint(c.ResponseWriter, "Hello, World!")
		})

		core.Run()
	}
*/
package compress
