package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/StarCraftII/Community/Ladder"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/sc2/ladder/grandmaster/:regionId", ginGrandmasterLeaderboard) //GrandmasterLeaderboard Returns ladder data for the current season's grandmaster leaderboard.

	app.Instance().RegisterRoute("GET", "/sc2/ladder/season/:regionId", ginSeason) //Season Returns data about the current season.

}




func ginGrandmasterLeaderboard(c *gin.Context) {
	// binding uri parameters
	var req StarCraftII_Ladder.GrandmasterLeaderboardFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := StarCraftII_Ladder.GrandmasterLeaderboard(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginSeason(c *gin.Context) {
	// binding uri parameters
	var req StarCraftII_Ladder.SeasonFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := StarCraftII_Ladder.Season(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
