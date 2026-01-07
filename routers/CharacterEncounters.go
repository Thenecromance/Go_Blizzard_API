package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/encounters", ginCharacter_Encounters_Summary)
	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/encounters/dungeons", ginCharacter_Dungeons)
	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/encounters/raids", ginCharacter_Raids)

}

func ginCharacter_Encounters_Summary(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginCharacter_Dungeons(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginCharacter_Raids(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
