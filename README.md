# Gin-Zerolog


[Gin](https://github.com/gin-gonic/gin) middleware for Logging with
[zerolog](https://github.com/rs/zerolog).

[![Travis branch](https://img.shields.io/travis/easonlin404/gin-zerolog/master.svg)](https://travis-ci.org/easonlin404/gin-zerolog)
[![Codecov branch](https://img.shields.io/codecov/c/github/easonlin404/gin-zerolog/master.svg)](https://codecov.io/gh/easonlin404/gin-zerolog)
[![Go Report Card](https://goreportcard.com/badge/github.com/easonlin404/gin-zerolog)](https://goreportcard.com/report/github.com/easonlin404/gin-zerolog)
[![GoDoc](https://godoc.org/github.com/easonlin404/gin-zerolog?status.svg)](https://godoc.org/github.com/easonlin404/gin-zerolog)
 

## Usage

### Start using it

Download and install it:

```sh
$ go get github.com/easonlin404/gin-zerolog
```

Import it in your code:

```go
import "github.com/easonlin404/gin-zerolog"
```

### Canonical example:

```go
package main

import (
	"github.com/easonlin404/gin-zerolog"
	"github.com/gin-gonic/gin"
)

func main() {

    r := gin.Default()
    r.Use(ginzerolog.Logger())

    r.GET("/", func(c *gin.Context) {
    	// your code
    })

    r.Run()
}
```