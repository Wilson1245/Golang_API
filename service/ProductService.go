package service

import (
	"blogService/test/Gin/middlewares"
	"blogService/test/Gin/pojo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// post "/"
func SaveProduct(c *gin.Context) {
	product := pojo.Products{}
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error creating product --> "+err.Error())
		return
	}
	if havecreate := pojo.DBCreateProduct(product); !havecreate {
		c.JSON(http.StatusBadRequest, "Product already exist")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Product created",
		"Product": product,
	})
}

// get "/"
func FindAllProducts(c *gin.Context) {
	product := pojo.DBFindAllProducts()
	c.JSON(http.StatusOK, product)
}

// get "/:id"
func FindOneProduct(c *gin.Context) {
	productid, _ := strconv.Atoi(c.Param("id"))
	product := pojo.DBFindOneProduct(productid)
	if product.ProductId != 0 {
		c.JSON(http.StatusOK, product)
		return
	}
	c.JSON(http.StatusNotFound, "Product not found")
}

// PUT "/:id"
func UpdateProduct(c *gin.Context) {
	productid, _ := strconv.Atoi(c.Param("id"))
	product := pojo.Products{}
	c.BindJSON(&product)
	reuserId := middlewares.GetSessionUserId(c)
	userId := int(reuserId)
	if dbresult := pojo.DBPutProduct(productid, product, userId); !dbresult {
		c.JSON(http.StatusNotFound, "Product not found")
		return
	}
	c.JSON(http.StatusOK, "Product updated")
}

// DELETE "/:id"
func DeleteProduct(c *gin.Context) {
	productid, _ := strconv.Atoi(c.Param("id"))
	if dbresult := pojo.DBDeleteProduct(productid); !dbresult {
		c.JSON(http.StatusNotFound, "Product not found")
		return
	}
	c.JSON(http.StatusOK, "Product Delete Success!")
}
