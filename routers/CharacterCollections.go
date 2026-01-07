package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/collections", ginCharacter_Collections_Index)
	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/collections/decor", ginCharacter_Decor_Collection_Summary)
	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/collections/heirlooms", ginCharacter_Heirlooms_Collection_Summary)
	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/collections/mounts", ginCharacter_Mounts_Collection_Summary)
	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/collections/pets", ginCharacter_Pets_Collection_Summary)
	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/collections/toys", ginCharacter_Toys_Collection_Summary)
	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/collections/transmogs", ginCharacter_Transmog_Collection_Summary)

}

func ginCharacter_Collections_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginCharacter_Decor_Collection_Summary(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginCharacter_Heirlooms_Collection_Summary(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginCharacter_Mounts_Collection_Summary(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginCharacter_Pets_Collection_Summary(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginCharacter_Toys_Collection_Summary(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginCharacter_Transmog_Collection_Summary(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
