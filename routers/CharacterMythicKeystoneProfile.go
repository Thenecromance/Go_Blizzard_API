package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/ProfileService/CharacterMythicKeystoneProfile"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/mythic-keystone-profile", ginCharacterMythicKeystoneProfileCharacterMythicKeystoneProfileIndex) /* CharacterMythicKeystoneProfileIndex Returns the Mythic Keystone profile index for a character. */

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/mythic-keystone-profile/season/:seasonId", ginCharacterMythicKeystoneProfileCharacterMythicKeystoneSeasonDetails) /* CharacterMythicKeystoneSeasonDetails Returns the Mythic Keystone season details for a character.<br/><br/>Returns a <strong>404 Not Found</strong> for characters that have not yet completed a Mythic Keystone dungeon for the specified season. */

}




func ginCharacterMythicKeystoneProfileCharacterMythicKeystoneProfileIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_CharacterMythicKeystoneProfile.CharacterMythicKeystoneProfileIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_CharacterMythicKeystoneProfile.CharacterMythicKeystoneProfileIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCharacterMythicKeystoneProfileCharacterMythicKeystoneSeasonDetails(c *gin.Context) {
	// binding uri parameters
	var req wow_CharacterMythicKeystoneProfile.CharacterMythicKeystoneSeasonDetailsFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_CharacterMythicKeystoneProfile.CharacterMythicKeystoneSeasonDetails(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
