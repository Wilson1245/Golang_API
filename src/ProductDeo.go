package src

import (
	"blogService/test/Gin/middlewares"
	"blogService/test/Gin/service"

	"github.com/gin-gonic/gin"
)

func AddProductRouter(r *gin.RouterGroup) {
	product := r.Group("/product")

	product.GET("/:id", service.FindOneProduct)
	product.GET("/", service.FindAllProducts)
	product.PUT("/:id", service.UpdateProduct)

	product.Use(middlewares.UserAuthSessionMiddle())
	{
		product.POST("/", service.SaveProduct)
	}

	// product.Use(middlewares.AdminAuthSessionMiddle())
	// {
	// 	product.DELETE("/:id", service.DeleteProduct)
	// }
}
