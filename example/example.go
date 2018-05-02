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
