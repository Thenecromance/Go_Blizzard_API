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

	// Capture all query parameters into a map if needed
    // Note: Gin's binding doesn't remove items, so this contains all query params
    queryParams := make(map[string]string)
    for k, v := range c.Request.URL.Query() {
    	if len(v) > 0 {
    		queryParams[k] = v[0]
    	}
    }
    // If the downstream function supports an extra map, pass it here.
    // Assuming req might have a field for this or you manipulate req based on queryParams.
    // For now, I'm just showing how to extract them.
    // Example: req.ExtraParams = queryParams

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

	// Capture all query parameters into a map if needed
    // Note: Gin's binding doesn't remove items, so this contains all query params
    queryParams := make(map[string]string)
    for k, v := range c.Request.URL.Query() {
    	if len(v) > 0 {
    		queryParams[k] = v[0]
    	}
    }
    // If the downstream function supports an extra map, pass it here.
    // Assuming req might have a field for this or you manipulate req based on queryParams.
    // For now, I'm just showing how to extract them.
    // Example: req.ExtraParams = queryParams

	resp, err := wow_CharacterMythicKeystoneProfile.CharacterMythicKeystoneSeasonDetails(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
