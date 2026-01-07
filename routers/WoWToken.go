package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/token/index", ginWoW_Token_Index_CN)

}

func ginWoW_Token_Index_CN(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
