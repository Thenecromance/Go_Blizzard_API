package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/toy/index", ginToy_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/toy/:toyId", ginToy)

}

func ginToy_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginToy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
