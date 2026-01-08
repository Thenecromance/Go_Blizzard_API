package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wowClassic/ProfileService/Guild"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/guild/:realmSlug/:nameSlug", ginGuildGuild) /* Guild Returns a single guild by its name and realm. */

	app.Instance().RegisterRoute("GET", "/data/wow/guild/:realmSlug/:nameSlug/activity", ginGuildGuildActivity) /* GuildActivity Returns a single guild's activity by name and realm. */

	app.Instance().RegisterRoute("GET", "/data/wow/guild/:realmSlug/:nameSlug/achievements", ginGuildGuildAchievements) /* GuildAchievements Returns a single guild's achievements by name and realm. */

	app.Instance().RegisterRoute("GET", "/data/wow/guild/:realmSlug/:nameSlug/roster", ginGuildGuildRoster) /* GuildRoster Returns a single guild's roster by its name and realm. */

}




func ginGuildGuild(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_Guild.GuildFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_Guild.Guild(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginGuildGuildActivity(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_Guild.GuildActivityFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_Guild.GuildActivity(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginGuildGuildAchievements(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_Guild.GuildAchievementsFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_Guild.GuildAchievements(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginGuildGuildRoster(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_Guild.GuildRosterFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_Guild.GuildRoster(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
