package routers

import (
	"net/http"

	"Unofficial_API/api/HeartStone/DataService/Decks"
	"Unofficial_API/app"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/hearthstone/deck", ginGetdeckbycode) //Getdeckbycode Finds a deck by deck code. For more information, see the <a href='/documentation/hearthstone/guides/decks'>Retrieving Decks Guide</a>.

	//app.Instance().RegisterRoute("GET", "/hearthstone/deck", ginGetdeckbycardlist) //Getdeckbycardlist Finds a deck by list of cards, including the hero. For more information, see the <a href='/documentation/hearthstone/guides/decks'>Retrieving Decks Guide</a>.

}

func ginGetdeckbycode(c *gin.Context) {
	// binding uri parameters
	var req HeartStone_Decks.GetdeckbycodeFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := HeartStone_Decks.Getdeckbycode(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginGetdeckbycardlist(c *gin.Context) {
	// binding uri parameters
	var req HeartStone_Decks.GetdeckbycardlistFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := HeartStone_Decks.Getdeckbycardlist(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
