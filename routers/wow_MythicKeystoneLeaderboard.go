package routers

import (
	"net/http"

	"github.com/Thenecromance/Go_Blizzard_API/api/wow/DataService/MythicKeystoneLeaderboard"
	"github.com/Thenecromance/Go_Blizzard_API/app"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/:connectedRealmId/mythic-leaderboard/index", ginwowMythicKeystoneLeaderboardsIndex) /* MythicKeystoneLeaderboardsIndex Returns an index of Mythic Keystone Leaderboard dungeon instances for a connected realm. */

	app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/:connectedRealmId/mythic-leaderboard/:dungeonId/period/:period", ginwowMythicKeystoneLeaderboard) /* MythicKeystoneLeaderboard Returns a weekly Mythic Keystone Leaderboard by period. */

}

func ginwowMythicKeystoneLeaderboardsIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_MythicKeystoneLeaderboard.MythicKeystoneLeaderboardsIndexFields
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

	resp, err := wow_MythicKeystoneLeaderboard.MythicKeystoneLeaderboardsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginwowMythicKeystoneLeaderboard(c *gin.Context) {
	// binding uri parameters
	var req wow_MythicKeystoneLeaderboard.MythicKeystoneLeaderboardFields
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

	resp, err := wow_MythicKeystoneLeaderboard.MythicKeystoneLeaderboard(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
