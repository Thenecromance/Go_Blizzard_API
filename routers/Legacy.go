package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/sc2/legacy/profile/:regionId/:realmId/:profileId", ginlegacyProfile)
	app.Instance().RegisterRoute("GET", "/sc2/legacy/profile/:regionId/:realmId/:profileId/ladders", ginlegacyLadder)
	app.Instance().RegisterRoute("GET", "/sc2/legacy/profile/:regionId/:realmId/:profileId/matches", ginMatch_History)
	app.Instance().RegisterRoute("GET", "/sc2/legacy/ladder/:regionId/:ladderId", ginlegacyLadders)
	app.Instance().RegisterRoute("GET", "/sc2/legacy/data/achievements/:regionId", ginAchievements)
	app.Instance().RegisterRoute("GET", "/sc2/legacy/data/rewards/:regionId", ginRewards)

}

func ginlegacyProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginlegacyLadders(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginMatch_History(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginlegacyLadder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginAchievements(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginRewards(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
