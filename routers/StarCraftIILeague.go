package routers

import (
	"net/http"

	"Unofficial_API/api/StarCraftII/DataService/StarCraftIILeague"
	"Unofficial_API/app"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/sc2/league/:seasonId/:queueId/:teamType/:leagueId", ginGetLeagueData) //GetLeagueData Returns data for the specified season, queue, team, and league.

	/***queueId**: the standard available queueIds are: `1`=WoL 1v1, `2`=WoL 2v2, `3`=WoL 3v3, `4`=WoL 4v4, `101`=HotS 1v1, `102`=HotS 2v2, `103`=HotS 3v3, `104`=HotS 4v4, `201`=LotV 1v1, `202`=LotV 2v2, `203`=LotV 3v3, `204`=LotV 4v4, `206`=LotV Archon. Note that other available queues may not be listed here.

	 **teamType**: there are two available teamTypes: `0`=arranged, `1`=random.

	 **leagueId**: available leagueIds are: `0`=Bronze, `1`=Silver, `2`=Gold, `3`=Platinum, `4`=Diamond, `5`=Master, `6`=Grandmaster.*/

}

func ginGetLeagueData(c *gin.Context) {
	// binding uri parameters
	var req StarCraftII_StarCraftIILeague.GetLeagueDataFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := StarCraftII_StarCraftIILeague.GetLeagueData(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
