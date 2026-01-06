package SC_routers

import "Unofficial_API/app"

func init() {

 app.Instance().RegisterRoute("GET", "/sc2/static/profile/:regionId",ginStatic)
 app.Instance().RegisterRoute("GET", "/sc2/metadata/profile/:regionId/:realmId/:profileId",ginMetadata)
 app.Instance().RegisterRoute("GET", "/sc2/profile/:regionID/:realmID/:profileId",ginProfile)
 app.Instance().RegisterRoute("GET", "/sc2/profile/:regionId/:realmId/:profileId/ladder/summary",ginLadderSummary)
 app.Instance().RegisterRoute("GET", "/sc2/profile/:regionId/:realmId/:profileId/ladder/:ladderId",ginLadder)
 app.Instance().RegisterRoute("GET", "/sc2/ladder/grandmaster/:regionId",ginGrandmasterLeaderboard)
 app.Instance().RegisterRoute("GET", "/sc2/ladder/season/:regionId",ginSeason)
 app.Instance().RegisterRoute("GET", "/sc2/player/:accountId",ginPlayer)
 app.Instance().RegisterRoute("GET", "/sc2/legacy/profile/:regionId/:realmId/:profileId",ginProfile)
 app.Instance().RegisterRoute("GET", "/sc2/legacy/profile/:regionId/:realmId/:profileId/ladders",ginLadders)
 app.Instance().RegisterRoute("GET", "/sc2/legacy/profile/:regionId/:realmId/:profileId/matches",ginMatchHistory)
 app.Instance().RegisterRoute("GET", "/sc2/legacy/ladder/:regionId/:ladderId",ginLadder)
 app.Instance().RegisterRoute("GET", "/sc2/legacy/data/achievements/:regionId",ginAchievements)
 app.Instance().RegisterRoute("GET", "/sc2/legacy/data/rewards/:regionId",ginRewards)

}

func ginStatic(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginMetadata(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginProfile(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginLadderSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginLadder(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginGrandmasterLeaderboard(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginSeason(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPlayer(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginProfile(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginLadders(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginMatchHistory(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginLadder(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginAchievements(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginRewards(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
