package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/DataService/MythicKeystoneLeaderboard"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/:connectedRealmId/mythic-leaderboard/index", ginMythicKeystoneLeaderboardsIndex) //MythicKeystoneLeaderboardsIndex Returns an index of Mythic Keystone Leaderboard dungeon instances for a connected realm.

	app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/:connectedRealmId/mythic-leaderboard/:dungeonId/period/:period", ginMythicKeystoneLeaderboard) //MythicKeystoneLeaderboard Returns a weekly Mythic Keystone Leaderboard by period.

}




func ginMythicKeystoneLeaderboardsIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_MythicKeystoneLeaderboard.MythicKeystoneLeaderboardsIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_MythicKeystoneLeaderboard.MythicKeystoneLeaderboardsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginMythicKeystoneLeaderboard(c *gin.Context) {
	// binding uri parameters
	var req wow_MythicKeystoneLeaderboard.MythicKeystoneLeaderboardFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_MythicKeystoneLeaderboard.MythicKeystoneLeaderboard(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
