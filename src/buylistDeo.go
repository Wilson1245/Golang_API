package src

import (
	"blogService/test/Gin/middlewares"
	"blogService/test/Gin/service"

	"github.com/gin-gonic/gin"
)

func AddBuyListRouter(r *gin.RouterGroup) {
	buy := r.Group("/buylist")

	buy.Use(middlewares.UserAuthSessionMiddle())
	{
		buy.POST("/", service.SaveBuylist)
	}

	// buy.Use(middlewares.AdminAuthSessionMiddle())
	// {
	// 	buy.DELETE("/:id", service.DeleteBuylist)

	// 	buy.GET("/:id", service.FindOneBuylist)

	// 	buy.GET("/", service.FindAllBuylists)

	// 	buy.PUT("/:id", service.UpdateBuylist)
	// }
}
