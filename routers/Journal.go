package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/DataService/Journal"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/journal-expansion/index", ginJournalJournalExpansionsIndex) /* JournalExpansionsIndex Returns an index of journal expansions. */

	app.Instance().RegisterRoute("GET", "/data/wow/journal-expansion/:journalExpansionId", ginJournalJournalExpansion) /* JournalExpansion Returns a journal expansion by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/journal-encounter/index", ginJournalJournalEncountersIndex) /* JournalEncountersIndex Returns an index of journal encounters. */

	app.Instance().RegisterRoute("GET", "/data/wow/journal-encounter/:journalEncounterId", ginJournalJournalEncounter) /* JournalEncounter Returns a journal encounter by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/search/journal-encounter", ginJournalJournalEncounterSearch) /* JournalEncounterSearch Performs a search of journal encounters. The fields below are provided for example. For more detail see the <a href="/documentation/world-of-warcraft/guides/search">Search Guide</a>. */

	app.Instance().RegisterRoute("GET", "/data/wow/journal-instance/index", ginJournalJournalInstancesIndex) /* JournalInstancesIndex Returns an index of journal instances. */

	app.Instance().RegisterRoute("GET", "/data/wow/journal-instance/:journalInstanceId", ginJournalJournalInstance) /* JournalInstance Returns a journal instance. */

	app.Instance().RegisterRoute("GET", "/data/wow/media/journal-instance/:journalInstanceId", ginJournalJournalInstanceMedia) /* JournalInstanceMedia Returns media for a journal instance by ID. */

}




func ginJournalJournalExpansionsIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_Journal.JournalExpansionsIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Journal.JournalExpansionsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginJournalJournalExpansion(c *gin.Context) {
	// binding uri parameters
	var req wow_Journal.JournalExpansionFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Journal.JournalExpansion(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginJournalJournalEncountersIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_Journal.JournalEncountersIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Journal.JournalEncountersIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginJournalJournalEncounter(c *gin.Context) {
	// binding uri parameters
	var req wow_Journal.JournalEncounterFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Journal.JournalEncounter(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginJournalJournalEncounterSearch(c *gin.Context) {
	// binding uri parameters
	var req wow_Journal.JournalEncounterSearchFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Journal.JournalEncounterSearch(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginJournalJournalInstancesIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_Journal.JournalInstancesIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Journal.JournalInstancesIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginJournalJournalInstance(c *gin.Context) {
	// binding uri parameters
	var req wow_Journal.JournalInstanceFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Journal.JournalInstance(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginJournalJournalInstanceMedia(c *gin.Context) {
	// binding uri parameters
	var req wow_Journal.JournalInstanceMediaFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Journal.JournalInstanceMedia(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
