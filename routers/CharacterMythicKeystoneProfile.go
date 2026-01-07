package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/mythic-keystone-profile", ginCharacter_Mythic_Keystone_Profile_Index)
	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/mythic-keystone-profile/season/:seasonId", ginCharacter_Mythic_Keystone_Season_Details)

}

func ginCharacter_Mythic_Keystone_Profile_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginCharacter_Mythic_Keystone_Season_Details(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
