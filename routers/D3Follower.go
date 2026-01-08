package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/D3/Community/D3Follower"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/d3/data/follower/:followerSlug", ginD3FollowerGetFollower) /* GetFollower Returns a single follower by slug. */

}




func ginD3FollowerGetFollower(c *gin.Context) {
	// binding uri parameters
	var req D3_D3Follower.GetFollowerFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := D3_D3Follower.GetFollower(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
