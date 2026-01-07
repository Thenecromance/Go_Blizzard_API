package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/d3/profile/:account/", gingetApiAccount)
	app.Instance().RegisterRoute("GET", "/d3/profile/:account/hero/:heroId", gingetApiHero)
	app.Instance().RegisterRoute("GET", "/d3/profile/:account/hero/:heroId/items", gingetApiDetailedHeroItems)
	app.Instance().RegisterRoute("GET", "/d3/profile/:account/hero/:heroId/follower-items", gingetApiDetailedFollowerItems)

}

func gingetApiAccount(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func gingetApiHero(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func gingetApiDetailedHeroItems(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func gingetApiDetailedFollowerItems(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
