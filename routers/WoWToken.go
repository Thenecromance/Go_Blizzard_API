package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wowClassic/DataService/WoWToken"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/token/index", ginWoWTokenWoWTokenIndexCN) /* WoWTokenIndexCN Returns the WoW Token index. */

}




func ginWoWTokenWoWTokenIndexCN(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_WoWToken.WoWTokenIndexCNFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_WoWToken.WoWTokenIndexCN(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
