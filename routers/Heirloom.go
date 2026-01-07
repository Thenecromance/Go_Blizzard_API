package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/heirloom/index", ginHeirloom_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/heirloom/:heirloomId", ginHeirloom)

}

func ginHeirloom_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginHeirloom(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
