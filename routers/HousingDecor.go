package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/decor/index", ginDecor_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/decor/:decorId", ginDecor)
	app.Instance().RegisterRoute("GET", "/data/wow/search/decor", ginDecor_Search)
	app.Instance().RegisterRoute("GET", "/data/wow/fixture/index", ginFixture_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/fixture/:fixtureId", ginFixture)
	app.Instance().RegisterRoute("GET", "/data/wow/search/fixture", ginFixture_Search)
	app.Instance().RegisterRoute("GET", "/data/wow/fixture-hook/index", ginFixture_Hook_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/fixture-hook/:fixtureHookId", ginFixture_Hook)
	app.Instance().RegisterRoute("GET", "/data/wow/search/fixture-hook", ginFixture_Hook_Search)
	app.Instance().RegisterRoute("GET", "/data/wow/room/index", ginRoom_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/room/:roomId", ginRoom)
	app.Instance().RegisterRoute("GET", "/data/wow/search/room", ginRoom_Search)

}

func ginDecor_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginDecor(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginDecor_Search(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginFixture_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginFixture(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginFixture_Search(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginFixture_Hook_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginFixture_Hook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginFixture_Hook_Search(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginRoom_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginRoom(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginRoom_Search(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
