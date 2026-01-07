package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/hearthstone/cards", ginCard_search)
	app.Instance().RegisterRoute("GET", "/hearthstone/cards/:idorslug", ginFetch_one_card)

}

func ginCard_search(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginDetailed_card_search_example(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginBattlegrounds_card_search(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginMercenaries_card_search(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginFetch_one_card(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
