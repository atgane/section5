package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	mode := os.Getenv("MODE")

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"mode":    mode,
			"version": 1.1,
		})
	})

	router.Run("0.0.0.0:3000")
}
