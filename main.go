package main

import (
	"jwt_example/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	config.InitEnv()
}
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
