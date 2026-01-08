package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/ProfileService/CharacterHouse"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/house/house-{houseNumber}", ginCharacterHouseCharacterHouseSummary) /* CharacterHouseSummary Returns a summary of a house a character has built. */

}




func ginCharacterHouseCharacterHouseSummary(c *gin.Context) {
	// binding uri parameters
	var req wow_CharacterHouse.CharacterHouseSummaryFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_CharacterHouse.CharacterHouseSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
