package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/playable-specialization/index", ginPlayable_Specializations_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/playable-specialization/:specId", ginPlayable_Specialization)
	app.Instance().RegisterRoute("GET", "/data/wow/media/playable-specialization/:specId", ginPlayable_Specialization_Media)

}

func ginPlayable_Specializations_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginPlayable_Specialization(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginPlayable_Specialization_Media(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
