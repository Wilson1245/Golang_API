package service

import (
	"blogService/test/Gin/pojo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginAdmin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	// log.Println("Login User", username, password)
	result, _ := pojo.DBCheckAdminLogin(username, password)
	if result {
		//middlewares.SaveAuthSession(c, userId)
		c.JSON(http.StatusOK, "Admin Login Success!")
		return
	}
	c.JSON(http.StatusBadRequest, "Admin Login Error!")
}

func FindAllAdmin(c *gin.Context) {
	admins := pojo.DBFindAllAdmin()
	c.JSON(http.StatusOK, admins)
}
