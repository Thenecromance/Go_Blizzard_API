package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/soulbinds", ginCharacter_Soulbinds)

}

func ginCharacter_Soulbinds(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
