package routers

import (
	"net/http"

	"github.com/Thenecromance/Go_Blizzard_API/api/wow/DataService/Journal"
	"github.com/Thenecromance/Go_Blizzard_API/app"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/journal-expansion/index", ginwowJournalExpansionsIndex) /* JournalExpansionsIndex Returns an index of journal expansions. */

	app.Instance().RegisterRoute("GET", "/data/wow/journal-expansion/:journalExpansionId", ginwowJournalExpansion) /* JournalExpansion Returns a journal expansion by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/journal-encounter/index", ginwowJournalEncountersIndex) /* JournalEncountersIndex Returns an index of journal encounters. */

	app.Instance().RegisterRoute("GET", "/data/wow/journal-encounter/:journalEncounterId", ginwowJournalEncounter) /* JournalEncounter Returns a journal encounter by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/search/journal-encounter", ginwowJournalEncounterSearch) /* JournalEncounterSearch Performs a search of journal encounters. The fields below are provided for example. For more detail see the <a href="/documentation/world-of-warcraft/guides/search">Search Guide</a>. */

	app.Instance().RegisterRoute("GET", "/data/wow/journal-instance/index", ginwowJournalInstancesIndex) /* JournalInstancesIndex Returns an index of journal instances. */

	app.Instance().RegisterRoute("GET", "/data/wow/journal-instance/:journalInstanceId", ginwowJournalInstance) /* JournalInstance Returns a journal instance. */

	app.Instance().RegisterRoute("GET", "/data/wow/media/journal-instance/:journalInstanceId", ginwowJournalInstanceMedia) /* JournalInstanceMedia Returns media for a journal instance by ID. */

}

func ginwowJournalExpansionsIndex(c *gin.Context) {
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

	// Capture all query parameters into a map if needed
	// Note: Gin's binding doesn't remove items, so this contains all query params
	queryParams := make(map[string]string)
	for k, v := range c.Request.URL.Query() {
		if len(v) > 0 {
			queryParams[k] = v[0]
		}
	}
	// If the downstream function supports an extra map, pass it here.
	// Assuming req might have a field for this or you manipulate req based on queryParams.
	// For now, I'm just showing how to extract them.
	// Example: req.ExtraParams = queryParams

	resp, err := wow_Journal.JournalExpansionsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginwowJournalExpansion(c *gin.Context) {
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

	// Capture all query parameters into a map if needed
	// Note: Gin's binding doesn't remove items, so this contains all query params
	queryParams := make(map[string]string)
	for k, v := range c.Request.URL.Query() {
		if len(v) > 0 {
			queryParams[k] = v[0]
		}
	}
	// If the downstream function supports an extra map, pass it here.
	// Assuming req might have a field for this or you manipulate req based on queryParams.
	// For now, I'm just showing how to extract them.
	// Example: req.ExtraParams = queryParams

	resp, err := wow_Journal.JournalExpansion(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginwowJournalEncountersIndex(c *gin.Context) {
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

	// Capture all query parameters into a map if needed
	// Note: Gin's binding doesn't remove items, so this contains all query params
	queryParams := make(map[string]string)
	for k, v := range c.Request.URL.Query() {
		if len(v) > 0 {
			queryParams[k] = v[0]
		}
	}
	// If the downstream function supports an extra map, pass it here.
	// Assuming req might have a field for this or you manipulate req based on queryParams.
	// For now, I'm just showing how to extract them.
	// Example: req.ExtraParams = queryParams

	resp, err := wow_Journal.JournalEncountersIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginwowJournalEncounter(c *gin.Context) {
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

	// Capture all query parameters into a map if needed
	// Note: Gin's binding doesn't remove items, so this contains all query params
	queryParams := make(map[string]string)
	for k, v := range c.Request.URL.Query() {
		if len(v) > 0 {
			queryParams[k] = v[0]
		}
	}
	// If the downstream function supports an extra map, pass it here.
	// Assuming req might have a field for this or you manipulate req based on queryParams.
	// For now, I'm just showing how to extract them.
	// Example: req.ExtraParams = queryParams

	resp, err := wow_Journal.JournalEncounter(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginwowJournalEncounterSearch(c *gin.Context) {
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

	// Capture all query parameters into a map if needed
	// Note: Gin's binding doesn't remove items, so this contains all query params
	queryParams := make(map[string]string)
	for k, v := range c.Request.URL.Query() {
		if len(v) > 0 {
			queryParams[k] = v[0]
		}
	}
	// If the downstream function supports an extra map, pass it here.
	// Assuming req might have a field for this or you manipulate req based on queryParams.
	// For now, I'm just showing how to extract them.
	// Example: req.ExtraParams = queryParams

	resp, err := wow_Journal.JournalEncounterSearch(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginwowJournalInstancesIndex(c *gin.Context) {
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

	// Capture all query parameters into a map if needed
	// Note: Gin's binding doesn't remove items, so this contains all query params
	queryParams := make(map[string]string)
	for k, v := range c.Request.URL.Query() {
		if len(v) > 0 {
			queryParams[k] = v[0]
		}
	}
	// If the downstream function supports an extra map, pass it here.
	// Assuming req might have a field for this or you manipulate req based on queryParams.
	// For now, I'm just showing how to extract them.
	// Example: req.ExtraParams = queryParams

	resp, err := wow_Journal.JournalInstancesIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginwowJournalInstance(c *gin.Context) {
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

	// Capture all query parameters into a map if needed
	// Note: Gin's binding doesn't remove items, so this contains all query params
	queryParams := make(map[string]string)
	for k, v := range c.Request.URL.Query() {
		if len(v) > 0 {
			queryParams[k] = v[0]
		}
	}
	// If the downstream function supports an extra map, pass it here.
	// Assuming req might have a field for this or you manipulate req based on queryParams.
	// For now, I'm just showing how to extract them.
	// Example: req.ExtraParams = queryParams

	resp, err := wow_Journal.JournalInstance(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginwowJournalInstanceMedia(c *gin.Context) {
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

	// Capture all query parameters into a map if needed
	// Note: Gin's binding doesn't remove items, so this contains all query params
	queryParams := make(map[string]string)
	for k, v := range c.Request.URL.Query() {
		if len(v) > 0 {
			queryParams[k] = v[0]
		}
	}
	// If the downstream function supports an extra map, pass it here.
	// Assuming req might have a field for this or you manipulate req based on queryParams.
	// For now, I'm just showing how to extract them.
	// Example: req.ExtraParams = queryParams

	resp, err := wow_Journal.JournalInstanceMedia(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
