package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/DataService/ItemAppearance"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/item-appearance/:appearanceId", ginItemAppearanceItemAppearance) /* ItemAppearance Returns an item appearance by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/search/item-appearance", ginItemAppearanceItemAppearanceSearch) /* ItemAppearanceSearch Performs a search of item appearances. The fields below are provided for example. For more detail see the <a href="/documentation/world-of-warcraft/guides/search">Search Guide</a>. */

	app.Instance().RegisterRoute("GET", "/data/wow/item-appearance/set/index", ginItemAppearanceItemAppearanceSetsIndex) /* ItemAppearanceSetsIndex Returns an index of item appearance sets. */

	app.Instance().RegisterRoute("GET", "/data/wow/item-appearance/set/:appearanceSetId", ginItemAppearanceItemAppearanceSet) /* ItemAppearanceSet Returns an item appearance set by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/item-appearance/slot/index", ginItemAppearanceItemAppearanceSlotIndex) /* ItemAppearanceSlotIndex Returns an index of item appearance slots. */

	app.Instance().RegisterRoute("GET", "/data/wow/item-appearance/slot/:slotType", ginItemAppearanceItemAppearanceSlot) /* ItemAppearanceSlot Returns an item appearance slot slot type. */

}




func ginItemAppearanceItemAppearance(c *gin.Context) {
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



func ginItemAppearanceItemAppearanceSearch(c *gin.Context) {
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



func ginItemAppearanceItemAppearanceSetsIndex(c *gin.Context) {
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



func ginItemAppearanceItemAppearanceSet(c *gin.Context) {
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



func ginItemAppearanceItemAppearanceSlotIndex(c *gin.Context) {
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



func ginItemAppearanceItemAppearanceSlot(c *gin.Context) {
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
