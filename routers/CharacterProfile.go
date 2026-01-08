package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wowClassic/ProfileService/CharacterProfile"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName", ginCharacterProfileCharacterProfileSummary) /* CharacterProfileSummary Returns a profile summary for a character. */

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/status", ginCharacterProfileCharacterProfileStatus) /* CharacterProfileStatus Returns the status and a unique ID for a character. A client should delete information about a character from their application if any of the following conditions occur:<ul><li>an HTTP 404 Not Found error is returned</li><li>the is_valid value is false</li><li>the returned character ID doesn't match the previously recorded value for the character</li></ul>The following example illustrates how to use this endpoint:<ol><li>A client requests and stores information about a character, including its unique character ID and the timestamp of the request.</li><li>After 30 days, the client makes a request to the status endpoint to verify if the character information is still valid.</li><li>If character cannot be found, is not valid, or the characters IDs do not match, the client removes the information from their application.</li><li>If the character is valid and the character IDs match, the client retains the data for another 30 days.</li></ol> */

}




func ginCharacterProfileCharacterProfileSummary(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_CharacterProfile.CharacterProfileSummaryFields
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

	resp, err := wowClassic_CharacterProfile.CharacterProfileSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCharacterProfileCharacterProfileStatus(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_CharacterProfile.CharacterProfileStatusFields
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

	resp, err := wowClassic_CharacterProfile.CharacterProfileStatus(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
