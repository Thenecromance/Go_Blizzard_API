package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/region/index", ginRegions_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/region/:regionId", ginRegion)

}

func ginRegions_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginRegion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
