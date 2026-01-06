package wowClassic_routers

import "Unofficial_API/app"

func init() {

 app.Instance().RegisterRoute("GET", "/profile/user/wow",ginAccountProfileSummary)
 app.Instance().RegisterRoute("GET", "/profile/user/wow/protected-character/:realmId-characterId",ginProtectedCharacterProfileSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/appearance",ginCharacterAppearanceSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/equipment",ginCharacterEquipmentSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/hunter-pets",ginCharacterHunterPetsSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/character-media",ginCharacterMediaSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName",ginCharacterProfileSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/status",ginCharacterProfileStatus)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/pvp-bracket/:pvpBracket",ginCharacterPvPBracketStatistics)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/pvp-summary",ginCharacterPvPSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/specializations",ginCharacterSpecializationsSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/statistics",ginCharacterStatisticsSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/achievements",ginCharacterAchievementsSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/achievements/statistics",ginCharacterAchievementStatistics)
 app.Instance().RegisterRoute("GET", "/data/wow/guild/:realmSlug/:nameSlug",ginGuild)
 app.Instance().RegisterRoute("GET", "/data/wow/guild/:realmSlug/:nameSlug/activity",ginGuildActivity)
 app.Instance().RegisterRoute("GET", "/data/wow/guild/:realmSlug/:nameSlug/achievements",ginGuildAchievements)
 app.Instance().RegisterRoute("GET", "/data/wow/guild/:realmSlug/:nameSlug/roster",ginGuildRoster)

}

func ginAccountProfileSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginProtectedCharacterProfileSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterAppearanceSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterEquipmentSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterHunterPetsSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterMediaSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterProfileSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterProfileStatus(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterPvPBracketStatistics(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterPvPSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterSpecializationsSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterStatisticsSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterAchievementsSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterAchievementStatistics(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginGuild(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginGuildActivity(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginGuildAchievements(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginGuildRoster(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
