<p align="center"><img src="http://volatile.whitedevops.com/images/repositories/compress/logo.png" alt="Volatile Compress" title="Volatile Compress"><br><br></p>

Volatile Compress is a handler for the [Core](https://github.com/volatile/core).  
If accepted by the client, it compress the server response with GZIP.



## Installation

```Shell
$ go get -u github.com/volatile/compress
```

## Usage

### Global

`compress.Use()` adds a handler to the Core so all the responses are compressed.

Make sure to include the handler above any other handler that alter the response body.

```Go
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
```

### Local

`compress.LocalUse(*core.Context, func(http.ResponseWriter))` can be used to compress the response inside a specific handler.  
Only the ResponseWriter is transmitted to avoid errors like calling c.Next() and risk to use multiple compressors over the response.

Make sure to not use a local compress if the global handler is set.

```Go
package main

import (
	"fmt"
	"net/http"

	"github.com/volatile/compress"
	"github.com/volatile/core"
)

func main() {
	core.Use(func(c *core.Context) {
		compress.LocalUse(c, func(w http.ResponseWriter) {
			fmt.Fprint(w, "Hello, World!")
		})
	})

	core.Run()
}
```

[![GoDoc](https://godoc.org/github.com/volatile/compress?status.svg)](https://godoc.org/github.com/volatile/compress)
