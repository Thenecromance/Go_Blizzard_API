package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/ProfileService/CharacterTitles"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/titles", ginCharacterTitlesCharacterTitlesSummary) /* CharacterTitlesSummary Returns a summary of titles a character has obtained. */

}




func ginCharacterTitlesCharacterTitlesSummary(c *gin.Context) {
	// binding uri parameters
	var req wow_CharacterTitles.CharacterTitlesSummaryFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_CharacterTitles.CharacterTitlesSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
