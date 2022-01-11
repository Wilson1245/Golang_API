package service

import (
	"blogService/test/Gin/middlewares"
	"blogService/test/Gin/pojo"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// var userlist = []pojo.Users{}

// post "/"
func SaveUser(c *gin.Context) {
	user := pojo.Users{}
	err := c.BindJSON(&user)
	user.UserIdentity = "U"
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error creating user --> "+err.Error())
		return
	}
	// userlist = append(userlist, user)
	if havecreate := pojo.DBCreateUser(user); !havecreate {
		c.JSON(http.StatusBadRequest, "User already exist")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User created",
		"User":    user,
	})
}

// get "/"
func FindAllUsers(c *gin.Context) {
	user := pojo.DBFindAllUsers()
	log.Println("User --> ", user)
	c.JSON(http.StatusOK, user)
}

// get "/:id"
func FindOneUser(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Param("id"))
	user := pojo.DBFindOneUser(userid)
	if user.UserId != 0 {
		c.JSON(http.StatusOK, user)
		return
	}
	c.JSON(http.StatusNotFound, "User not found")
}

// PUT "/:id"
func UpdateUser(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Param("id"))
	user := pojo.Users{}
	c.BindJSON(&user)
	if dbresult := pojo.DBPutUser(userid, user); !dbresult {
		c.JSON(http.StatusNotFound, "User not found")
		return
	}
	c.JSON(http.StatusOK, "User updated")
}

// DELETE "/:id"
func DeleteUser(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Param("id"))
	if dbresult := pojo.DBDeleteUser(userid); !dbresult {
		c.JSON(http.StatusNotFound, "User not found")
		return
	}
	c.JSON(http.StatusOK, "User Delete Success!")
}

func LoginUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	log.Println("Login User", username, password)
	result, userId := pojo.DBCheckLogin(username, password)
	if result {
		middlewares.SaveAuthSession(c, userId)
		c.JSON(http.StatusOK, "User Login Success!")
		return
	}
	c.JSON(http.StatusBadRequest, "User Login Error!")
}

func LoginOutUser(c *gin.Context) {
	if result := middlewares.HasSession(c); !result {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "You never Login.",
		})
		return
	}
	middlewares.ClearAuthSession(c)
	c.JSON(http.StatusOK, "Login Out User!")
}
