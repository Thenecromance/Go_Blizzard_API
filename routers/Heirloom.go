package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/DataService/Heirloom"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/heirloom/index", ginHeirloomHeirloomIndex) /* HeirloomIndex Returns an index of heirlooms. */

	app.Instance().RegisterRoute("GET", "/data/wow/heirloom/:heirloomId", ginHeirloomHeirloom) /* Heirloom Returns an heirloom by id. */

}




func ginHeirloomHeirloomIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_Heirloom.HeirloomIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Heirloom.HeirloomIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginHeirloomHeirloom(c *gin.Context) {
	// binding uri parameters
	var req wow_Heirloom.HeirloomFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Heirloom.Heirloom(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
