package routers

import (
	"net/http"

	"github.com/Thenecromance/Go_Blizzard_API/api/HeartStone/DataService/Cards"
	"github.com/Thenecromance/Go_Blizzard_API/app"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/hearthstone/cards", ginHeartStoneCardsearch) /* Cardsearch Returns an up-to-date list of all cards matching the search criteria. For more information about the search parameters, see the <a href='/documentation/hearthstone/guides'>Card Search Guide</a>. */

	//app.Instance().RegisterRoute("GET", "/hearthstone/cards", ginHeartStoneDetailedcardsearchexample) /* Detailedcardsearchexample Returns a specific card by using detailed search criteria. For more information about the search parameters, see the <a href='/documentation/hearthstone/guides/card-search'>Card Search Guide</a>. */
	//
	//app.Instance().RegisterRoute("GET", "/hearthstone/cards", ginHeartStoneBattlegroundscardsearch) /* Battlegroundscardsearch Returns an up-to-date list of all cards matching the search criteria for the specified game mode. For more information about the search parameters, see the <a href='/documentation/hearthstone/guides/card-search'>Card Search Guide</a>. */
	//
	//app.Instance().RegisterRoute("GET", "/hearthstone/cards", ginHeartStoneMercenariescardsearch) /* Mercenariescardsearch Returns an up-to-date list of all cards matching the search criteria for the specified game mode. For more information about the search parameters, see the <a href='/documentation/hearthstone/guides/card-search'>Card Search Guide</a>. */

	app.Instance().RegisterRoute("GET", "/hearthstone/cards/:idorslug", ginHeartStoneFetchonecard) /* Fetchonecard Returns the card with an ID or slug that matches the one you specify. For more information, see the <a href='/documentation/hearthstone/guides/card-search'>Card Search Guide</a>. */

}

func ginHeartStoneCardsearch(c *gin.Context) {
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

	// Capture all query parameters into a map if needed
	// Note: Gin's binding doesn't remove items, so this contains all query params
	queryParams := make(map[string]string)
	for k, v := range c.Request.URL.Query() {
		if len(v) > 0 {
			queryParams[k] = v[0]
		}
	}
	// If the downstream function supports an extra map, pass it here.
	// Assuming req might have a field for this or you manipulate req based on queryParams.
	// For now, I'm just showing how to extract them.
	// Example: req.ExtraParams = queryParams

	resp, err := HeartStone_Cards.Cardsearch(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginHeartStoneDetailedcardsearchexample(c *gin.Context) {
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

	// Capture all query parameters into a map if needed
	// Note: Gin's binding doesn't remove items, so this contains all query params
	queryParams := make(map[string]string)
	for k, v := range c.Request.URL.Query() {
		if len(v) > 0 {
			queryParams[k] = v[0]
		}
	}
	// If the downstream function supports an extra map, pass it here.
	// Assuming req might have a field for this or you manipulate req based on queryParams.
	// For now, I'm just showing how to extract them.
	// Example: req.ExtraParams = queryParams

	resp, err := HeartStone_Cards.Detailedcardsearchexample(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginHeartStoneBattlegroundscardsearch(c *gin.Context) {
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

	// Capture all query parameters into a map if needed
	// Note: Gin's binding doesn't remove items, so this contains all query params
	queryParams := make(map[string]string)
	for k, v := range c.Request.URL.Query() {
		if len(v) > 0 {
			queryParams[k] = v[0]
		}
	}
	// If the downstream function supports an extra map, pass it here.
	// Assuming req might have a field for this or you manipulate req based on queryParams.
	// For now, I'm just showing how to extract them.
	// Example: req.ExtraParams = queryParams

	resp, err := HeartStone_Cards.Battlegroundscardsearch(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginHeartStoneMercenariescardsearch(c *gin.Context) {
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

	// Capture all query parameters into a map if needed
	// Note: Gin's binding doesn't remove items, so this contains all query params
	queryParams := make(map[string]string)
	for k, v := range c.Request.URL.Query() {
		if len(v) > 0 {
			queryParams[k] = v[0]
		}
	}
	// If the downstream function supports an extra map, pass it here.
	// Assuming req might have a field for this or you manipulate req based on queryParams.
	// For now, I'm just showing how to extract them.
	// Example: req.ExtraParams = queryParams

	resp, err := HeartStone_Cards.Mercenariescardsearch(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ginHeartStoneFetchonecard(c *gin.Context) {
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

	// Capture all query parameters into a map if needed
	// Note: Gin's binding doesn't remove items, so this contains all query params
	queryParams := make(map[string]string)
	for k, v := range c.Request.URL.Query() {
		if len(v) > 0 {
			queryParams[k] = v[0]
		}
	}
	// If the downstream function supports an extra map, pass it here.
	// Assuming req might have a field for this or you manipulate req based on queryParams.
	// For now, I'm just showing how to extract them.
	// Example: req.ExtraParams = queryParams

	resp, err := HeartStone_Cards.Fetchonecard(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
