package pojo

import (
	"blogService/test/Gin/database"
	"log"
)

type Buylists struct {
	BuyListId    int `json:"buylistid"`
	BuyUserId    int `json:"buyuserid" binding:"required"`
	BuyProductId int `json:"buyproductlist" binding:"required"`
}

type BuylistView struct {
	BuyListId  int          `json:"buylistid"`
	BuyUser    Users        `json:"buyuser"`
	BuyProduct ProductsView `json:"buyproduct"`
}

// func (b Buylists) setbuylistid(buylistId int) {
// 	b.BuyListId = buylistId
// }

// func (b Buylists) setbuyUser(buy_user_id int) {
// 	b.BuyUserId = buy_user_id
// }

// func (b Buylists) setbuybuylist(buy_buylist_id int) {
// 	b.BuyBuylistId = buy_buylist_id
// }

//OK
func DBFindAllBuylists() []BuylistView {
	buylists := []Buylists{}
	database.DBconnect.Find(&buylists)
	buylistviews := DBBuylistFindUserAndProduct(buylists)
	return buylistviews
}

//OK
func DBFindOneBuylist(buylistId int) (buylist Buylists) {
	database.DBconnect.Where("buylist_id = ?", buylistId).Find(&buylist)
	return
}

//OK
func DBCreateBuylist(buylist Buylists) (b bool) {
	return database.DBconnect.Create(&buylist).RowsAffected == 1
}

func DBDeleteBuylist(buylistId int) (b bool) {
	buylist := Buylists{}
	return database.DBconnect.Where("buy_list_id = ?", buylistId).Delete(&buylist).RowsAffected == 1
}

//OK
func DBPutBuylist(buylistId int, buylist Buylists) (b bool) {
	beforebuylist := Buylists{}
	if database.DBconnect.Where("buy_list_id = ?", buylistId).Find(&beforebuylist).RowsAffected != 1 {
		return false
	}
	log.Println("brforeBuylist --> ", beforebuylist)
	beforebuylist.BuyUserId = buylist.BuyUserId
	beforebuylist.BuyProductId = buylist.BuyProductId
	dbresult := database.DBconnect.Where("buy_list_id", buylistId).Save(&beforebuylist)
	return dbresult.RowsAffected == 1
}

func DBBuylistFindUserAndProduct(b []Buylists) []BuylistView {
	buylistviews := []BuylistView{}
	buylistview := BuylistView{}
	for _, buylist := range b {
		user := DBFindOneUser(buylist.BuyUserId)
		product := DBFindOneProduct(buylist.BuyProductId)
		buylistview.BuyListId = buylist.BuyListId
		buylistview.BuyUser = user
		buylistview.BuyProduct = product
		buylistviews = append(buylistviews, buylistview)
	}
	return buylistviews
}
