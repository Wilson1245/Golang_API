package service

import (
	"blogService/test/Gin/middlewares"
	"blogService/test/Gin/pojo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginAdmin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	// log.Println("Login User", username, password)
	result, userId := pojo.DBCheckAdminLogin(username, password)
	if result {
		middlewares.SaveAuthSession(c, userId)
		c.JSON(http.StatusOK, "Admin Login Success!")
		return
	}
	c.JSON(http.StatusBadRequest, "Admin Login Error!")
}

func FindAllAdmin(c *gin.Context) {
	admins := pojo.DBFindAllAdmin()
	c.JSON(http.StatusOK, admins)
}

func AdminLoginOut(c *gin.Context) {
	if middlewares.HasSession(c) {
		middlewares.ClearAuthSession(c)
		c.JSON(http.StatusOK, "Admin Login Out!")
		return
	} else {
		c.JSON(http.StatusBadRequest, "Admin Login Out Error!")
		return
	}
}

func AdminAutoMaticDeleteBuyList(c *gin.Context) {
	buylists := pojo.DBFindAllBuylists()
	for _, b := range buylists {
		if b.BuyUser.UserId == 0 || b.BuyProduct.ProductId == 0 {
			pojo.DBDeleteBuylist(b.BuyListId)
		}
	}
	c.JSON(http.StatusOK, "Admin Auto Matic Clear complete.")
}
