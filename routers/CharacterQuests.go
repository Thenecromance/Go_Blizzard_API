package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/quests", ginCharacter_Quests)
	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/quests/completed", ginCharacter_Completed_Quests)

}

func ginCharacter_Quests(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginCharacter_Completed_Quests(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
