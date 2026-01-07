package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/modified-crafting/index", ginModified_Crafting_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/modified-crafting/category/index", ginModified_Crafting_Category_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/modified-crafting/category/:categoryId", ginModified_Crafting_Category)
	app.Instance().RegisterRoute("GET", "/data/wow/modified-crafting/reagent-slot-type/index", ginModified_Crafting_Reagent_Slot_Type_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/modified-crafting/reagent-slot-type/:slotTypeId", ginModified_Crafting_Reagent_Slot_Type)

}

func ginModified_Crafting_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginModified_Crafting_Category_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginModified_Crafting_Category(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginModified_Crafting_Reagent_Slot_Type_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginModified_Crafting_Reagent_Slot_Type(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
