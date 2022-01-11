package middlewares

import (
	"blogService/test/Gin/pojo"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const key = "AWE123"

//使用Cookie 保存 Session
func EnableCookieSessions() gin.HandlerFunc {
	stroe := cookie.NewStore([]byte(key))
	return sessions.Sessions("SIMPLE", stroe)
}

// Admin Session 中間件
func AdminAuthSessionMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionValue := session.Get("userId")
		user := pojo.DBFindOneUser(sessionValue.(int))
		if user.UserIdentity != "T" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "You not Admin.",
			})
			c.Abort()
			return
		} else {
			c.Set("userId", sessionValue)
			c.Next()
		}
	}
}

// User Session 中間件
func UserAuthSessionMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionValue := session.Get("userId")
		// user := pojo.DBFindOneUser(sessionValue.(int))
		if sessionValue == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "You not User.",
			})
			c.Abort()
			return
		} else {
			c.Set("userId", sessionValue)
			c.Next()
		}
	}
}

//註冊或登入時儲存Session
func SaveAuthSession(c *gin.Context, userId int) {
	session := sessions.Default(c)
	session.Set("userId", userId)
	session.Save()
}

//登出或離開時清除Session
func ClearAuthSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func HasSession(c *gin.Context) bool {
	session := sessions.Default(c)
	if sessionValue := session.Get("userId"); sessionValue == nil {
		return false
	}
	return true
}

func GetSessionUserId(c *gin.Context) int {
	session := sessions.Default(c)
	sessionValue := session.Get("userId")
	if sessionValue == nil {
		return 0
	}
	return sessionValue.(int)
}
