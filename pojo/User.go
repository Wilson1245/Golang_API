package pojo

import (
	"blogService/test/Gin/database"
	"log"
)

type Users struct {
	UserId       int    `json:"userid"`
	UserName     string `json:"username" binding:"required"`
	UserPassword string `json:"userpass" binding:"required"`
	UserEmail    string `json:"useremail" binding:"required"`
	UserIdentity string `json:"useridentity binding:"omitempty"`
}

//OK
func DBFindAllUsers() []Users {
	users := []Users{}
	database.DBconnect.Find(&users)
	return users
}

//OK
func DBFindOneUser(userId int) (user Users) {
	database.DBconnect.Where("user_id = ?", userId).Find(&user)
	return
}

//OK
func DBCreateUser(user Users) (b bool) {
	if database.DBconnect.Where("user_name = ?", user.UserName).Find(&user).RowsAffected == 1 {
		return false
	}
	return database.DBconnect.Create(&user).RowsAffected == 1
}

func DBDeleteUser(userId int) (b bool) {
	user := Users{}
	return database.DBconnect.Where("user_id = ?", userId).Delete(&user).RowsAffected == 1
}

//OK
func DBPutUser(userId int, user Users) (b bool) {
	beforeuser := Users{}
	if database.DBconnect.Where("user_id = ?", userId).Find(&beforeuser).RowsAffected != 1 {
		return false
	}
	log.Println("brforeUser --> ", beforeuser)
	beforeuser.UserName = user.UserName
	beforeuser.UserPassword = user.UserPassword
	beforeuser.UserEmail = user.UserEmail
	dbresult := database.DBconnect.Where("user_id", userId).Save(&beforeuser)
	return dbresult.RowsAffected == 1
}

func DBCheckLogin(username string, password string) (bool, int) {
	user := Users{}
	dbresult := database.DBconnect.Where("user_name = ? and user_password = ?", username, password).Find(&user)
	userid := user.UserId
	return dbresult.RowsAffected == 1, userid
}
