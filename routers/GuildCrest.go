package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/guild-crest/index", ginGuild_Crest_Components_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/media/guild-crest/border/:borderId", ginGuild_Crest_Border_Media)
	app.Instance().RegisterRoute("GET", "/data/wow/media/guild-crest/emblem/:emblemId", ginGuild_Crest_Emblem_Media)

}

func ginGuild_Crest_Components_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginGuild_Crest_Border_Media(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginGuild_Crest_Emblem_Media(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
