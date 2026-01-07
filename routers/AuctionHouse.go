package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/:connectedRealmId/auctions/index", ginAuction_House_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/:connectedRealmId/auctions/:auctionHouseId", ginAuctions)

}

func ginAuction_House_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginAuctions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
