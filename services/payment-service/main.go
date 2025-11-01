package main

import (
	gin "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello Payment Service"})
	})

	err := r.Run(":8083")
	if err != nil {
		return
	} // port za Payment Service
}
