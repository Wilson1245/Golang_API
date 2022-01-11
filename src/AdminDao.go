package src

import (
	"blogService/test/Gin/service"

	"github.com/gin-gonic/gin"
)

func AddAdminRouter(r *gin.RouterGroup) {
	auth := r.Group("/auth")

	auth.GET("/", service.FindAllAdmin)

	auth.POST("/login", service.LoginAdmin)

	admin := auth.Group("/admin")

	admin.DELETE("/user/:id", service.DeleteUser)

	admin.DELETE("/product/:id", service.DeleteProduct)

	admin.DELETE("/buy/:id", service.DeleteBuylist)

	admin.GET("/", service.FindAllBuylists)

	admin.GET("/;id", service.FindOneBuylist)

	admin.PUT("/:id", service.UpdateBuylist)

}
