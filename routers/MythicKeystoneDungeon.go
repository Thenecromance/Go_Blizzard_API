package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/DataService/MythicKeystoneDungeon"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/index", ginMythicKeystoneIndex) //MythicKeystoneIndex Returns an index of links to other documents related to Mythic Keystone dungeons.

	app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/dungeon/index", ginMythicKeystoneDungeonsIndex) //MythicKeystoneDungeonsIndex Returns an index of Mythic Keystone dungeons.

	app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/dungeon/:dungeonId", ginMythicKeystoneDungeon) //MythicKeystoneDungeon Returns a Mythic Keystone dungeon by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/period/index", ginMythicKeystonePeriodsIndex) //MythicKeystonePeriodsIndex Returns an index of Mythic Keystone periods.

	app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/period/:periodId", ginMythicKeystonePeriod) //MythicKeystonePeriod Returns a Mythic Keystone period by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/season/index", ginMythicKeystoneSeasonsIndex) //MythicKeystoneSeasonsIndex Returns an index of Mythic Keystone seasons.

	app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/season/:seasonId", ginMythicKeystoneSeason) //MythicKeystoneSeason Returns a Mythic Keystone season by ID.

}




func ginMythicKeystoneIndex(c *gin.Context) {
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



func ginMythicKeystoneDungeonsIndex(c *gin.Context) {
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



func ginMythicKeystoneDungeon(c *gin.Context) {
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



func ginMythicKeystonePeriodsIndex(c *gin.Context) {
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



func ginMythicKeystonePeriod(c *gin.Context) {
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



func ginMythicKeystoneSeasonsIndex(c *gin.Context) {
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



func ginMythicKeystoneSeason(c *gin.Context) {
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
