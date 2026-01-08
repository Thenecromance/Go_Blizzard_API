package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/StarCraftII/Community/Account"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/sc2/player/:accountId", ginAccountPlayer) /* Player Returns metadata for an individual's account. */

}




func ginAccountPlayer(c *gin.Context) {
	// binding uri parameters
	var req StarCraftII_Account.PlayerFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := StarCraftII_Account.Player(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
