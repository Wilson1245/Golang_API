package src

import (
	"blogService/test/Gin/middlewares"
	"blogService/test/Gin/service"

	"github.com/gin-gonic/gin"
)

func AddAdminRouter(r *gin.RouterGroup) {
	auth := r.Group("/auth", middlewares.BasicAuth())

	// auth.GET("/", service.FindAllAdmin)

	auth.POST("/login", service.LoginAdmin)

	admin := auth.Group("/admin", middlewares.AdminAuthSessionMiddle())

	admin.GET("/loginout", service.AdminLoginOut)

	admin.GET("/user", service.FindAllUsers)

	admin.GET("/product", service.FindAllProducts)

	admin.GET("/buy", service.FindAllBuylists)

	admin.DELETE("/user/:id", service.DeleteUser)

	admin.DELETE("/product/:id", service.DeleteProduct)

	admin.DELETE("/buy/:id", service.DeleteBuylist)

	admin.DELETE("/buy/automatic", service.AdminAutoMaticDeleteBuyList)

	admin.GET("/", service.FindAllBuylists)

	admin.GET("/;id", service.FindOneBuylist)

	admin.PUT("/:id", service.UpdateBuylist)

}
