package SC_routers

import "Unofficial_API/app"

func init() {

 app.Instance().RegisterRoute("GET", "/data/sc2/league/:seasonId/:queueId/:teamType/:leagueId",gingetLeagueData)

}

func gingetLeagueData(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
