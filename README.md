<p align="center"><img src="https://cloud.githubusercontent.com/assets/9503891/8640486/b71678e6-28fa-11e5-8596-5fd6e63896d1.png" alt="Volatile Compress" title="Volatile Compress"><br><br></p>

Volatile Compress is a handler for the [Core](https://github.com/volatile/core).  
If accepted by the client, it compress the server response with GZIP.

Make sure to include the handler above any other handler that alter the response body.

## Installation

```Shell
$ go get -u github.com/volatile/compress
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
