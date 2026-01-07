package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/pvp-bracket/:pvpBracket", ginCharacter_PvP_Bracket_Statistics)
	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/pvp-summary", ginCharacter_PvP_Summary)

}

func ginCharacter_PvP_Bracket_Statistics(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginCharacter_PvP_Summary(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
