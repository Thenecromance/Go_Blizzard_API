package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/playable-race/index", ginPlayable_Races_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/playable-race/:playableRaceId", ginPlayable_Race)

}

func ginPlayable_Races_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginPlayable_Race(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
