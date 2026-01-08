package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/HeartStone/DataService/Metadata"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/hearthstone/metadata", ginMetadataAllmetadata) /* Allmetadata Returns information about the categorization of cards. Metadata includes the card set, set group (for example, Standard or Year of the Dragon), rarity, class, card type, minion type, and keywords. For more information, see the <a href='/documentation/hearthstone/guides/metadata'>Metadata Guide</a>. */

	app.Instance().RegisterRoute("GET", "/hearthstone/metadata/:type", ginMetadataSpecificmetadata) /* Specificmetadata Returns information about just one type of metadata. For more information, see the <a href='/documentation/hearthstone/guides/metadata'>Metadata Guide</a>. */

}




func ginMetadataAllmetadata(c *gin.Context) {
	// binding uri parameters
	var req HeartStone_Metadata.AllmetadataFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := HeartStone_Metadata.Allmetadata(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginMetadataSpecificmetadata(c *gin.Context) {
	// binding uri parameters
	var req HeartStone_Metadata.SpecificmetadataFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := HeartStone_Metadata.Specificmetadata(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
