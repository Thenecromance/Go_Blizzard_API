package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/creature-family/index", ginCreature_Families_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/creature-family/:creatureFamilyId", ginCreature_Family)
	app.Instance().RegisterRoute("GET", "/data/wow/creature-type/index", ginCreature_Types_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/creature-type/:creatureTypeId", ginCreature_Type)
	app.Instance().RegisterRoute("GET", "/data/wow/creature/:creatureId", ginCreature)
	app.Instance().RegisterRoute("GET", "/data/wow/search/creature", ginCreature_Search)
	app.Instance().RegisterRoute("GET", "/data/wow/media/creature-display/:creatureDisplayId", ginCreature_Display_Media)
	app.Instance().RegisterRoute("GET", "/data/wow/media/creature-family/:creatureFamilyId", ginCreature_Family_Media)

}

func ginCreature_Families_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginCreature_Family(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginCreature_Types_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginCreature_Type(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginCreature(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginCreature_Search(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginCreature_Display_Media(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginCreature_Family_Media(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
