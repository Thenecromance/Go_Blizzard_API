package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/d3/data/hero/:classSlug", gingetCharacterClass)
	app.Instance().RegisterRoute("GET", "/d3/data/hero/:classSlug/skill/:skillSlug", gingetApiSkill)

}

func gingetCharacterClass(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func gingetApiSkill(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
