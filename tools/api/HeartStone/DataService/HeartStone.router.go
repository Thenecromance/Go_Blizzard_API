package HeartStone_routers

import "Unofficial_API/app"

func init() {

 app.Instance().RegisterRoute("GET", "/hearthstone/cards",ginCardsearch)
 app.Instance().RegisterRoute("GET", "/hearthstone/cards",ginDetailedcardsearchexample)
 app.Instance().RegisterRoute("GET", "/hearthstone/cards",ginBattlegroundscardsearch)
 app.Instance().RegisterRoute("GET", "/hearthstone/cards",ginMercenariescardsearch)
 app.Instance().RegisterRoute("GET", "/hearthstone/cards/:idorslug",ginFetchonecard)
 app.Instance().RegisterRoute("GET", "/hearthstone/cardbacks",ginCardBackSearch)
 app.Instance().RegisterRoute("GET", "/hearthstone/cardbacks/:idorslug",ginFetchonecardback)
 app.Instance().RegisterRoute("GET", "/hearthstone/deck",ginGetdeckbycode)
 app.Instance().RegisterRoute("GET", "/hearthstone/deck",ginGetdeckbycardlist)
 app.Instance().RegisterRoute("GET", "/hearthstone/metadata",ginAllmetadata)
 app.Instance().RegisterRoute("GET", "/hearthstone/metadata/:type",ginSpecificmetadata)

}

func ginCardsearch(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginDetailedcardsearchexample(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginBattlegroundscardsearch(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginMercenariescardsearch(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginFetchonecard(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCardBackSearch(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginFetchonecardback(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginGetdeckbycode(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginGetdeckbycardlist(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginAllmetadata(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginSpecificmetadata(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
