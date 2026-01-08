package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/D3/Community/D3ArtisanandRecipe"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/d3/data/artisan/:artisanSlug", ginD3ArtisanandRecipeGetArtisan) /* GetArtisan Returns a single artisan by slug. */

	app.Instance().RegisterRoute("GET", "/d3/data/artisan/:artisanSlug/recipe/:recipeSlug", ginD3ArtisanandRecipeGetRecipe) /* GetRecipe Returns a single recipe by slug for the specified artisan. */

}




func ginD3ArtisanandRecipeGetArtisan(c *gin.Context) {
	// binding uri parameters
	var req D3_D3ArtisanandRecipe.GetArtisanFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := D3_D3ArtisanandRecipe.GetArtisan(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginD3ArtisanandRecipeGetRecipe(c *gin.Context) {
	// binding uri parameters
	var req D3_D3ArtisanandRecipe.GetRecipeFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := D3_D3ArtisanandRecipe.GetRecipe(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
