package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/ProfileService/CharacterQuests"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/quests", ginCharacterQuestsCharacterQuests) /* CharacterQuests Returns a character's active quests as well as a link to the character's completed quests. */

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/quests/completed", ginCharacterQuestsCharacterCompletedQuests) /* CharacterCompletedQuests Returns a list of quests that a character has completed. */

}




func ginCharacterQuestsCharacterQuests(c *gin.Context) {
	// binding uri parameters
	var req wow_CharacterQuests.CharacterQuestsFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_CharacterQuests.CharacterQuests(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCharacterQuestsCharacterCompletedQuests(c *gin.Context) {
	// binding uri parameters
	var req wow_CharacterQuests.CharacterCompletedQuestsFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_CharacterQuests.CharacterCompletedQuests(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
