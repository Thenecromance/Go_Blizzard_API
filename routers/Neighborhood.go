package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/neighborhood-map/index", ginNeighborhood_Map_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/neighborhood-map/:neighborhoodMapId", ginNeighborhood_Map)
	app.Instance().RegisterRoute("GET", "/data/wow/neighborhood-map/:neighborhoodMapId/neighborhood/:neighborhoodId", ginNeighborhood)

}

func ginNeighborhood_Map_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginNeighborhood_Map(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginNeighborhood(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
