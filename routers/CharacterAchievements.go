package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/achievements", ginCharacter_Achievements_Summary)
	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/achievements/statistics", ginCharacter_Achievement_Statistics)

}

func ginCharacter_Achievements_Summary(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginCharacter_Achievement_Statistics(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
