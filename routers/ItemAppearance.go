package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/item-appearance/:appearanceId", ginItem_Appearance)
	app.Instance().RegisterRoute("GET", "/data/wow/search/item-appearance", ginItem_Appearance_Search)
	app.Instance().RegisterRoute("GET", "/data/wow/item-appearance/set/index", ginItem_Appearance_Sets_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/item-appearance/set/:appearanceSetId", ginItem_Appearance_Set)
	app.Instance().RegisterRoute("GET", "/data/wow/item-appearance/slot/index", ginItem_Appearance_Slot_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/item-appearance/slot/:slotType", ginItem_Appearance_Slot)

}

func ginItem_Appearance(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginItem_Appearance_Search(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginItem_Appearance_Sets_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginItem_Appearance_Set(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginItem_Appearance_Slot_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginItem_Appearance_Slot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
