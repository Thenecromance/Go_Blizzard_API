package routers

import (
	"net/http"

	"Unofficial_API/api/HeartStone/DataService/Cards"
	"Unofficial_API/app"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/hearthstone/cards", ginCardsearch) //Cardsearch Returns an up-to-date list of all cards matching the search criteria. For more information about the search parameters, see the <a href='/documentation/hearthstone/guides'>Card Search Guide</a>.

	/*app.Instance().RegisterRoute("GET", "/hearthstone/cards", ginDetailedcardsearchexample) //Detailedcardsearchexample Returns a specific card by using detailed search criteria. For more information about the search parameters, see the <a href='/documentation/hearthstone/guides/card-search'>Card Search Guide</a>.

	app.Instance().RegisterRoute("GET", "/hearthstone/cards", ginBattlegroundscardsearch) //Battlegroundscardsearch Returns an up-to-date list of all cards matching the search criteria for the specified game mode. For more information about the search parameters, see the <a href='/documentation/hearthstone/guides/card-search'>Card Search Guide</a>.

	app.Instance().RegisterRoute("GET", "/hearthstone/cards", ginMercenariescardsearch) //Mercenariescardsearch Returns an up-to-date list of all cards matching the search criteria for the specified game mode. For more information about the search parameters, see the <a href='/documentation/hearthstone/guides/card-search'>Card Search Guide</a>.
	*/
	app.Instance().RegisterRoute("GET", "/hearthstone/cards/:idorslug", ginFetchonecard) //Fetchonecard Returns the card with an ID or slug that matches the one you specify. For more information, see the <a href='/documentation/hearthstone/guides/card-search'>Card Search Guide</a>.

}

func ginCardsearch(c *gin.Context) {
	// binding uri parameters
	var req HeartStone_Cards.CardsearchFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := HeartStone_Cards.Cardsearch(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginDetailedcardsearchexample(c *gin.Context) {
	// binding uri parameters
	var req HeartStone_Cards.DetailedcardsearchexampleFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := HeartStone_Cards.Detailedcardsearchexample(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginBattlegroundscardsearch(c *gin.Context) {
	// binding uri parameters
	var req HeartStone_Cards.BattlegroundscardsearchFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := HeartStone_Cards.Battlegroundscardsearch(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginMercenariescardsearch(c *gin.Context) {
	// binding uri parameters
	var req HeartStone_Cards.MercenariescardsearchFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := HeartStone_Cards.Mercenariescardsearch(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginFetchonecard(c *gin.Context) {
	// binding uri parameters
	var req HeartStone_Cards.FetchonecardFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := HeartStone_Cards.Fetchonecard(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
