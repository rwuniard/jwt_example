package main

import (
	"jwt_example/config"
	"jwt_example/controllers"
	"jwt_example/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	config.InitEnv()
	config.ConnectDB()
	config.CreateDB()
}
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("/logout", middleware.RequireAuth, controllers.Logout)
	r.Run()
}
