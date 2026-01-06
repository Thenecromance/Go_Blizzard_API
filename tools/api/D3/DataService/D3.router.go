package D3_routers

import "Unofficial_API/app"

func init() {

 app.Instance().RegisterRoute("GET", "/data/d3/season/",ginSeasonIndex)
 app.Instance().RegisterRoute("GET", "/data/d3/season/:id",ginSeason)
 app.Instance().RegisterRoute("GET", "/data/d3/season/:id/leaderboard/:leaderboard",ginSeasonLeaderboard)
 app.Instance().RegisterRoute("GET", "/data/d3/era/",ginEraIndex)
 app.Instance().RegisterRoute("GET", "/data/d3/era/:id",ginEra)
 app.Instance().RegisterRoute("GET", "/data/d3/era/:id/leaderboard/:leaderboard",ginEraLeaderboard)

}

func ginSeasonIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginSeason(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginSeasonLeaderboard(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginEraIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginEra(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginEraLeaderboard(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
