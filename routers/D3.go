package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/d3/season/", ginSeason_Index)
	app.Instance().RegisterRoute("GET", "/data/d3/season/:id", ginSeason)
	app.Instance().RegisterRoute("GET", "/data/d3/season/:id/leaderboard/:leaderboard", ginSeason_Leaderboard)
	app.Instance().RegisterRoute("GET", "/data/d3/era/", ginEra_Index)
	app.Instance().RegisterRoute("GET", "/data/d3/era/:id", ginEra)
	app.Instance().RegisterRoute("GET", "/data/d3/era/:id/leaderboard/:leaderboard", ginEra_Leaderboard)

}

func ginSeason_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginSeason(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginSeason_Leaderboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginEra_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginEra(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginEra_Leaderboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
