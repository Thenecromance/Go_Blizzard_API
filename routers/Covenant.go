package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/covenant/index", ginCovenant_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/covenant/:covenantId", ginCovenant)
	app.Instance().RegisterRoute("GET", "/data/wow/media/covenant/:covenantId", ginCovenant_Media)
	app.Instance().RegisterRoute("GET", "/data/wow/covenant/soulbind/index", ginSoulbind_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/covenant/soulbind/:soulbindId", ginSoulbind)
	app.Instance().RegisterRoute("GET", "/data/wow/covenant/conduit/index", ginConduit_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/covenant/conduit/:conduitId", ginConduit)

}

func ginCovenant_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginCovenant(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginCovenant_Media(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginSoulbind_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginSoulbind(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginConduit_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginConduit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
