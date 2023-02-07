package main

import (
	"github.com/gin-gonic/gin"
)

func handler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	// get an instence of gin engine
	r := gin.New()

	r.GET("/ping", handler)

	r.Run("127.0.0.1:3000") // listen and serve on 127.0.0.1:3000
}
