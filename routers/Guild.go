package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/guild/:realmSlug/:nameSlug", ginGuild)
	app.Instance().RegisterRoute("GET", "/data/wow/guild/:realmSlug/:nameSlug/activity", ginGuild_Activity)
	app.Instance().RegisterRoute("GET", "/data/wow/guild/:realmSlug/:nameSlug/achievements", ginGuild_Achievements)
	app.Instance().RegisterRoute("GET", "/data/wow/guild/:realmSlug/:nameSlug/roster", ginGuild_Roster)

}

func ginGuild(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginGuild_Activity(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginGuild_Achievements(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginGuild_Roster(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
