package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/DataService/Covenant"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/covenant/index", ginCovenantIndex) //CovenantIndex Returns an index of covenants.

	app.Instance().RegisterRoute("GET", "/data/wow/covenant/:covenantId", ginCovenant) //Covenant Returns a covenant by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/media/covenant/:covenantId", ginCovenantMedia) //CovenantMedia Returns media for a covenant by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/covenant/soulbind/index", ginSoulbindIndex) //SoulbindIndex Returns an index of soulbinds.

	app.Instance().RegisterRoute("GET", "/data/wow/covenant/soulbind/:soulbindId", ginSoulbind) //Soulbind Returns a soulbind by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/covenant/conduit/index", ginConduitIndex) //ConduitIndex Returns an index of conduits.

	app.Instance().RegisterRoute("GET", "/data/wow/covenant/conduit/:conduitId", ginConduit) //Conduit Returns a conduit by ID.

}




func ginCovenantIndex(c *gin.Context) {
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



func ginCovenant(c *gin.Context) {
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



func ginCovenantMedia(c *gin.Context) {
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



func ginSoulbindIndex(c *gin.Context) {
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



func ginSoulbind(c *gin.Context) {
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



func ginConduitIndex(c *gin.Context) {
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



func ginConduit(c *gin.Context) {
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
