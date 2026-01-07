package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/sc2/ladder/grandmaster/:regionId", ginGrandmaster_Leaderboard)
	app.Instance().RegisterRoute("GET", "/sc2/ladder/season/:regionId", ginladderSeason)

}

func ginGrandmaster_Leaderboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginladderSeason(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
