package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/azerite-essence/index", ginAzerite_Essences_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/azerite-essence/:azeriteEssenceId", ginAzerite_Essence)
	app.Instance().RegisterRoute("GET", "/data/wow/search/azerite-essence", ginAzerite_Essence_Search)
	app.Instance().RegisterRoute("GET", "/data/wow/media/azerite-essence/:azeriteEssenceId", ginAzerite_Essence_Media)

}

func ginAzerite_Essences_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginAzerite_Essence(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginAzerite_Essence_Search(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginAzerite_Essence_Media(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
