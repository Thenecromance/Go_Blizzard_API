package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/D3/Community/D3Profile"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/d3/profile/:account/", ginD3ProfileGetApiAccount) /* GetApiAccount Returns the specified account profile. */

	app.Instance().RegisterRoute("GET", "/d3/profile/:account/hero/:heroId", ginD3ProfileGetApiHero) /* GetApiHero Returns a single hero. */

	app.Instance().RegisterRoute("GET", "/d3/profile/:account/hero/:heroId/items", ginD3ProfileGetApiDetailedHeroItems) /* GetApiDetailedHeroItems Returns a list of items for the specified hero. */

	app.Instance().RegisterRoute("GET", "/d3/profile/:account/hero/:heroId/follower-items", ginD3ProfileGetApiDetailedFollowerItems) /* GetApiDetailedFollowerItems Returns a list of items for the specified hero's followers. */

}




func ginD3ProfileGetApiAccount(c *gin.Context) {
	// binding uri parameters
	var req D3_D3Profile.GetApiAccountFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := D3_D3Profile.GetApiAccount(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginD3ProfileGetApiHero(c *gin.Context) {
	// binding uri parameters
	var req D3_D3Profile.GetApiHeroFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := D3_D3Profile.GetApiHero(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginD3ProfileGetApiDetailedHeroItems(c *gin.Context) {
	// binding uri parameters
	var req D3_D3Profile.GetApiDetailedHeroItemsFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := D3_D3Profile.GetApiDetailedHeroItems(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginD3ProfileGetApiDetailedFollowerItems(c *gin.Context) {
	// binding uri parameters
	var req D3_D3Profile.GetApiDetailedFollowerItemsFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := D3_D3Profile.GetApiDetailedFollowerItems(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
