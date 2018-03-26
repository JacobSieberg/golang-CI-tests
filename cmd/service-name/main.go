package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//Version number
var Version = "7"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, fmt.Sprintf("Hello World Version: %s", Version))
	})
	r.Run()
}
