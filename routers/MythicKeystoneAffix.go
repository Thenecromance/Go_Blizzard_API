package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/keystone-affix/index", ginMythic_Keystone_Affixes_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/keystone-affix/:keystoneAffixId", ginMythic_Keystone_Affix)
	app.Instance().RegisterRoute("GET", "/data/wow/media/keystone-affix/:keystoneAffixId", ginMythic_Keystone_Affix_Media)

}

func ginMythic_Keystone_Affixes_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginMythic_Keystone_Affix(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginMythic_Keystone_Affix_Media(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
