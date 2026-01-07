package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wowClassic/DataService/PlayableClass"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/playable-class/index", ginPlayableClassesIndex) //PlayableClassesIndex Returns an index of playable classes.

	app.Instance().RegisterRoute("GET", "/data/wow/playable-class/:classId", ginPlayableClass) //PlayableClass Returns a playable class by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/media/playable-class/:playableClassId", ginPlayableClassMedia) //PlayableClassMedia Returns media for a playable class by ID.

}




func ginPlayableClassesIndex(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_PlayableClass.PlayableClassesIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_PlayableClass.PlayableClassesIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginPlayableClass(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_PlayableClass.PlayableClassFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_PlayableClass.PlayableClass(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginPlayableClassMedia(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_PlayableClass.PlayableClassMediaFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_PlayableClass.PlayableClassMedia(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
