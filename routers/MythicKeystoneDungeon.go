package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/index", ginMythic_Keystone_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/dungeon/index", ginMythic_Keystone_Dungeons_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/dungeon/:dungeonId", ginMythic_Keystone_Dungeon)
	app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/period/index", ginMythic_Keystone_Periods_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/period/:periodId", ginMythic_Keystone_Period)
	app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/season/index", ginMythic_Keystone_Seasons_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/season/:seasonId", ginMythic_Keystone_Season)

}

func ginMythic_Keystone_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginMythic_Keystone_Dungeons_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginMythic_Keystone_Dungeon(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginMythic_Keystone_Periods_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginMythic_Keystone_Period(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginMythic_Keystone_Seasons_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginMythic_Keystone_Season(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
