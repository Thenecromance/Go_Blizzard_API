package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/DataService/PvPTier"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/pvp-tier/index", ginPvPTiersIndex) //PvPTiersIndex Returns an index of PvP tiers.

	app.Instance().RegisterRoute("GET", "/data/wow/pvp-tier/:pvpTierId", ginPvPTier) //PvPTier Returns a PvP tier by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/media/pvp-tier/:pvpTierId", ginPvPTierMedia) //PvPTierMedia Returns media for a PvP tier by ID.

}




func ginPvPTiersIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_PvPTier.PvPTiersIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_PvPTier.PvPTiersIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginPvPTier(c *gin.Context) {
	// binding uri parameters
	var req wow_PvPTier.PvPTierFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_PvPTier.PvPTier(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginPvPTierMedia(c *gin.Context) {
	// binding uri parameters
	var req wow_PvPTier.PvPTierMediaFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_PvPTier.PvPTierMedia(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
