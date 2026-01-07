package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/d3/data/act", gingetActIndex)
	app.Instance().RegisterRoute("GET", "/d3/data/act/:actId", gingetAct)

}

func gingetActIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func gingetAct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
