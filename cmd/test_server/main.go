package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/test_hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Test Hello!"})
	})
	err := r.Run(":7979")
	if err != nil {
		return
	}
}
