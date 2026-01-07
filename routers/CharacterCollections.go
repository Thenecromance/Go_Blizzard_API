package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/ProfileService/CharacterCollections"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/collections", ginCharacterCollectionsIndex) //CharacterCollectionsIndex Returns an index of collection types for a character.

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/collections/decor", ginCharacterDecorCollectionSummary) //CharacterDecorCollectionSummary Returns a summary of the housing decor collection a character has obtained.

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/collections/heirlooms", ginCharacterHeirloomsCollectionSummary) //CharacterHeirloomsCollectionSummary Returns a summary of the heirlooms a character has obtained.

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/collections/mounts", ginCharacterMountsCollectionSummary) //CharacterMountsCollectionSummary Returns a summary of the mounts a character has obtained.

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/collections/pets", ginCharacterPetsCollectionSummary) //CharacterPetsCollectionSummary Returns a summary of the battle pets a character has obtained.

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/collections/toys", ginCharacterToysCollectionSummary) //CharacterToysCollectionSummary Returns a summary of the toys a character has obtained.

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/collections/transmogs", ginCharacterTransmogCollectionSummary) //CharacterTransmogCollectionSummary Returns a summary of the transmog unlocks a character has obtained.

}




func ginCharacterCollectionsIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_CharacterCollections.CharacterCollectionsIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_CharacterCollections.CharacterCollectionsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCharacterDecorCollectionSummary(c *gin.Context) {
	// binding uri parameters
	var req wow_CharacterCollections.CharacterDecorCollectionSummaryFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_CharacterCollections.CharacterDecorCollectionSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCharacterHeirloomsCollectionSummary(c *gin.Context) {
	// binding uri parameters
	var req wow_CharacterCollections.CharacterHeirloomsCollectionSummaryFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_CharacterCollections.CharacterHeirloomsCollectionSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCharacterMountsCollectionSummary(c *gin.Context) {
	// binding uri parameters
	var req wow_CharacterCollections.CharacterMountsCollectionSummaryFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_CharacterCollections.CharacterMountsCollectionSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCharacterPetsCollectionSummary(c *gin.Context) {
	// binding uri parameters
	var req wow_CharacterCollections.CharacterPetsCollectionSummaryFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_CharacterCollections.CharacterPetsCollectionSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCharacterToysCollectionSummary(c *gin.Context) {
	// binding uri parameters
	var req wow_CharacterCollections.CharacterToysCollectionSummaryFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_CharacterCollections.CharacterToysCollectionSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCharacterTransmogCollectionSummary(c *gin.Context) {
	// binding uri parameters
	var req wow_CharacterCollections.CharacterTransmogCollectionSummaryFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_CharacterCollections.CharacterTransmogCollectionSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
