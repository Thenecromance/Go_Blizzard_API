package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wowClassic/ProfileService/CharacterSpecializations"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/specializations", ginCharacterSpecializationsCharacterSpecializationsSummary) /* CharacterSpecializationsSummary Returns a summary of a character's specializations. */

}




func ginCharacterSpecializationsCharacterSpecializationsSummary(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_CharacterSpecializations.CharacterSpecializationsSummaryFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_CharacterSpecializations.CharacterSpecializationsSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
