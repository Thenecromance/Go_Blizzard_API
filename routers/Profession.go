package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/DataService/Profession"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/profession/index", ginProfessionsIndex) //ProfessionsIndex Returns an index of professions.

	app.Instance().RegisterRoute("GET", "/data/wow/profession/:professionId", ginProfession) //Profession Returns a profession by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/media/profession/:professionId", ginProfessionMedia) //ProfessionMedia Returns media for a profession by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/profession/:professionId/skill-tier/:skillTierId", ginProfessionSkillTier) //ProfessionSkillTier Returns a skill tier for a profession by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/recipe/:recipeId", ginRecipe) //Recipe Returns a recipe by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/media/recipe/:recipeId", ginRecipeMedia) //RecipeMedia Returns media for a recipe by ID.

}




func ginProfessionsIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_Profession.ProfessionsIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Profession.ProfessionsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginProfession(c *gin.Context) {
	// binding uri parameters
	var req wow_Profession.ProfessionFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Profession.Profession(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginProfessionMedia(c *gin.Context) {
	// binding uri parameters
	var req wow_Profession.ProfessionMediaFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Profession.ProfessionMedia(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginProfessionSkillTier(c *gin.Context) {
	// binding uri parameters
	var req wow_Profession.ProfessionSkillTierFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Profession.ProfessionSkillTier(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginRecipe(c *gin.Context) {
	// binding uri parameters
	var req wow_Profession.RecipeFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Profession.Recipe(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginRecipeMedia(c *gin.Context) {
	// binding uri parameters
	var req wow_Profession.RecipeMediaFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Profession.RecipeMedia(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
