package main

import (
	"os"

	"aviary/fox"

	"github.com/gin-gonic/gin"
)

func getPort() string {
	port := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		port = ":" + val
	}
	return port
}

func main() {
	r := gin.Default()
	r.POST("/fox", fox.Handler)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found."})
	})
	r.Run(getPort())
}
