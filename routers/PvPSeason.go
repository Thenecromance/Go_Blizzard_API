package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/pvp-season/index", ginPvP_Seasons_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/pvp-season/:pvpSeasonId", ginPvP_Season)
	app.Instance().RegisterRoute("GET", "/data/wow/pvp-region/index", ginPvP_Region_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/pvp-region/:pvpRegionId/pvp-season/index", ginPvP_Regional_Season_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/pvp-region/:pvpRegionId/pvp-season/:pvpSeasonId", ginPvP_Regional_Season)
	app.Instance().RegisterRoute("GET", "/data/wow/pvp-region/:pvpRegionId/pvp-season/:pvpSeasonId/pvp-leaderboard/index", ginPvP_Leaderboards_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/pvp-region/:pvpRegionId/pvp-season/:pvpSeasonId/pvp-leaderboard/:pvpBracket", ginPvP_Leaderboard)
	app.Instance().RegisterRoute("GET", "/data/wow/pvp-region/:pvpRegionId/pvp-season/:pvpSeasonId/pvp-reward/index", ginPvP_Rewards_Index)

}

func ginPvP_Seasons_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginPvP_Season(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginPvP_Region_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginPvP_Regional_Season_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginPvP_Regional_Season(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginPvP_Leaderboards_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginPvP_Leaderboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginPvP_Rewards_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
