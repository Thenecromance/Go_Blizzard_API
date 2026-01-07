package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/realm/index", ginRealms_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/realm/:realmSlug", ginRealm)
	app.Instance().RegisterRoute("GET", "/data/wow/search/realm", ginRealm_Search)

}

func ginRealms_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginRealm(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginRealm_Search(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
