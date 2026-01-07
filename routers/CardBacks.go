package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/hearthstone/cardbacks", ginCard_Back_Search)
	app.Instance().RegisterRoute("GET", "/hearthstone/cardbacks/:idorslug", ginFetch_one_card_back)

}

func ginCard_Back_Search(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginFetch_one_card_back(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
