package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wowClassic/ProfileService/CharacterEquipment"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/equipment", ginCharacterEquipmentCharacterEquipmentSummary) /* CharacterEquipmentSummary Returns a summary of the items equipped by a character. */

}




func ginCharacterEquipmentCharacterEquipmentSummary(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_CharacterEquipment.CharacterEquipmentSummaryFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_CharacterEquipment.CharacterEquipmentSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
