package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wowClassic/DataService/PvPSeason"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/pvp-season/index", ginPvPSeasonsIndex) //PvPSeasonsIndex Returns an index of PvP seasons.

	app.Instance().RegisterRoute("GET", "/data/wow/pvp-season/:pvpSeasonId", ginPvPSeason) //PvPSeason Returns a PvP season by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/pvp-region/index", ginPvPRegionIndex) //PvPRegionIndex Returns an index of PvP Regions.

	app.Instance().RegisterRoute("GET", "/data/wow/pvp-region/:pvpRegionId/pvp-season/index", ginPvPRegionalSeasonIndex) //PvPRegionalSeasonIndex Returns an index of PvP Seasons in a PvP region.

	app.Instance().RegisterRoute("GET", "/data/wow/pvp-region/:pvpRegionId/pvp-season/:pvpSeasonId", ginPvPRegionalSeason) //PvPRegionalSeason Returns a PvP season by region ID and season ID.

	app.Instance().RegisterRoute("GET", "/data/wow/pvp-region/:pvpRegionId/pvp-season/:pvpSeasonId/pvp-leaderboard/index", ginPvPLeaderboardsIndex) //PvPLeaderboardsIndex Returns an index of PvP leaderboards for a PvP season in a given PvP region.

	app.Instance().RegisterRoute("GET", "/data/wow/pvp-region/:pvpRegionId/pvp-season/:pvpSeasonId/pvp-leaderboard/:pvpBracket", ginPvPLeaderboard) //PvPLeaderboard Returns the PvP leaderboard of a specific PvP bracket for a PvP season in a given PvP region.

	app.Instance().RegisterRoute("GET", "/data/wow/pvp-region/:pvpRegionId/pvp-season/:pvpSeasonId/pvp-reward/index", ginPvPRewardsIndex) //PvPRewardsIndex Returns an index of PvP rewards for a PvP season in a given PvP region.

}




func ginPvPSeasonsIndex(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_PvPSeason.PvPSeasonsIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_PvPSeason.PvPSeasonsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginPvPSeason(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_PvPSeason.PvPSeasonFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_PvPSeason.PvPSeason(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginPvPRegionIndex(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_PvPSeason.PvPRegionIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_PvPSeason.PvPRegionIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginPvPRegionalSeasonIndex(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_PvPSeason.PvPRegionalSeasonIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_PvPSeason.PvPRegionalSeasonIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginPvPRegionalSeason(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_PvPSeason.PvPRegionalSeasonFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_PvPSeason.PvPRegionalSeason(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginPvPLeaderboardsIndex(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_PvPSeason.PvPLeaderboardsIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_PvPSeason.PvPLeaderboardsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginPvPLeaderboard(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_PvPSeason.PvPLeaderboardFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_PvPSeason.PvPLeaderboard(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginPvPRewardsIndex(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_PvPSeason.PvPRewardsIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_PvPSeason.PvPRewardsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
