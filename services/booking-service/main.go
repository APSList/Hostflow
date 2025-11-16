package main

import (
	"net/http"

	gin "github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/health", healthCheck)

	err := router.Run("0.0.0.0:8084")
	if err != nil {
		return
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"health": true})
}
