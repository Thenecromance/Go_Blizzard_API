package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/d3/data/item-type", gingetItemTypeIndex)
	app.Instance().RegisterRoute("GET", "/d3/data/item-type/:itemTypeSlug", gingetItemType)

}

func gingetItemTypeIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func gingetItemType(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
