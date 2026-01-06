package wow_routers

import "Unofficial_API/app"

func init() {

 app.Instance().RegisterRoute("GET", "/profile/user/wow",ginAccountProfileSummary)
 app.Instance().RegisterRoute("GET", "/profile/user/wow/protected-character/:realmId-characterId",ginProtectedCharacterProfileSummary)
 app.Instance().RegisterRoute("GET", "/profile/user/wow/collections",ginAccountCollectionsIndex)
 app.Instance().RegisterRoute("GET", "/profile/user/wow/collections/decor",ginAccountDecorCollectionSummary)
 app.Instance().RegisterRoute("GET", "/profile/user/wow/collections/heirlooms",ginAccountHeirloomsCollectionSummary)
 app.Instance().RegisterRoute("GET", "/profile/user/wow/collections/mounts",ginAccountMountsCollectionSummary)
 app.Instance().RegisterRoute("GET", "/profile/user/wow/collections/pets",ginAccountPetsCollectionSummary)
 app.Instance().RegisterRoute("GET", "/profile/user/wow/collections/toys",ginAccountToysCollectionSummary)
 app.Instance().RegisterRoute("GET", "/profile/user/wow/collections/transmogs",ginAccountTransmogCollectionSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/achievements",ginCharacterAchievementsSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/achievements/statistics",ginCharacterAchievementStatistics)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/appearance",ginCharacterAppearanceSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/collections",ginCharacterCollectionsIndex)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/collections/decor",ginCharacterDecorCollectionSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/collections/heirlooms",ginCharacterHeirloomsCollectionSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/collections/mounts",ginCharacterMountsCollectionSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/collections/pets",ginCharacterPetsCollectionSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/collections/toys",ginCharacterToysCollectionSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/collections/transmogs",ginCharacterTransmogCollectionSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/encounters",ginCharacterEncountersSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/encounters/dungeons",ginCharacterDungeons)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/encounters/raids",ginCharacterRaids)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/equipment",ginCharacterEquipmentSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/house/house-{houseNumber}",ginCharacterHouseSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/hunter-pets",ginCharacterHunterPetsSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/character-media",ginCharacterMediaSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/mythic-keystone-profile",ginCharacterMythicKeystoneProfileIndex)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/mythic-keystone-profile/season/:seasonId",ginCharacterMythicKeystoneSeasonDetails)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/professions",ginCharacterProfessionsSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName",ginCharacterProfileSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/status",ginCharacterProfileStatus)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/pvp-bracket/:pvpBracket",ginCharacterPvPBracketStatistics)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/pvp-summary",ginCharacterPvPSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/quests",ginCharacterQuests)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/quests/completed",ginCharacterCompletedQuests)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/reputations",ginCharacterReputationsSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/soulbinds",ginCharacterSoulbinds)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/specializations",ginCharacterSpecializationsSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/statistics",ginCharacterStatisticsSummary)
 app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/titles",ginCharacterTitlesSummary)
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
func ginAccountCollectionsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginAccountDecorCollectionSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginAccountHeirloomsCollectionSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginAccountMountsCollectionSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginAccountPetsCollectionSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginAccountToysCollectionSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginAccountTransmogCollectionSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterAchievementsSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterAchievementStatistics(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterAppearanceSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterCollectionsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterDecorCollectionSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterHeirloomsCollectionSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterMountsCollectionSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterPetsCollectionSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterToysCollectionSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterTransmogCollectionSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterEncountersSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterDungeons(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterRaids(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterEquipmentSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterHouseSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterHunterPetsSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterMediaSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterMythicKeystoneProfileIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterMythicKeystoneSeasonDetails(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterProfessionsSummary(c *gin.Context){
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
func ginCharacterQuests(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterCompletedQuests(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterReputationsSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterSoulbinds(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterSpecializationsSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterStatisticsSummary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCharacterTitlesSummary(c *gin.Context){
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
