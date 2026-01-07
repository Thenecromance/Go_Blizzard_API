package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/playable-class/index", ginPlayable_Classes_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/playable-class/:classId", ginPlayable_Class)
	app.Instance().RegisterRoute("GET", "/data/wow/media/playable-class/:playableClassId", ginPlayable_Class_Media)

}

func ginPlayable_Classes_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginPlayable_Class(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginPlayable_Class_Media(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
