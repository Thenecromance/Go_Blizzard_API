package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wowClassic/DataService/ConnectedRealm"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/index", ginConnectedRealmsIndex) //ConnectedRealmsIndex Returns an index of connected realms.

	app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/:connectedRealmId", ginConnectedRealm) //ConnectedRealm Returns a connected realm by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/search/connected-realm", ginConnectedRealmsSearch) //ConnectedRealmsSearch Performs a search of connected realms. The fields below are provided for example. For more detail see the <a href="/documentation/world-of-warcraft-classic/guides/search">Search Guide</a>.

}




func ginConnectedRealmsIndex(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_ConnectedRealm.ConnectedRealmsIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_ConnectedRealm.ConnectedRealmsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginConnectedRealm(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_ConnectedRealm.ConnectedRealmFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_ConnectedRealm.ConnectedRealm(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginConnectedRealmsSearch(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_ConnectedRealm.ConnectedRealmsSearchFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_ConnectedRealm.ConnectedRealmsSearch(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
