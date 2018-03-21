package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/login", loginEndpoint)
		v1.POST("/login", loginEndpoint)
	}

	router.Run()
}

func loginEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "login"})
}
