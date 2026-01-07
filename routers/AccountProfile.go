package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wowClassic/ProfileService/AccountProfile"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/profile/user/wow", ginAccountProfileSummary) //AccountProfileSummary Returns a profile summary for an account.<br/><br/>Because this endpoint provides data about the current logged-in user's World of Warcraft account, it requires an access token with the <strong>wow.profile</strong> scope acquired via the <a href="/documentation/guides/using-oauth/authorization-code-flow">Authorization Code Flow</a>.

	app.Instance().RegisterRoute("GET", "/profile/user/wow/protected-character/:realmId-characterId", ginProtectedCharacterProfileSummary) //ProtectedCharacterProfileSummary Returns a protected profile summary for a character.<br/><br/>Because this endpoint provides data about the current logged-in user's World of Warcraft account, it requires an access token with the <strong>wow.profile</strong> scope acquired via the <a href="/documentation/guides/using-oauth/authorization-code-flow">Authorization Code Flow</a>.

}




func ginAccountProfileSummary(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_AccountProfile.AccountProfileSummaryFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_AccountProfile.AccountProfileSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginProtectedCharacterProfileSummary(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_AccountProfile.ProtectedCharacterProfileSummaryFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_AccountProfile.ProtectedCharacterProfileSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
