package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wowClassic/DataService/AuctionHouse"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/:connectedRealmId/auctions/index", ginAuctionHouseIndex) //AuctionHouseIndex Returns an index of auction houses for a connected realm.<br/><br/>See the <strong>Connected Realm API</strong> for information about retrieving a list of connected realm IDs.

	app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/:connectedRealmId/auctions/:auctionHouseId", ginAuctions) //Auctions Returns all active auctions for a specific auction house on a connected realm.<br/><br/>See the <strong>Connected Realm API</strong> for information about retrieving a list of connected realm IDs.<br/><br/>Auction house data updates at a set interval. The value was initially set at 1 hour; however, it might change over time without notice.<br/><br/>Depending on the number of active auctions on the specified connected realm, the response from this endpoint may be rather large, sometimes exceeding 10 MB.

}




func ginAuctionHouseIndex(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_AuctionHouse.AuctionHouseIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_AuctionHouse.AuctionHouseIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginAuctions(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_AuctionHouse.AuctionsFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_AuctionHouse.Auctions(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
