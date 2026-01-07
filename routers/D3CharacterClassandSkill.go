package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/D3/Community/D3CharacterClassandSkill"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/d3/data/hero/:classSlug", ginGetCharacterClass) //GetCharacterClass Returns a single character class by slug.

	app.Instance().RegisterRoute("GET", "/d3/data/hero/:classSlug/skill/:skillSlug", ginGetApiSkill) //GetApiSkill Returns a single skill by slug for a specific character class.

}




func ginGetCharacterClass(c *gin.Context) {
	// binding uri parameters
	var req D3_D3CharacterClassandSkill.GetCharacterClassFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := D3_D3CharacterClassandSkill.GetCharacterClass(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginGetApiSkill(c *gin.Context) {
	// binding uri parameters
	var req D3_D3CharacterClassandSkill.GetApiSkillFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := D3_D3CharacterClassandSkill.GetApiSkill(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
