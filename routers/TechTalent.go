package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/DataService/TechTalent"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/tech-talent-tree/index", ginTechTalentTreeIndex) //TechTalentTreeIndex Returns an index of tech talent trees.

	app.Instance().RegisterRoute("GET", "/data/wow/tech-talent-tree/:techTalentTreeId", ginTechTalentTree) //TechTalentTree Returns a tech talent tree by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/tech-talent/index", ginTechTalentIndex) //TechTalentIndex Returns an index of tech talents.

	app.Instance().RegisterRoute("GET", "/data/wow/tech-talent/:techTalentId", ginTechTalent) //TechTalent Returns a tech talent by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/media/tech-talent/:techTalentId", ginTechTalentMedia) //TechTalentMedia Returns media for a tech talent by ID.

}




func ginTechTalentTreeIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_TechTalent.TechTalentTreeIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_TechTalent.TechTalentTreeIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginTechTalentTree(c *gin.Context) {
	// binding uri parameters
	var req wow_TechTalent.TechTalentTreeFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_TechTalent.TechTalentTree(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginTechTalentIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_TechTalent.TechTalentIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_TechTalent.TechTalentIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginTechTalent(c *gin.Context) {
	// binding uri parameters
	var req wow_TechTalent.TechTalentFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_TechTalent.TechTalent(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginTechTalentMedia(c *gin.Context) {
	// binding uri parameters
	var req wow_TechTalent.TechTalentMediaFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_TechTalent.TechTalentMedia(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
