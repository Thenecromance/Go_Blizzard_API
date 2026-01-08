package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/DataService/Covenant"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/covenant/index", ginCovenantCovenantIndex) /* CovenantIndex Returns an index of covenants. */

	app.Instance().RegisterRoute("GET", "/data/wow/covenant/:covenantId", ginCovenantCovenant) /* Covenant Returns a covenant by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/media/covenant/:covenantId", ginCovenantCovenantMedia) /* CovenantMedia Returns media for a covenant by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/covenant/soulbind/index", ginCovenantSoulbindIndex) /* SoulbindIndex Returns an index of soulbinds. */

	app.Instance().RegisterRoute("GET", "/data/wow/covenant/soulbind/:soulbindId", ginCovenantSoulbind) /* Soulbind Returns a soulbind by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/covenant/conduit/index", ginCovenantConduitIndex) /* ConduitIndex Returns an index of conduits. */

	app.Instance().RegisterRoute("GET", "/data/wow/covenant/conduit/:conduitId", ginCovenantConduit) /* Conduit Returns a conduit by ID. */

}




func ginCovenantCovenantIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_Covenant.CovenantIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Covenant.CovenantIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCovenantCovenant(c *gin.Context) {
	// binding uri parameters
	var req wow_Covenant.CovenantFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Covenant.Covenant(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCovenantCovenantMedia(c *gin.Context) {
	// binding uri parameters
	var req wow_Covenant.CovenantMediaFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Covenant.CovenantMedia(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCovenantSoulbindIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_Covenant.SoulbindIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Covenant.SoulbindIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCovenantSoulbind(c *gin.Context) {
	// binding uri parameters
	var req wow_Covenant.SoulbindFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Covenant.Soulbind(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCovenantConduitIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_Covenant.ConduitIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Covenant.ConduitIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCovenantConduit(c *gin.Context) {
	// binding uri parameters
	var req wow_Covenant.ConduitFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Covenant.Conduit(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
