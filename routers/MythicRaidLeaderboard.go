package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/DataService/MythicRaidLeaderboard"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/leaderboard/hall-of-fame/:raid/:faction", ginMythicRaidLeaderboard) //MythicRaidLeaderboard Returns the leaderboard for a given raid and faction.

}




func ginMythicRaidLeaderboard(c *gin.Context) {
	// binding uri parameters
	var req wow_MythicRaidLeaderboard.MythicRaidLeaderboardFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_MythicRaidLeaderboard.MythicRaidLeaderboard(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
