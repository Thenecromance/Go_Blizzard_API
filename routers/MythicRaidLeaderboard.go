package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/leaderboard/hall-of-fame/:raid/:faction", ginMythic_Raid_Leaderboard)

}

func ginMythic_Raid_Leaderboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
