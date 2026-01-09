package routers

import (
	"net/http"

	"github.com/Thenecromance/BlizzardAPI/app"
	"github.com/Thenecromance/BlizzardAPI/api/wow/ProfileService/AccountProfile"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/profile/user/wow", ginwowAccountProfileSummary) /* AccountProfileSummary Returns a profile summary for an account.<br/><br/>Because this endpoint provides data about the current logged-in user's World of Warcraft account, it requires an access token with the <strong>wow.profile</strong> scope acquired via the <a href="/documentation/guides/using-oauth/authorization-code-flow">Authorization Code Flow</a>. */

	app.Instance().RegisterRoute("GET", "/profile/user/wow/protected-character/:realmId-characterId", ginwowProtectedCharacterProfileSummary) /* ProtectedCharacterProfileSummary Returns a protected profile summary for a character.<br/><br/>Because this endpoint provides data about the current logged-in user's World of Warcraft account, it requires an access token with the <strong>wow.profile</strong> scope acquired via the <a href="/documentation/guides/using-oauth/authorization-code-flow">Authorization Code Flow</a>. */

	app.Instance().RegisterRoute("GET", "/profile/user/wow/collections", ginwowAccountCollectionsIndex) /* AccountCollectionsIndex Returns an index of collection types for an account.<br/><br/>Because this endpoint provides data about the current logged-in user's World of Warcraft account, it requires an access token with the <strong>wow.profile</strong> scope acquired via the <a href="/documentation/guides/using-oauth/authorization-code-flow">Authorization Code Flow</a>. */

	app.Instance().RegisterRoute("GET", "/profile/user/wow/collections/decor", ginwowAccountDecorCollectionSummary) /* AccountDecorCollectionSummary Returns a summary of the housing decor an account has collected.<br/><br/>Because this endpoint provides data about the current logged-in user's World of Warcraft account, it requires an access token with the <strong>wow.profile</strong> scope acquired via the <a href="/documentation/guides/using-oauth/authorization-code-flow">Authorization Code Flow</a>. */

	app.Instance().RegisterRoute("GET", "/profile/user/wow/collections/heirlooms", ginwowAccountHeirloomsCollectionSummary) /* AccountHeirloomsCollectionSummary Returns a summary of the heirlooms an account has obtained.<br/><br/>Because this endpoint provides data about the current logged-in user's World of Warcraft account, it requires an access token with the <strong>wow.profile</strong> scope acquired via the <a href="/documentation/guides/using-oauth/authorization-code-flow">Authorization Code Flow</a>. */

	app.Instance().RegisterRoute("GET", "/profile/user/wow/collections/mounts", ginwowAccountMountsCollectionSummary) /* AccountMountsCollectionSummary Returns a summary of the mounts an account has obtained.<br/><br/>Because this endpoint provides data about the current logged-in user's World of Warcraft account, it requires an access token with the <strong>wow.profile</strong> scope acquired via the <a href="/documentation/guides/using-oauth/authorization-code-flow">Authorization Code Flow</a>. */

	app.Instance().RegisterRoute("GET", "/profile/user/wow/collections/pets", ginwowAccountPetsCollectionSummary) /* AccountPetsCollectionSummary Returns a summary of the battle pets an account has obtained.<br/><br/>Because this endpoint provides data about the current logged-in user's World of Warcraft account, it requires an access token with the <strong>wow.profile</strong> scope acquired via the <a href="/documentation/guides/using-oauth/authorization-code-flow">Authorization Code Flow</a>. */

	app.Instance().RegisterRoute("GET", "/profile/user/wow/collections/toys", ginwowAccountToysCollectionSummary) /* AccountToysCollectionSummary Returns a summary of the toys an account has obtained.<br/><br/>Because this endpoint provides data about the current logged-in user's World of Warcraft account, it requires an access token with the <strong>wow.profile</strong> scope acquired via the <a href="/documentation/guides/using-oauth/authorization-code-flow">Authorization Code Flow</a>. */

	app.Instance().RegisterRoute("GET", "/profile/user/wow/collections/transmogs", ginwowAccountTransmogCollectionSummary) /* AccountTransmogCollectionSummary Returns a summary of the transmog unlocks an account has obtained.<br/><br/>Because this endpoint provides data about the current logged-in user's World of Warcraft account, it requires an access token with the <strong>wow.profile</strong> scope acquired via the <a href="/documentation/guides/using-oauth/authorization-code-flow">Authorization Code Flow</a>. */

}




func ginwowAccountProfileSummary(c *gin.Context) {
	// binding uri parameters
	var req wow_AccountProfile.AccountProfileSummaryFields
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

	resp, err := wow_AccountProfile.AccountProfileSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginwowProtectedCharacterProfileSummary(c *gin.Context) {
	// binding uri parameters
	var req wow_AccountProfile.ProtectedCharacterProfileSummaryFields
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

	resp, err := wow_AccountProfile.ProtectedCharacterProfileSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginwowAccountCollectionsIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_AccountProfile.AccountCollectionsIndexFields
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

	resp, err := wow_AccountProfile.AccountCollectionsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginwowAccountDecorCollectionSummary(c *gin.Context) {
	// binding uri parameters
	var req wow_AccountProfile.AccountDecorCollectionSummaryFields
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

	resp, err := wow_AccountProfile.AccountDecorCollectionSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginwowAccountHeirloomsCollectionSummary(c *gin.Context) {
	// binding uri parameters
	var req wow_AccountProfile.AccountHeirloomsCollectionSummaryFields
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

	resp, err := wow_AccountProfile.AccountHeirloomsCollectionSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginwowAccountMountsCollectionSummary(c *gin.Context) {
	// binding uri parameters
	var req wow_AccountProfile.AccountMountsCollectionSummaryFields
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

	resp, err := wow_AccountProfile.AccountMountsCollectionSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginwowAccountPetsCollectionSummary(c *gin.Context) {
	// binding uri parameters
	var req wow_AccountProfile.AccountPetsCollectionSummaryFields
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

	resp, err := wow_AccountProfile.AccountPetsCollectionSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginwowAccountToysCollectionSummary(c *gin.Context) {
	// binding uri parameters
	var req wow_AccountProfile.AccountToysCollectionSummaryFields
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

	resp, err := wow_AccountProfile.AccountToysCollectionSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginwowAccountTransmogCollectionSummary(c *gin.Context) {
	// binding uri parameters
	var req wow_AccountProfile.AccountTransmogCollectionSummaryFields
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

	resp, err := wow_AccountProfile.AccountTransmogCollectionSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
