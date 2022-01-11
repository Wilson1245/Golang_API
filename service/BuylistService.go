package service

import (
	"blogService/test/Gin/pojo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// post "/"
func SaveBuylist(c *gin.Context) {
	buylist := pojo.Buylists{}
	err := c.BindJSON(&buylist)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error creating buylist --> "+err.Error())
		return
	}
	if havecreate := pojo.DBCreateBuylist(buylist); !havecreate {
		c.JSON(http.StatusBadRequest, "Buylist created error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Buylist created",
		"Buylist": buylist,
	})
}

// get "/"
func FindAllBuylists(c *gin.Context) {
	buylist := pojo.DBFindAllBuylists()
	c.JSON(http.StatusOK, buylist)
}

// get "/:id"
func FindOneBuylist(c *gin.Context) {
	buylistid, _ := strconv.Atoi(c.Param("id"))
	buylist := pojo.DBFindOneBuylist(buylistid)
	if buylist.BuyListId != 0 {
		c.JSON(http.StatusOK, buylist)
		return
	}
	c.JSON(http.StatusNotFound, "Buylist not found")
}

// PUT "/:id"
func UpdateBuylist(c *gin.Context) {
	buylistid, _ := strconv.Atoi(c.Param("id"))
	buylist := pojo.Buylists{}
	c.BindJSON(&buylist)
	if dbresult := pojo.DBPutBuylist(buylistid, buylist); !dbresult {
		c.JSON(http.StatusNotFound, "Buylist not found")
		return
	}
	c.JSON(http.StatusOK, "Buylist updated")
}

// DELETE "/:id"
func DeleteBuylist(c *gin.Context) {
	buylistid, _ := strconv.Atoi(c.Param("id"))
	if dbresult := pojo.DBDeleteBuylist(buylistid); !dbresult {
		c.JSON(http.StatusNotFound, "Buylist not found")
		return
	}
	c.JSON(http.StatusOK, "Buylist Delete Success!")
}
