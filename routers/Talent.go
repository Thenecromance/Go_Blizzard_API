package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/talent-tree/index", ginTalent_Tree_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/talent-tree/:talentTreeId/playable-specialization/:specId", ginTalent_Tree)
	app.Instance().RegisterRoute("GET", "/data/wow/talent-tree/:talentTreeId", ginTalent_Tree_Nodes)
	app.Instance().RegisterRoute("GET", "/data/wow/talent/index", ginTalents_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/talent/:talentId", ginTalent)
	app.Instance().RegisterRoute("GET", "/data/wow/pvp-talent/index", ginPvP_Talents_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/pvp-talent/:pvpTalentId", ginPvP_Talent)

}

func ginTalent_Tree_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginTalent_Tree(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginTalent_Tree_Nodes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginTalents_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginTalent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginPvP_Talents_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginPvP_Talent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
