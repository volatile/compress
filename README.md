<p align="center"><img src="http://volatile.whitedevops.com/images/repositories/compress/logo.png" alt="Volatile Compress" title="Volatile Compress"><br><br></p>

Volatile Compress is a handler for the [Core](https://github.com/volatile/core).  
If accepted by the client, it compress the server response with GZIP.

Unlike almost all other compressing packages, it takes care to not handle contents that are already compressed (like JPEG, MPEG or PDF).  
Trying to GZIP them not only wastes CPU but can potentially increase file sizes.

Make sure to include the handler above any other handler that alter the response body.

## Installation

```Shell
$ go get github.com/volatile/compress
```

## Usage

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

[![GoDoc](https://godoc.org/github.com/volatile/compress?status.svg)](https://godoc.org/github.com/volatile/compress)
