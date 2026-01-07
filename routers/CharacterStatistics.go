package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/statistics", ginCharacter_Statistics_Summary)

}

func ginCharacter_Statistics_Summary(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
