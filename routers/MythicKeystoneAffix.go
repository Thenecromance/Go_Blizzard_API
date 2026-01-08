package routers

import (
	"net/http"

	"Unofficial_API/api/wow/DataService/MythicKeystoneAffix"
	"Unofficial_API/app"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/keystone-affix/index", ginMythicKeystoneAffixesIndex) //MythicKeystoneAffixesIndex Returns an index of mythic keystone affixes.

	app.Instance().RegisterRoute("GET", "/data/wow/keystone-affix/:keystoneAffixId", ginMythicKeystoneAffix) //MythicKeystoneAffix Returns a mythic keystone affix by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/media/keystone-affix/:keystoneAffixId", ginMythicKeystoneAffixMedia) //MythicKeystoneAffixMedia Returns media for a mythic keystone affix by ID.

}

func ginMythicKeystoneAffixesIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_MythicKeystoneAffix.MythicKeystoneAffixesIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_MythicKeystoneAffix.MythicKeystoneAffixesIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.(*wow_MythicKeystoneAffix.BNetMythicKeystoneAffixesIndex))
}

func ginMythicKeystoneAffix(c *gin.Context) {
	// binding uri parameters
	var req wow_MythicKeystoneAffix.MythicKeystoneAffixFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_MythicKeystoneAffix.MythicKeystoneAffix(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginMythicKeystoneAffixMedia(c *gin.Context) {
	// binding uri parameters
	var req wow_MythicKeystoneAffix.MythicKeystoneAffixMediaFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_MythicKeystoneAffix.MythicKeystoneAffixMedia(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
