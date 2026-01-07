package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName", ginCharacter_Profile_Summary)
	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/status", ginCharacter_Profile_Status)

}

func ginCharacter_Profile_Summary(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginCharacter_Profile_Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
