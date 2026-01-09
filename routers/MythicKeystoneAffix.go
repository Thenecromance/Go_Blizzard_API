package routers

import (
	"net/http"

	"github.com/Thenecromance/BlizzardAPI/app"
	"github.com/Thenecromance/BlizzardAPI/api/wow/DataService/MythicKeystoneAffix"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/keystone-affix/index", ginMythicKeystoneAffixMythicKeystoneAffixesIndex) /* MythicKeystoneAffixesIndex Returns an index of mythic keystone affixes. */

	app.Instance().RegisterRoute("GET", "/data/wow/keystone-affix/:keystoneAffixId", ginMythicKeystoneAffixMythicKeystoneAffix) /* MythicKeystoneAffix Returns a mythic keystone affix by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/media/keystone-affix/:keystoneAffixId", ginMythicKeystoneAffixMythicKeystoneAffixMedia) /* MythicKeystoneAffixMedia Returns media for a mythic keystone affix by ID. */

}




func ginMythicKeystoneAffixMythicKeystoneAffixesIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_MythicKeystoneAffix.MythicKeystoneAffixesIndexFields
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

	resp, err := wow_MythicKeystoneAffix.MythicKeystoneAffixesIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginMythicKeystoneAffixMythicKeystoneAffix(c *gin.Context) {
	// binding uri parameters
	var req wow_MythicKeystoneAffix.MythicKeystoneAffixFields
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

	resp, err := wow_MythicKeystoneAffix.MythicKeystoneAffix(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginMythicKeystoneAffixMythicKeystoneAffixMedia(c *gin.Context) {
	// binding uri parameters
	var req wow_MythicKeystoneAffix.MythicKeystoneAffixMediaFields
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

	resp, err := wow_MythicKeystoneAffix.MythicKeystoneAffixMedia(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
