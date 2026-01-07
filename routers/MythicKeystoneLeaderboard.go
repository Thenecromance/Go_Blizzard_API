package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/:connectedRealmId/mythic-leaderboard/index", ginMythic_Keystone_Leaderboards_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/:connectedRealmId/mythic-leaderboard/:dungeonId/period/:period", ginMythic_Keystone_Leaderboard)

}

func ginMythic_Keystone_Leaderboards_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginMythic_Keystone_Leaderboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
