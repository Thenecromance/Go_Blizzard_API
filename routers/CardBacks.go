package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/HeartStone/DataService/CardBacks"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/hearthstone/cardbacks", ginCardBackSearch) //CardBackSearch Returns an up-to-date list of all card backs matching the search criteria. For more information about the search parameters, see the <a href='/documentation/hearthstone/guides/card-backs'>Card Backs Guide</a>.

	app.Instance().RegisterRoute("GET", "/hearthstone/cardbacks/:idorslug", ginFetchonecardback) //Fetchonecardback Returns a specific card back by using card back ID or slug.

}




func ginCardBackSearch(c *gin.Context) {
	// binding uri parameters
	var req HeartStone_CardBacks.CardBackSearchFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := HeartStone_CardBacks.CardBackSearch(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginFetchonecardback(c *gin.Context) {
	// binding uri parameters
	var req HeartStone_CardBacks.FetchonecardbackFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := HeartStone_CardBacks.Fetchonecardback(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
