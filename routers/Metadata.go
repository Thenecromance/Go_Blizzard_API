package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/hearthstone/metadata", ginAll_metadata)
	app.Instance().RegisterRoute("GET", "/hearthstone/metadata/:type", ginSpecific_metadata)

}

func ginAll_metadata(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginSpecific_metadata(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
