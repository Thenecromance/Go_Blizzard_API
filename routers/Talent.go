package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/DataService/Talent"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/talent-tree/index", ginTalentTreeIndex) //TalentTreeIndex Returns an index of talent trees.

	app.Instance().RegisterRoute("GET", "/data/wow/talent-tree/:talentTreeId/playable-specialization/:specId", ginTalentTree) //TalentTree Returns a talent tree by specialization ID.

	app.Instance().RegisterRoute("GET", "/data/wow/talent-tree/:talentTreeId", ginTalentTreeNodes) //TalentTreeNodes Returns all talent tree nodes as well as links to associated playable specializations given a talent tree id. This is useful to generate loadout export codes.

	app.Instance().RegisterRoute("GET", "/data/wow/talent/index", ginTalentsIndex) //TalentsIndex Returns an index of talents.

	app.Instance().RegisterRoute("GET", "/data/wow/talent/:talentId", ginTalent) //Talent Returns a talent by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/pvp-talent/index", ginPvPTalentsIndex) //PvPTalentsIndex Returns an index of PvP talents.

	app.Instance().RegisterRoute("GET", "/data/wow/pvp-talent/:pvpTalentId", ginPvPTalent) //PvPTalent Returns a PvP talent by ID.

}




func ginTalentTreeIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_Talent.TalentTreeIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Talent.TalentTreeIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginTalentTree(c *gin.Context) {
	// binding uri parameters
	var req wow_Talent.TalentTreeFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Talent.TalentTree(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginTalentTreeNodes(c *gin.Context) {
	// binding uri parameters
	var req wow_Talent.TalentTreeNodesFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Talent.TalentTreeNodes(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginTalentsIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_Talent.TalentsIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Talent.TalentsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginTalent(c *gin.Context) {
	// binding uri parameters
	var req wow_Talent.TalentFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Talent.Talent(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginPvPTalentsIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_Talent.PvPTalentsIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Talent.PvPTalentsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginPvPTalent(c *gin.Context) {
	// binding uri parameters
	var req wow_Talent.PvPTalentFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Talent.PvPTalent(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
