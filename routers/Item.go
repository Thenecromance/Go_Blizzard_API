package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wowClassic/DataService/Item"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/item-class/index", ginItemItemClassesIndex) /* ItemClassesIndex Returns an index of item classes. */

	app.Instance().RegisterRoute("GET", "/data/wow/item-class/:itemClassId", ginItemItemClass) /* ItemClass Returns an item class by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/item-class/:itemClassId/item-subclass/:itemSubclassId", ginItemItemSubclass) /* ItemSubclass Returns an item subclass by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/item/:itemId", ginItemItem) /* Item Returns an item by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/media/item/:itemId", ginItemItemMedia) /* ItemMedia Returns media for an item by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/search/item", ginItemItemSearch) /* ItemSearch Performs a search of items. The fields below are provided for example. For more detail see the <a href="/documentation/world-of-warcraft/guides/search">Search Guide</a>. */

}




func ginItemItemClassesIndex(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_Item.ItemClassesIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_Item.ItemClassesIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginItemItemClass(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_Item.ItemClassFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_Item.ItemClass(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginItemItemSubclass(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_Item.ItemSubclassFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_Item.ItemSubclass(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginItemItem(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_Item.ItemFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_Item.Item(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginItemItemMedia(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_Item.ItemMediaFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_Item.ItemMedia(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginItemItemSearch(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_Item.ItemSearchFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_Item.ItemSearch(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
