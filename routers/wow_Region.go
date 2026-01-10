package routers

import (
	"net/http"

	"github.com/Thenecromance/Go_Blizzard_API/api/wow/DataService/Region"
	"github.com/Thenecromance/Go_Blizzard_API/app"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/region/index", ginwowRegionsIndex) /* RegionsIndex Returns an index of regions. */

	app.Instance().RegisterRoute("GET", "/data/wow/region/:regionId", ginwowRegion) /* Region Returns a region by ID. */

}

func ginwowRegionsIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_Region.RegionsIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// Capture all query parameters into a map if needed
	// Note: Gin's binding doesn't remove items, so this contains all query params
	queryParams := make(map[string]string)
	for k, v := range c.Request.URL.Query() {
		if len(v) > 0 {
			queryParams[k] = v[0]
		}
	}
	// If the downstream function supports an extra map, pass it here.
	// Assuming req might have a field for this or you manipulate req based on queryParams.
	// For now, I'm just showing how to extract them.
	// Example: req.ExtraParams = queryParams

	resp, err := wow_Region.RegionsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginwowRegion(c *gin.Context) {
	// binding uri parameters
	var req wow_Region.RegionFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// Capture all query parameters into a map if needed
	// Note: Gin's binding doesn't remove items, so this contains all query params
	queryParams := make(map[string]string)
	for k, v := range c.Request.URL.Query() {
		if len(v) > 0 {
			queryParams[k] = v[0]
		}
	}
	// If the downstream function supports an extra map, pass it here.
	// Assuming req might have a field for this or you manipulate req based on queryParams.
	// For now, I'm just showing how to extract them.
	// Example: req.ExtraParams = queryParams

	resp, err := wow_Region.Region(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
