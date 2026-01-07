package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/DataService/AzeriteEssence"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/azerite-essence/index", ginAzeriteEssencesIndex) //AzeriteEssencesIndex Returns an index of azerite essences.

	app.Instance().RegisterRoute("GET", "/data/wow/azerite-essence/:azeriteEssenceId", ginAzeriteEssence) //AzeriteEssence Returns an azerite essence by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/search/azerite-essence", ginAzeriteEssenceSearch) //AzeriteEssenceSearch Performs a search of azerite essences. The fields below are provided for example. For more detail see the <a href="/documentation/world-of-warcraft/guides/search">Search Guide</a>.

	app.Instance().RegisterRoute("GET", "/data/wow/media/azerite-essence/:azeriteEssenceId", ginAzeriteEssenceMedia) //AzeriteEssenceMedia Returns media for an azerite essence by ID.

}




func ginAzeriteEssencesIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_AzeriteEssence.AzeriteEssencesIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_AzeriteEssence.AzeriteEssencesIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginAzeriteEssence(c *gin.Context) {
	// binding uri parameters
	var req wow_AzeriteEssence.AzeriteEssenceFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_AzeriteEssence.AzeriteEssence(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginAzeriteEssenceSearch(c *gin.Context) {
	// binding uri parameters
	var req wow_AzeriteEssence.AzeriteEssenceSearchFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_AzeriteEssence.AzeriteEssenceSearch(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginAzeriteEssenceMedia(c *gin.Context) {
	// binding uri parameters
	var req wow_AzeriteEssence.AzeriteEssenceMediaFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_AzeriteEssence.AzeriteEssenceMedia(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
