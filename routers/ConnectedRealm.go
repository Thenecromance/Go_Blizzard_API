package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wowClassic/DataService/ConnectedRealm"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/index", ginConnectedRealmConnectedRealmsIndex) /* ConnectedRealmsIndex Returns an index of connected realms. */

	app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/:connectedRealmId", ginConnectedRealmConnectedRealm) /* ConnectedRealm Returns a connected realm by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/search/connected-realm", ginConnectedRealmConnectedRealmsSearch) /* ConnectedRealmsSearch Performs a search of connected realms. The fields below are provided for example. For more detail see the <a href="/documentation/world-of-warcraft-classic/guides/search">Search Guide</a>. */

}




func ginConnectedRealmConnectedRealmsIndex(c *gin.Context) {
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



func ginConnectedRealmConnectedRealm(c *gin.Context) {
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



func ginConnectedRealmConnectedRealmsSearch(c *gin.Context) {
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
