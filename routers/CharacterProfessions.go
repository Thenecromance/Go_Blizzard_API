package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/ProfileService/CharacterProfessions"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/professions", ginCharacterProfessionsCharacterProfessionsSummary) /* CharacterProfessionsSummary Returns a summary of professions for a character. */

}




func ginCharacterProfessionsCharacterProfessionsSummary(c *gin.Context) {
	// binding uri parameters
	var req wow_CharacterProfessions.CharacterProfessionsSummaryFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_CharacterProfessions.CharacterProfessionsSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
