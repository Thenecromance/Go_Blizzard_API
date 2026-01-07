package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/quest/index", ginQuests_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/quest/:questId", ginQuest)
	app.Instance().RegisterRoute("GET", "/data/wow/quest/category/index", ginQuest_Categories_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/quest/category/:questCategoryId", ginQuest_Category)
	app.Instance().RegisterRoute("GET", "/data/wow/quest/area/index", ginQuest_Areas_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/quest/area/:questAreaId", ginQuest_Area)
	app.Instance().RegisterRoute("GET", "/data/wow/quest/type/index", ginQuest_Types_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/quest/type/:questTypeId", ginQuest_Type)

}

func ginQuests_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginQuest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginQuest_Categories_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginQuest_Category(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginQuest_Areas_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginQuest_Area(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginQuest_Types_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginQuest_Type(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
