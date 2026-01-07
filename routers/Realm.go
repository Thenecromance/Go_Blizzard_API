package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wowClassic/DataService/Realm"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/realm/index", ginRealmsIndex) //RealmsIndex Returns an index of realms.

	app.Instance().RegisterRoute("GET", "/data/wow/realm/:realmSlug", ginRealm) //Realm Returns a single realm by slug or ID.

	app.Instance().RegisterRoute("GET", "/data/wow/search/realm", ginRealmSearch) //RealmSearch Performs a search of realms. The fields below are examples only. For more detail see the <a href="/documentation/world-of-warcraft-classic/guides/search">Search Guide</a>.

}




func ginRealmsIndex(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_Realm.RealmsIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_Realm.RealmsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginRealm(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_Realm.RealmFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_Realm.Realm(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginRealmSearch(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_Realm.RealmSearchFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_Realm.RealmSearch(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
