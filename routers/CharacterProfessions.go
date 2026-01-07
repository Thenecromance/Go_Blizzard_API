package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/professions", ginCharacter_Professions_Summary)

}

func ginCharacter_Professions_Summary(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
