package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wowClassic/DataService/PlayableRace"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/playable-race/index", ginPlayableRacesIndex) //PlayableRacesIndex Returns an index of playable races.

	app.Instance().RegisterRoute("GET", "/data/wow/playable-race/:playableRaceId", ginPlayableRace) //PlayableRace Returns a playable race by ID.

}




func ginPlayableRacesIndex(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_PlayableRace.PlayableRacesIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_PlayableRace.PlayableRacesIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginPlayableRace(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_PlayableRace.PlayableRaceFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_PlayableRace.PlayableRace(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
