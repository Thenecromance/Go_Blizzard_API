package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/StarCraftII/Community/Legacy"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/sc2/legacy/profile/:regionId/:realmId/:profileId", ginLegacyProfile) /* Profile Retrieves data about an individual SC2 profile. */

	app.Instance().RegisterRoute("GET", "/sc2/legacy/profile/:regionId/:realmId/:profileId/ladders", ginLegacyLadders) /* Ladders Retrieves data about an individual SC2 profile's ladders. */

	app.Instance().RegisterRoute("GET", "/sc2/legacy/profile/:regionId/:realmId/:profileId/matches", ginLegacyMatchHistory) /* MatchHistory Returns data about an individual SC2 profile's match history. */

	app.Instance().RegisterRoute("GET", "/sc2/legacy/ladder/:regionId/:ladderId", ginLegacyLadder) /* Ladder Returns data about an individual SC2 ladder. */

	app.Instance().RegisterRoute("GET", "/sc2/legacy/data/achievements/:regionId", ginLegacyAchievements) /* Achievements Returns data about the achievements available in SC2. */

	app.Instance().RegisterRoute("GET", "/sc2/legacy/data/rewards/:regionId", ginLegacyRewards) /* Rewards Returns data about the rewards available in SC2. */

}




func ginLegacyProfile(c *gin.Context) {
	// binding uri parameters
	var req StarCraftII_Legacy.ProfileFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := StarCraftII_Legacy.Profile(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginLegacyLadders(c *gin.Context) {
	// binding uri parameters
	var req StarCraftII_Legacy.LaddersFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := StarCraftII_Legacy.Ladders(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginLegacyMatchHistory(c *gin.Context) {
	// binding uri parameters
	var req StarCraftII_Legacy.MatchHistoryFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := StarCraftII_Legacy.MatchHistory(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginLegacyLadder(c *gin.Context) {
	// binding uri parameters
	var req StarCraftII_Legacy.LadderFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := StarCraftII_Legacy.Ladder(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginLegacyAchievements(c *gin.Context) {
	// binding uri parameters
	var req StarCraftII_Legacy.AchievementsFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := StarCraftII_Legacy.Achievements(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginLegacyRewards(c *gin.Context) {
	// binding uri parameters
	var req StarCraftII_Legacy.RewardsFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := StarCraftII_Legacy.Rewards(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
