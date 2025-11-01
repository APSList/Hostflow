package main

import (
	gin "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello Property Service"})
	})

	err := r.Run(":8082")
	if err != nil {
		return
	} // port za Payment Service
}
