package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/DataService/ModifiedCrafting"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/modified-crafting/index", ginModifiedCraftingIndex) //ModifiedCraftingIndex Returns the parent index for Modified Crafting.

	app.Instance().RegisterRoute("GET", "/data/wow/modified-crafting/category/index", ginModifiedCraftingCategoryIndex) //ModifiedCraftingCategoryIndex Returns the index of Modified Crafting categories.

	app.Instance().RegisterRoute("GET", "/data/wow/modified-crafting/category/:categoryId", ginModifiedCraftingCategory) //ModifiedCraftingCategory Returns a Modified Crafting category by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/modified-crafting/reagent-slot-type/index", ginModifiedCraftingReagentSlotTypeIndex) //ModifiedCraftingReagentSlotTypeIndex Returns the index of Modified Crafting reagent slot types.

	app.Instance().RegisterRoute("GET", "/data/wow/modified-crafting/reagent-slot-type/:slotTypeId", ginModifiedCraftingReagentSlotType) //ModifiedCraftingReagentSlotType Returns a Modified Crafting reagent slot type by ID.

}




func ginModifiedCraftingIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_ModifiedCrafting.ModifiedCraftingIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_ModifiedCrafting.ModifiedCraftingIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginModifiedCraftingCategoryIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_ModifiedCrafting.ModifiedCraftingCategoryIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_ModifiedCrafting.ModifiedCraftingCategoryIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginModifiedCraftingCategory(c *gin.Context) {
	// binding uri parameters
	var req wow_ModifiedCrafting.ModifiedCraftingCategoryFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_ModifiedCrafting.ModifiedCraftingCategory(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginModifiedCraftingReagentSlotTypeIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_ModifiedCrafting.ModifiedCraftingReagentSlotTypeIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_ModifiedCrafting.ModifiedCraftingReagentSlotTypeIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginModifiedCraftingReagentSlotType(c *gin.Context) {
	// binding uri parameters
	var req wow_ModifiedCrafting.ModifiedCraftingReagentSlotTypeFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_ModifiedCrafting.ModifiedCraftingReagentSlotType(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
