package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/DataService/Neighborhood"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/neighborhood-map/index", ginNeighborhoodMapIndex) //NeighborhoodMapIndex Returns an index of neighborhood maps.

	app.Instance().RegisterRoute("GET", "/data/wow/neighborhood-map/:neighborhoodMapId", ginNeighborhoodMap) //NeighborhoodMap Returns a neighborhood map by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/neighborhood-map/:neighborhoodMapId/neighborhood/:neighborhoodId", ginNeighborhood) //Neighborhood Returns a neighborhood by ID.

}




func ginNeighborhoodMapIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_Neighborhood.NeighborhoodMapIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Neighborhood.NeighborhoodMapIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginNeighborhoodMap(c *gin.Context) {
	// binding uri parameters
	var req wow_Neighborhood.NeighborhoodMapFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Neighborhood.NeighborhoodMap(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginNeighborhood(c *gin.Context) {
	// binding uri parameters
	var req wow_Neighborhood.NeighborhoodFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Neighborhood.Neighborhood(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
