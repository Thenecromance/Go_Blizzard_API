package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/hearthstone/deck", ginGet_deck_by_code)

}

func ginGet_deck_by_code(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginGet_deck_by_card_list(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
