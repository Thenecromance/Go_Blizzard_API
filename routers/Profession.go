package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/profession/index", ginProfessions_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/profession/:professionId", ginProfession)
	app.Instance().RegisterRoute("GET", "/data/wow/media/profession/:professionId", ginProfession_Media)
	app.Instance().RegisterRoute("GET", "/data/wow/profession/:professionId/skill-tier/:skillTierId", ginProfession_Skill_Tier)
	app.Instance().RegisterRoute("GET", "/data/wow/recipe/:recipeId", ginRecipe)
	app.Instance().RegisterRoute("GET", "/data/wow/media/recipe/:recipeId", ginRecipe_Media)

}

func ginProfessions_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginProfession(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginProfession_Media(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginProfession_Skill_Tier(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginRecipe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginRecipe_Media(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
