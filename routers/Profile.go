package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/sc2/static/profile/:regionId", ginStatic)
	app.Instance().RegisterRoute("GET", "/sc2/metadata/profile/:regionId/:realmId/:profileId", ginMetadata)
	app.Instance().RegisterRoute("GET", "/sc2/profile/:regionId/:realmId/:profileId", ginProfile)
	app.Instance().RegisterRoute("GET", "/sc2/profile/:regionId/:realmId/:profileId/ladder/summary", ginLadder_Summary)
	app.Instance().RegisterRoute("GET", "/sc2/profile/:regionId/:realmId/:profileId/ladder/:ladderId", ginLadder)

}

func ginStatic(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginMetadata(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginLadder_Summary(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginLadder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
