package D3_routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/d3/data/act", gingetActIndex)
	app.Instance().RegisterRoute("GET", "/d3/data/act/:actId", gingetAct)
	app.Instance().RegisterRoute("GET", "/d3/data/artisan/:artisanSlug", gingetArtisan)
	app.Instance().RegisterRoute("GET", "/d3/data/artisan/:artisanSlug/recipe/:recipeSlug", gingetRecipe)
	app.Instance().RegisterRoute("GET", "/d3/data/follower/:followerSlug", gingetFollower)
	app.Instance().RegisterRoute("GET", "/d3/data/hero/:classSlug", gingetCharacterClass)
	app.Instance().RegisterRoute("GET", "/d3/data/hero/:classSlug/skill/:skillSlug", gingetApiSkill)
	app.Instance().RegisterRoute("GET", "/d3/data/item-type", gingetItemTypeIndex)
	app.Instance().RegisterRoute("GET", "/d3/data/item-type/:itemTypeSlug", gingetItemType)
	app.Instance().RegisterRoute("GET", "/d3/data/item/:itemSlugAndId", gingetItem)
	app.Instance().RegisterRoute("GET", "/d3/profile/:account/", gingetApiAccount)
	app.Instance().RegisterRoute("GET", "/d3/profile/:account/hero/:heroId", gingetApiHero)
	app.Instance().RegisterRoute("GET", "/d3/profile/:account/hero/:heroId/items", gingetApiDetailedHeroItems)
	app.Instance().RegisterRoute("GET", "/d3/profile/:account/hero/:heroId/follower-items", gingetApiDetailedFollowerItems)

}

func gingetActIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func gingetAct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func gingetArtisan(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func gingetRecipe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func gingetFollower(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func gingetCharacterClass(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func gingetApiSkill(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func gingetItemTypeIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func gingetItemType(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func gingetItem(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
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
