package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/search/media", ginMedia_Search)

}

func ginMedia_Search(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
