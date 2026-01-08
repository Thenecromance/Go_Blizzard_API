package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/D3/Community/D3Act"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/d3/data/act", ginD3ActGetActIndex) /* GetActIndex Returns an index of acts. */

	app.Instance().RegisterRoute("GET", "/d3/data/act/:actId", ginD3ActGetAct) /* GetAct Returns a single act by ID. */

}




func ginD3ActGetActIndex(c *gin.Context) {
	// binding uri parameters
	var req D3_D3Act.GetActIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := D3_D3Act.GetActIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginD3ActGetAct(c *gin.Context) {
	// binding uri parameters
	var req D3_D3Act.GetActFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := D3_D3Act.GetAct(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
