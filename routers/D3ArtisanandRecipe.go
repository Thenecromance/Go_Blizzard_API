package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/d3/data/artisan/:artisanSlug", gingetArtisan)
	app.Instance().RegisterRoute("GET", "/d3/data/artisan/:artisanSlug/recipe/:recipeSlug", gingetRecipe)

}

func gingetArtisan(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func gingetRecipe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
