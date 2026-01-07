package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/pvp-tier/index", ginPvP_Tiers_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/pvp-tier/:pvpTierId", ginPvP_Tier)
	app.Instance().RegisterRoute("GET", "/data/wow/media/pvp-tier/:pvpTierId", ginPvP_Tier_Media)

}

func ginPvP_Tiers_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginPvP_Tier(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginPvP_Tier_Media(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
