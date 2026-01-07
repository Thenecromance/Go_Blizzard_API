package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/mount/index", ginMounts_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/mount/:mountId", ginMount)
	app.Instance().RegisterRoute("GET", "/data/wow/search/mount", ginMount_Search)

}

func ginMounts_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginMount(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginMount_Search(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
