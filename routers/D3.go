package routers

import (
	"net/http"

	"Unofficial_API/api/D3/DataService/D3"
	"Unofficial_API/app"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/d3/season/", ginSeasonIndex) //SeasonIndex Returns an index of available seasons.

	app.Instance().RegisterRoute("GET", "/data/d3/season/:id", ginD3Season) //Season Returns a leaderboard list for the specified season.

	app.Instance().RegisterRoute("GET", "/data/d3/season/:id/leaderboard/:leaderboard", ginSeasonLeaderboard) //SeasonLeaderboard Returns a the specified leaderboard for the specified season.

	app.Instance().RegisterRoute("GET", "/data/d3/era/", ginEraIndex) //EraIndex Returns an index of available eras.

	app.Instance().RegisterRoute("GET", "/data/d3/era/:id", ginEra) //Era Returns a leaderboard list for a particular era.

	app.Instance().RegisterRoute("GET", "/data/d3/era/:id/leaderboard/:leaderboard", ginEraLeaderboard) //EraLeaderboard Returns the specified leaderboard for the specified era.

}

func ginSeasonIndex(c *gin.Context) {
	// binding uri parameters
	var req D3_D3.SeasonIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := D3_D3.SeasonIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginD3Season(c *gin.Context) {
	// binding uri parameters
	var req D3_D3.SeasonFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := D3_D3.Season(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginSeasonLeaderboard(c *gin.Context) {
	// binding uri parameters
	var req D3_D3.SeasonLeaderboardFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := D3_D3.SeasonLeaderboard(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginEraIndex(c *gin.Context) {
	// binding uri parameters
	var req D3_D3.EraIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := D3_D3.EraIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginEra(c *gin.Context) {
	// binding uri parameters
	var req D3_D3.EraFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := D3_D3.Era(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginEraLeaderboard(c *gin.Context) {
	// binding uri parameters
	var req D3_D3.EraLeaderboardFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := D3_D3.EraLeaderboard(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
