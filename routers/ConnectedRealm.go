package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/index", ginConnected_Realms_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/:connectedRealmId", ginConnected_Realm)
	app.Instance().RegisterRoute("GET", "/data/wow/search/connected-realm", ginConnected_Realms_Search)

}

func ginConnected_Realms_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginConnected_Realm(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginConnected_Realms_Search(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
