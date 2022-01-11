package src

import (
	"blogService/test/Gin/middlewares"
	"blogService/test/Gin/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/user")

	user.POST("/login", service.LoginUser)

	user.POST("/", service.SaveUser)

	user.Use(middlewares.UserAuthSessionMiddle())
	{
		user.GET("/loginout", service.LoginOutUser)

		user.GET("/:id", service.FindOneUser)

		user.GET("/", service.FindAllUsers)

		user.PUT("/:id", service.UpdateUser)

	}
}
