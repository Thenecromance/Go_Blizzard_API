package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wowClassic/DataService/GuildCrest"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/guild-crest/index", ginGuildCrestGuildCrestComponentsIndex) /* GuildCrestComponentsIndex Returns an index of guild crest media. */

	app.Instance().RegisterRoute("GET", "/data/wow/media/guild-crest/border/:borderId", ginGuildCrestGuildCrestBorderMedia) /* GuildCrestBorderMedia Returns media for a guild crest border by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/media/guild-crest/emblem/:emblemId", ginGuildCrestGuildCrestEmblemMedia) /* GuildCrestEmblemMedia Returns media for a guild crest emblem by ID. */

}




func ginGuildCrestGuildCrestComponentsIndex(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_GuildCrest.GuildCrestComponentsIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_GuildCrest.GuildCrestComponentsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginGuildCrestGuildCrestBorderMedia(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_GuildCrest.GuildCrestBorderMediaFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_GuildCrest.GuildCrestBorderMedia(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginGuildCrestGuildCrestEmblemMedia(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_GuildCrest.GuildCrestEmblemMediaFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_GuildCrest.GuildCrestEmblemMedia(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
