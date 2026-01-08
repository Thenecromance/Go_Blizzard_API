package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/DataService/MythicKeystoneDungeon"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/index", ginMythicKeystoneDungeonMythicKeystoneIndex) /* MythicKeystoneIndex Returns an index of links to other documents related to Mythic Keystone dungeons. */

	app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/dungeon/index", ginMythicKeystoneDungeonMythicKeystoneDungeonsIndex) /* MythicKeystoneDungeonsIndex Returns an index of Mythic Keystone dungeons. */

	app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/dungeon/:dungeonId", ginMythicKeystoneDungeonMythicKeystoneDungeon) /* MythicKeystoneDungeon Returns a Mythic Keystone dungeon by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/period/index", ginMythicKeystoneDungeonMythicKeystonePeriodsIndex) /* MythicKeystonePeriodsIndex Returns an index of Mythic Keystone periods. */

	app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/period/:periodId", ginMythicKeystoneDungeonMythicKeystonePeriod) /* MythicKeystonePeriod Returns a Mythic Keystone period by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/season/index", ginMythicKeystoneDungeonMythicKeystoneSeasonsIndex) /* MythicKeystoneSeasonsIndex Returns an index of Mythic Keystone seasons. */

	app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/season/:seasonId", ginMythicKeystoneDungeonMythicKeystoneSeason) /* MythicKeystoneSeason Returns a Mythic Keystone season by ID. */

}




func ginMythicKeystoneDungeonMythicKeystoneIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_MythicKeystoneDungeon.MythicKeystoneIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_MythicKeystoneDungeon.MythicKeystoneIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginMythicKeystoneDungeonMythicKeystoneDungeonsIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_MythicKeystoneDungeon.MythicKeystoneDungeonsIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_MythicKeystoneDungeon.MythicKeystoneDungeonsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginMythicKeystoneDungeonMythicKeystoneDungeon(c *gin.Context) {
	// binding uri parameters
	var req wow_MythicKeystoneDungeon.MythicKeystoneDungeonFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_MythicKeystoneDungeon.MythicKeystoneDungeon(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginMythicKeystoneDungeonMythicKeystonePeriodsIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_MythicKeystoneDungeon.MythicKeystonePeriodsIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_MythicKeystoneDungeon.MythicKeystonePeriodsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginMythicKeystoneDungeonMythicKeystonePeriod(c *gin.Context) {
	// binding uri parameters
	var req wow_MythicKeystoneDungeon.MythicKeystonePeriodFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_MythicKeystoneDungeon.MythicKeystonePeriod(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginMythicKeystoneDungeonMythicKeystoneSeasonsIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_MythicKeystoneDungeon.MythicKeystoneSeasonsIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_MythicKeystoneDungeon.MythicKeystoneSeasonsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginMythicKeystoneDungeonMythicKeystoneSeason(c *gin.Context) {
	// binding uri parameters
	var req wow_MythicKeystoneDungeon.MythicKeystoneSeasonFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_MythicKeystoneDungeon.MythicKeystoneSeason(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
