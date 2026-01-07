package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/d3/data/follower/:followerSlug", gingetFollower)

}

func gingetFollower(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
