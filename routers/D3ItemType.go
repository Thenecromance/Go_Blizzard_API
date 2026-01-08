package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/D3/Community/D3ItemType"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/d3/data/item-type", ginD3ItemTypeGetItemTypeIndex) /* GetItemTypeIndex Returns an index of item types. */

	app.Instance().RegisterRoute("GET", "/d3/data/item-type/:itemTypeSlug", ginD3ItemTypeGetItemType) /* GetItemType Returns a single item type by slug. */

}




func ginD3ItemTypeGetItemTypeIndex(c *gin.Context) {
	// binding uri parameters
	var req D3_D3ItemType.GetItemTypeIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := D3_D3ItemType.GetItemTypeIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginD3ItemTypeGetItemType(c *gin.Context) {
	// binding uri parameters
	var req D3_D3ItemType.GetItemTypeFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := D3_D3ItemType.GetItemType(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
