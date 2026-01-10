package routers

import (
	"net/http"

	"github.com/Thenecromance/Go_Blizzard_API/api/wow/ProfileService/Guild"
	"github.com/Thenecromance/Go_Blizzard_API/app"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/guild/:realmSlug/:nameSlug", ginwowGuild) /* Guild Returns a single guild by its name and realm. */

	app.Instance().RegisterRoute("GET", "/data/wow/guild/:realmSlug/:nameSlug/activity", ginwowGuildActivity) /* GuildActivity Returns a single guild's activity by name and realm. */

	app.Instance().RegisterRoute("GET", "/data/wow/guild/:realmSlug/:nameSlug/achievements", ginwowGuildAchievements) /* GuildAchievements Returns a single guild's achievements by name and realm. */

	app.Instance().RegisterRoute("GET", "/data/wow/guild/:realmSlug/:nameSlug/roster", ginwowGuildRoster) /* GuildRoster Returns a single guild's roster by its name and realm. */

}

func ginwowGuild(c *gin.Context) {
	// binding uri parameters
	var req wow_Guild.GuildFields
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

	resp, err := wow_Guild.Guild(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginwowGuildActivity(c *gin.Context) {
	// binding uri parameters
	var req wow_Guild.GuildActivityFields
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

	resp, err := wow_Guild.GuildActivity(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginwowGuildAchievements(c *gin.Context) {
	// binding uri parameters
	var req wow_Guild.GuildAchievementsFields
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

	resp, err := wow_Guild.GuildAchievements(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginwowGuildRoster(c *gin.Context) {
	// binding uri parameters
	var req wow_Guild.GuildRosterFields
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

	resp, err := wow_Guild.GuildRoster(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
