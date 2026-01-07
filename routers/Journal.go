package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/journal-expansion/index", ginJournal_Expansions_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/journal-expansion/:journalExpansionId", ginJournal_Expansion)
	app.Instance().RegisterRoute("GET", "/data/wow/journal-encounter/index", ginJournal_Encounters_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/journal-encounter/:journalEncounterId", ginJournal_Encounter)
	app.Instance().RegisterRoute("GET", "/data/wow/search/journal-encounter", ginJournal_Encounter_Search)
	app.Instance().RegisterRoute("GET", "/data/wow/journal-instance/index", ginJournal_Instances_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/journal-instance/:journalInstanceId", ginJournal_Instance)
	app.Instance().RegisterRoute("GET", "/data/wow/media/journal-instance/:journalInstanceId", ginJournal_Instance_Media)

}

func ginJournal_Expansions_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginJournal_Expansion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginJournal_Encounters_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginJournal_Encounter(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginJournal_Encounter_Search(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginJournal_Instances_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginJournal_Instance(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginJournal_Instance_Media(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
