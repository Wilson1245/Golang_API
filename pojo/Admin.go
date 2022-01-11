package pojo

import "blogService/test/Gin/database"

type Admins struct {
	AdminId       int    `json:"admin"`
	AdminName     string `json:"adminname"`
	AdminPassword string `json:"adminpass"`
	AdminIdentity string `json:"adminidentity"`
}

func DBCheckAdminLogin(username string, password string) (bool, int) {
	admin := Admins{}
	dbresult := database.DBconnect.Where("admin_name = ? and admin_password = ?", username, password).Find(&admin)
	id := admin.AdminId
	return dbresult.RowsAffected == 1, id
}

func DBFindOneAdmin(id int) Admins {
	admin := Admins{}
	database.DBconnect.Where("admin_id = ?", id).Find(&admin)
	return admin
}

func DBFindAllAdmin() []Admins {
	admins := []Admins{}
	database.DBconnect.Find(&admins)
	return admins
}
