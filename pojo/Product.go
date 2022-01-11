package pojo

import (
	"blogService/test/Gin/database"
)

type Products struct {
	ProductId   int    `json:"productid"`
	ProductName string `json:"productname" binding:"required"`
	ProductQty  int    `json:"productqty" binding:"required"`
	ProductUser int    `json:"productuser" binding:"omitempty"`
}

type ProductsView struct {
	ProductId   int    `json:"productid"`
	ProductName string `json:"productname"`
	ProductQty  int    `json:"productqty" `
	ProductUser Users  `json:"productuser'`
}

// func (p *Products) setProductId(id int) {
// 	p.ProductId = id
// }

// func (p *Products) setProductName(name string) {
// 	p.ProductName = name
// }

// func (p *Products) setProductQty(qty int) {
// 	p.ProductQty = qty
// }
//OK
func DBFindAllProducts() []ProductsView {
	products := []Products{}
	database.DBconnect.Find(&products)
	productView := DBProductFindAllUser(products)
	return productView
}

//OK
func DBFindOneProduct(productId int) ProductsView {
	product := Products{}
	database.DBconnect.Where("product_id = ?", productId).Find(&product)
	productView := DBProductFindOneUser(product)
	return productView
}

//OK
func DBCreateProduct(product Products) (b bool) {
	if database.DBconnect.Where("product_name = ?", product.ProductName).Find(&product).RowsAffected == 1 {
		return false
	}
	return database.DBconnect.Create(&product).RowsAffected == 1
}

func DBDeleteProduct(productId int) (b bool) {
	product := Products{}
	return database.DBconnect.Where("product_id = ?", productId).Delete(&product).RowsAffected == 1
}

//OK
func DBPutProduct(productId int, product Products, userId int) (b bool) {
	beforeproduct := Products{}
	if database.DBconnect.Where("product_id = ?", productId).Find(&beforeproduct).RowsAffected != 1 {
		return false
	}
	if beforeproduct.ProductUser != userId {
		return false
	}
	beforeproduct.ProductName = product.ProductName
	beforeproduct.ProductQty = product.ProductQty
	dbresult := database.DBconnect.Where("product_id", productId).Save(&beforeproduct)
	return dbresult.RowsAffected == 1
}

//Find All Product
func DBProductFindAllUser(p []Products) []ProductsView {
	productView := ProductsView{}
	productViews := []ProductsView{}
	for _, product := range p {
		user := DBFindOneUser(product.ProductUser)
		productView.ProductId = product.ProductId
		productView.ProductName = product.ProductName
		productView.ProductQty = product.ProductQty
		productView.ProductUser = user
		productViews = append(productViews, productView)
	}
	return productViews
}

//Find One Product
func DBProductFindOneUser(p Products) ProductsView {
	productView := ProductsView{}
	user := DBFindOneUser(p.ProductUser)
	productView.ProductId = p.ProductId
	productView.ProductName = p.ProductName
	productView.ProductQty = p.ProductQty
	productView.ProductUser = user
	return productView
}
