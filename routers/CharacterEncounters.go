package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/ProfileService/CharacterEncounters"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/encounters", ginCharacterEncountersCharacterEncountersSummary) /* CharacterEncountersSummary Returns a summary of a character's encounters. */

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/encounters/dungeons", ginCharacterEncountersCharacterDungeons) /* CharacterDungeons Returns a summary of a character's completed dungeons. */

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/encounters/raids", ginCharacterEncountersCharacterRaids) /* CharacterRaids Returns a summary of a character's completed raids. */

}




func ginCharacterEncountersCharacterEncountersSummary(c *gin.Context) {
	// binding uri parameters
	var req wow_CharacterEncounters.CharacterEncountersSummaryFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_CharacterEncounters.CharacterEncountersSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCharacterEncountersCharacterDungeons(c *gin.Context) {
	// binding uri parameters
	var req wow_CharacterEncounters.CharacterDungeonsFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_CharacterEncounters.CharacterDungeons(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCharacterEncountersCharacterRaids(c *gin.Context) {
	// binding uri parameters
	var req wow_CharacterEncounters.CharacterRaidsFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_CharacterEncounters.CharacterRaids(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
