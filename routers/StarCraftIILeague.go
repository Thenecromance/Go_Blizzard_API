package routers

import (
	"net/http"

	"github.com/Thenecromance/BlizzardAPI/app"
	"github.com/Thenecromance/BlizzardAPI/api/StarCraftII/DataService/StarCraftIILeague"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/sc2/league/:seasonId/:queueId/:teamType/:leagueId", ginStarCraftIILeagueGetLeagueData) /* GetLeagueData Returns data for the specified season, queue, team, and league.

**queueId**: the standard available queueIds are: `1`=WoL 1v1, `2`=WoL 2v2, `3`=WoL 3v3, `4`=WoL 4v4, `101`=HotS 1v1, `102`=HotS 2v2, `103`=HotS 3v3, `104`=HotS 4v4, `201`=LotV 1v1, `202`=LotV 2v2, `203`=LotV 3v3, `204`=LotV 4v4, `206`=LotV Archon. Note that other available queues may not be listed here.

**teamType**: there are two available teamTypes: `0`=arranged, `1`=random.

**leagueId**: available leagueIds are: `0`=Bronze, `1`=Silver, `2`=Gold, `3`=Platinum, `4`=Diamond, `5`=Master, `6`=Grandmaster. */

}




func ginStarCraftIILeagueGetLeagueData(c *gin.Context) {
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

	// Capture all query parameters into a map if needed
    // Note: Gin's binding doesn't remove items, so this contains all query params
    queryParams := make(map[string]string)
    for k, v := range c.Request.URL.Query() {
    	if len(v) > 0 {
    		queryParams[k] = v[0]
    	}
    }
    // If the downstream function supports an extra map, pass it here.
    // Assuming req might have a field for this or you manipulate req based on queryParams.
    // For now, I'm just showing how to extract them.
    // Example: req.ExtraParams = queryParams

	resp, err := StarCraftII_StarCraftIILeague.GetLeagueData(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
