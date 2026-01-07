package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/reputation-faction/index", ginReputation_Factions_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/reputation-faction/:reputationFactionId", ginReputation_Faction)
	app.Instance().RegisterRoute("GET", "/data/wow/reputation-tiers/index", ginReputation_Tiers_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/reputation-tiers/:reputationTiersId", ginReputation_Tiers)

}

func ginReputation_Factions_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginReputation_Faction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginReputation_Tiers_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginReputation_Tiers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
