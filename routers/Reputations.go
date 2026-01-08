package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/DataService/Reputations"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/reputation-faction/index", ginReputationsReputationFactionsIndex) /* ReputationFactionsIndex Returns an index of reputation factions. */

	app.Instance().RegisterRoute("GET", "/data/wow/reputation-faction/:reputationFactionId", ginReputationsReputationFaction) /* ReputationFaction Returns a single reputation faction by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/reputation-tiers/index", ginReputationsReputationTiersIndex) /* ReputationTiersIndex Returns an index of reputation tiers. */

	app.Instance().RegisterRoute("GET", "/data/wow/reputation-tiers/:reputationTiersId", ginReputationsReputationTiers) /* ReputationTiers Returns a single set of reputation tiers by ID. */

}




func ginReputationsReputationFactionsIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_Reputations.ReputationFactionsIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Reputations.ReputationFactionsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginReputationsReputationFaction(c *gin.Context) {
	// binding uri parameters
	var req wow_Reputations.ReputationFactionFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Reputations.ReputationFaction(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginReputationsReputationTiersIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_Reputations.ReputationTiersIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Reputations.ReputationTiersIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginReputationsReputationTiers(c *gin.Context) {
	// binding uri parameters
	var req wow_Reputations.ReputationTiersFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Reputations.ReputationTiers(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
