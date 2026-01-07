package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/tech-talent-tree/index", ginTech_Talent_Tree_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/tech-talent-tree/:techTalentTreeId", ginTech_Talent_Tree)
	app.Instance().RegisterRoute("GET", "/data/wow/tech-talent/index", ginTech_Talent_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/tech-talent/:techTalentId", ginTech_Talent)
	app.Instance().RegisterRoute("GET", "/data/wow/media/tech-talent/:techTalentId", ginTech_Talent_Media)

}

func ginTech_Talent_Tree_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginTech_Talent_Tree(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginTech_Talent_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginTech_Talent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginTech_Talent_Media(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
