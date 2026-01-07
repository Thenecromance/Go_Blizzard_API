package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/DataService/ItemAppearance"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/item-appearance/:appearanceId", ginItemAppearance) //ItemAppearance Returns an item appearance by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/search/item-appearance", ginItemAppearanceSearch) //ItemAppearanceSearch Performs a search of item appearances. The fields below are provided for example. For more detail see the <a href="/documentation/world-of-warcraft/guides/search">Search Guide</a>.

	app.Instance().RegisterRoute("GET", "/data/wow/item-appearance/set/index", ginItemAppearanceSetsIndex) //ItemAppearanceSetsIndex Returns an index of item appearance sets.

	app.Instance().RegisterRoute("GET", "/data/wow/item-appearance/set/:appearanceSetId", ginItemAppearanceSet) //ItemAppearanceSet Returns an item appearance set by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/item-appearance/slot/index", ginItemAppearanceSlotIndex) //ItemAppearanceSlotIndex Returns an index of item appearance slots.

	app.Instance().RegisterRoute("GET", "/data/wow/item-appearance/slot/:slotType", ginItemAppearanceSlot) //ItemAppearanceSlot Returns an item appearance slot slot type.

}




func ginItemAppearance(c *gin.Context) {
	// binding uri parameters
	var req wow_ItemAppearance.ItemAppearanceFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_ItemAppearance.ItemAppearance(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginItemAppearanceSearch(c *gin.Context) {
	// binding uri parameters
	var req wow_ItemAppearance.ItemAppearanceSearchFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_ItemAppearance.ItemAppearanceSearch(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginItemAppearanceSetsIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_ItemAppearance.ItemAppearanceSetsIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_ItemAppearance.ItemAppearanceSetsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginItemAppearanceSet(c *gin.Context) {
	// binding uri parameters
	var req wow_ItemAppearance.ItemAppearanceSetFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_ItemAppearance.ItemAppearanceSet(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginItemAppearanceSlotIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_ItemAppearance.ItemAppearanceSlotIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_ItemAppearance.ItemAppearanceSlotIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginItemAppearanceSlot(c *gin.Context) {
	// binding uri parameters
	var req wow_ItemAppearance.ItemAppearanceSlotFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_ItemAppearance.ItemAppearanceSlot(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
