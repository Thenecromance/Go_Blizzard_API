package routers

import (
	"net/http"

	"github.com/Thenecromance/BlizzardAPI/app"
	"github.com/Thenecromance/BlizzardAPI/api/wow/DataService/AuctionHouse"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/:connectedRealmId/auctions", ginwowAuctions) /* Auctions Returns all active auctions for a connected realm.<br/><br/>See the <strong>Connected Realm API</strong> for information about retrieving a list of connected realm IDs.<br/><br/>Auction house data updates at a set interval. The value was initially set at 1 hour; however, it might change over time without notice.<br/><br/>Depending on the number of active auctions on the specified connected realm, the response from this endpoint may be rather large, sometimes exceeding 10 MB. */

	app.Instance().RegisterRoute("GET", "/data/wow/auctions/commodities", ginwowCommodities) /* Commodities Returns all active auctions for commodity items for the entire game region.<br/><br/>Auction house data updates at a set interval. The value was initially set at 1 hour; however, it might change over time without notice.<br/><br/>Depending on the number of active auctions on the specified connected realm, the response from this endpoint may be rather large, sometimes exceeding 10 MB. */

}




func ginwowAuctions(c *gin.Context) {
	// binding uri parameters
	var req wow_AuctionHouse.AuctionsFields
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

	resp, err := wow_AuctionHouse.Auctions(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginwowCommodities(c *gin.Context) {
	// binding uri parameters
	var req wow_AuctionHouse.CommoditiesFields
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

	resp, err := wow_AuctionHouse.Commodities(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
