package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/title/index", ginTitles_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/title/:titleId", ginTitle)

}

func ginTitles_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginTitle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
