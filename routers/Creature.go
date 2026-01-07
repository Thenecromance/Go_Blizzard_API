package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wowClassic/DataService/Creature"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/creature-family/index", ginCreatureFamiliesIndex) //CreatureFamiliesIndex Returns an index of creature families.

	app.Instance().RegisterRoute("GET", "/data/wow/creature-family/:creatureFamilyId", ginCreatureFamily) //CreatureFamily Returns a creature family by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/creature-type/index", ginCreatureTypesIndex) //CreatureTypesIndex Returns an index of creature types.

	app.Instance().RegisterRoute("GET", "/data/wow/creature-type/:creatureTypeId", ginCreatureType) //CreatureType Returns a creature type by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/creature/:creatureId", ginCreature) //Creature Returns a creature by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/search/creature", ginCreatureSearch) //CreatureSearch Performs a search of creatures. The fields below are provided for example. For more detail see the <a href="/documentation/world-of-warcraft/guides/search">Search Guide</a>.

	app.Instance().RegisterRoute("GET", "/data/wow/media/creature-display/:creatureDisplayId", ginCreatureDisplayMedia) //CreatureDisplayMedia Returns media for a creature display by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/media/creature-family/:creatureFamilyId", ginCreatureFamilyMedia) //CreatureFamilyMedia Returns media for a creature family by ID.

}




func ginCreatureFamiliesIndex(c *gin.Context) {
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



func ginCreatureFamily(c *gin.Context) {
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



func ginCreatureTypesIndex(c *gin.Context) {
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



func ginCreatureType(c *gin.Context) {
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



func ginCreature(c *gin.Context) {
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



func ginCreatureSearch(c *gin.Context) {
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



func ginCreatureDisplayMedia(c *gin.Context) {
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



func ginCreatureFamilyMedia(c *gin.Context) {
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
