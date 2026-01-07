package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/item-class/index", ginItem_Classes_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/item-class/:itemClassId", ginItem_Class)
	app.Instance().RegisterRoute("GET", "/data/wow/item-class/:itemClassId/item-subclass/:itemSubclassId", ginItem_Subclass)
	app.Instance().RegisterRoute("GET", "/data/wow/item/:itemId", ginItem)
	app.Instance().RegisterRoute("GET", "/data/wow/media/item/:itemId", ginItem_Media)
	app.Instance().RegisterRoute("GET", "/data/wow/search/item", ginItem_Search)

}

func ginItem_Classes_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginItem_Class(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginItem_Subclass(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginItem(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginItem_Media(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginItem_Search(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
