package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/power-type/index", ginPower_Types_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/power-type/:powerTypeId", ginPower_Type)

}

func ginPower_Types_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginPower_Type(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
