package main

import (
	"github.com/gin-gonic/gin"
)

func info(c *gin.Context) {
	address := gin.H{"address": c.Request.RemoteAddr}
	c.JSON(200, address)
}

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/info", info)
	}

	router.Run(":8080")
}
