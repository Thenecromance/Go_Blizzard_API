package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wowClassic/ProfileService/CharacterPvP"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/pvp-bracket/:pvpBracket", ginCharacterPvPBracketStatistics) //CharacterPvPBracketStatistics Returns the PvP bracket statistics for a character.

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/pvp-summary", ginCharacterPvPSummary) //CharacterPvPSummary Returns a PvP summary for a character.

}




func ginCharacterPvPBracketStatistics(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_CharacterPvP.CharacterPvPBracketStatisticsFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_CharacterPvP.CharacterPvPBracketStatistics(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCharacterPvPSummary(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_CharacterPvP.CharacterPvPSummaryFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_CharacterPvP.CharacterPvPSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
