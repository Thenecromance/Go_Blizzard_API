package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/DataService/PlayableSpecialization"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/playable-specialization/index", ginPlayableSpecializationsIndex) //PlayableSpecializationsIndex Returns an index of playable specializations.

	app.Instance().RegisterRoute("GET", "/data/wow/playable-specialization/:specId", ginPlayableSpecialization) //PlayableSpecialization Returns a playable specialization by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/media/playable-specialization/:specId", ginPlayableSpecializationMedia) //PlayableSpecializationMedia Returns media for a playable specialization by ID.

}




func ginPlayableSpecializationsIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_PlayableSpecialization.PlayableSpecializationsIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_PlayableSpecialization.PlayableSpecializationsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginPlayableSpecialization(c *gin.Context) {
	// binding uri parameters
	var req wow_PlayableSpecialization.PlayableSpecializationFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_PlayableSpecialization.PlayableSpecialization(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginPlayableSpecializationMedia(c *gin.Context) {
	// binding uri parameters
	var req wow_PlayableSpecialization.PlayableSpecializationMediaFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_PlayableSpecialization.PlayableSpecializationMedia(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
