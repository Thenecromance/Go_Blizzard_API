package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wowClassic/DataService/Creature"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/creature-family/index", ginCreatureCreatureFamiliesIndex) /* CreatureFamiliesIndex Returns an index of creature families. */

	app.Instance().RegisterRoute("GET", "/data/wow/creature-family/:creatureFamilyId", ginCreatureCreatureFamily) /* CreatureFamily Returns a creature family by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/creature-type/index", ginCreatureCreatureTypesIndex) /* CreatureTypesIndex Returns an index of creature types. */

	app.Instance().RegisterRoute("GET", "/data/wow/creature-type/:creatureTypeId", ginCreatureCreatureType) /* CreatureType Returns a creature type by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/creature/:creatureId", ginCreatureCreature) /* Creature Returns a creature by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/search/creature", ginCreatureCreatureSearch) /* CreatureSearch Performs a search of creatures. The fields below are provided for example. For more detail see the <a href="/documentation/world-of-warcraft/guides/search">Search Guide</a>. */

	app.Instance().RegisterRoute("GET", "/data/wow/media/creature-display/:creatureDisplayId", ginCreatureCreatureDisplayMedia) /* CreatureDisplayMedia Returns media for a creature display by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/media/creature-family/:creatureFamilyId", ginCreatureCreatureFamilyMedia) /* CreatureFamilyMedia Returns media for a creature family by ID. */

}




func ginCreatureCreatureFamiliesIndex(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_Creature.CreatureFamiliesIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_Creature.CreatureFamiliesIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCreatureCreatureFamily(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_Creature.CreatureFamilyFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_Creature.CreatureFamily(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCreatureCreatureTypesIndex(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_Creature.CreatureTypesIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_Creature.CreatureTypesIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCreatureCreatureType(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_Creature.CreatureTypeFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_Creature.CreatureType(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCreatureCreature(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_Creature.CreatureFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_Creature.Creature(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCreatureCreatureSearch(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_Creature.CreatureSearchFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_Creature.CreatureSearch(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCreatureCreatureDisplayMedia(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_Creature.CreatureDisplayMediaFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_Creature.CreatureDisplayMedia(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCreatureCreatureFamilyMedia(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_Creature.CreatureFamilyMediaFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_Creature.CreatureFamilyMedia(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
