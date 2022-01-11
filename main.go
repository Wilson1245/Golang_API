package main

import (
	"blogService/test/Gin/database"
	"blogService/test/Gin/middlewares"
	"blogService/test/Gin/src"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1", middlewares.BasicAuth(), middlewares.EnableCookieSessions())
	src.AddUserRouter(v1)
	src.AddProductRouter(v1)
	src.AddBuyListRouter(v1)
	src.AddAdminRouter(v1)

	// v2 := r.Group("/v2")
	// src.AddProductRouter(v2)

	go func() {
		database.DD()
	}()
	r.Run(":8080")
}
