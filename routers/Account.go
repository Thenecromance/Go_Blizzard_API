package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/sc2/player/:accountId", ginPlayer)

}

func ginPlayer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
