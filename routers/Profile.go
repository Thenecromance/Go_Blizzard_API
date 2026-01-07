package routers

import (
	"net/http"

	"Unofficial_API/api/StarCraftII/Community/Profile"
	"Unofficial_API/app"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/sc2/static/profile/:regionId", ginStatic) //Static Returns all static SC2 profile data (achievements, categories, criteria, and rewards).

	app.Instance().RegisterRoute("GET", "/sc2/metadata/profile/:regionId/:realmId/:profileId", ginMetadata) //Metadata Returns metadata for an individual's profile.

	app.Instance().RegisterRoute("GET", "/sc2/profile/:regionId/:realmId/:profileId", ginProfile) //Profile Returns data about an individual SC2 profile.

	app.Instance().RegisterRoute("GET", "/sc2/profile/:regionId/:realmId/:profileId/ladder/summary", ginLadderSummary) //LadderSummary Returns a ladder summary for an individual SC2 profile.

	app.Instance().RegisterRoute("GET", "/sc2/profile/:regionId/:realmId/:profileId/ladder/:ladderId", ginLadder) //Ladder Returns data about an individual profile's ladder.

}

func ginStatic(c *gin.Context) {
	// binding uri parameters
	var req StarCraftII_Profile.StaticFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := StarCraftII_Profile.Static(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginMetadata(c *gin.Context) {
	// binding uri parameters
	var req StarCraftII_Profile.MetadataFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := StarCraftII_Profile.Metadata(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginProfile(c *gin.Context) {
	// binding uri parameters
	var req StarCraftII_Profile.ProfileFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := StarCraftII_Profile.Profile(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginLadderSummary(c *gin.Context) {
	// binding uri parameters
	var req StarCraftII_Profile.LadderSummaryFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := StarCraftII_Profile.LadderSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginLadder(c *gin.Context) {
	// binding uri parameters
	var req StarCraftII_Profile.LadderFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := StarCraftII_Profile.Ladder(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
