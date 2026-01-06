package wowClassic_routers

import "Unofficial_API/app"

func init() {

 app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/:connectedRealmId/auctions/index",ginAuctionHouseIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/:connectedRealmId/auctions/:auctionHouseId",ginAuctions)
 app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/index",ginConnectedRealmsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/:connectedRealmId",ginConnectedRealm)
 app.Instance().RegisterRoute("GET", "/data/wow/search/connected-realm",ginConnectedRealmsSearch)
 app.Instance().RegisterRoute("GET", "/data/wow/creature-family/index",ginCreatureFamiliesIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/creature-family/:creatureFamilyId",ginCreatureFamily)
 app.Instance().RegisterRoute("GET", "/data/wow/creature-type/index",ginCreatureTypesIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/creature-type/:creatureTypeId",ginCreatureType)
 app.Instance().RegisterRoute("GET", "/data/wow/creature/:creatureId",ginCreature)
 app.Instance().RegisterRoute("GET", "/data/wow/search/creature",ginCreatureSearch)
 app.Instance().RegisterRoute("GET", "/data/wow/media/creature-display/:creatureDisplayId",ginCreatureDisplayMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/media/creature-family/:creatureFamilyId",ginCreatureFamilyMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/guild-crest/index",ginGuildCrestComponentsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/media/guild-crest/border/:borderId",ginGuildCrestBorderMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/media/guild-crest/emblem/:emblemId",ginGuildCrestEmblemMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/item-class/index",ginItemClassesIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/item-class/:itemClassId",ginItemClass)
 app.Instance().RegisterRoute("GET", "/data/wow/item-class/:itemClassId/item-subclass/:itemSubclassId",ginItemSubclass)
 app.Instance().RegisterRoute("GET", "/data/wow/item/:itemId",ginItem)
 app.Instance().RegisterRoute("GET", "/data/wow/media/item/:itemId",ginItemMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/search/item",ginItemSearch)
 app.Instance().RegisterRoute("GET", "/data/wow/search/media",ginMediaSearch)
 app.Instance().RegisterRoute("GET", "/data/wow/playable-class/index",ginPlayableClassesIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/playable-class/:classId",ginPlayableClass)
 app.Instance().RegisterRoute("GET", "/data/wow/media/playable-class/:playableClassId",ginPlayableClassMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/playable-race/index",ginPlayableRacesIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/playable-race/:playableRaceId",ginPlayableRace)
 app.Instance().RegisterRoute("GET", "/data/wow/power-type/index",ginPowerTypesIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/power-type/:powerTypeId",ginPowerType)
 app.Instance().RegisterRoute("GET", "/data/wow/pvp-season/index",ginPvPSeasonsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/pvp-season/:pvpSeasonId",ginPvPSeason)
 app.Instance().RegisterRoute("GET", "/data/wow/pvp-region/index",ginPvPRegionIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/pvp-region/:pvpRegionId/pvp-season/index",ginPvPRegionalSeasonIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/pvp-region/:pvpRegionId/pvp-season/:pvpSeasonId",ginPvPRegionalSeason)
 app.Instance().RegisterRoute("GET", "/data/wow/pvp-region/:pvpRegionId/pvp-season/:pvpSeasonId/pvp-leaderboard/index",ginPvPLeaderboardsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/pvp-region/:pvpRegionId/pvp-season/:pvpSeasonId/pvp-leaderboard/:pvpBracket",ginPvPLeaderboard)
 app.Instance().RegisterRoute("GET", "/data/wow/pvp-region/:pvpRegionId/pvp-season/:pvpSeasonId/pvp-reward/index",ginPvPRewardsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/realm/index",ginRealmsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/realm/:realmSlug",ginRealm)
 app.Instance().RegisterRoute("GET", "/data/wow/search/realm",ginRealmSearch)
 app.Instance().RegisterRoute("GET", "/data/wow/region/index",ginRegionsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/region/:regionId",ginRegion)
 app.Instance().RegisterRoute("GET", "/data/wow/token/index",ginWoWTokenIndex(CN))

}

func ginAuctionHouseIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginAuctions(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginConnectedRealmsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginConnectedRealm(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginConnectedRealmsSearch(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCreatureFamiliesIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCreatureFamily(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCreatureTypesIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCreatureType(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCreature(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCreatureSearch(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCreatureDisplayMedia(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCreatureFamilyMedia(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginGuildCrestComponentsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginGuildCrestBorderMedia(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginGuildCrestEmblemMedia(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginItemClassesIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginItemClass(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginItemSubclass(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginItem(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginItemMedia(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginItemSearch(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginMediaSearch(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPlayableClassesIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPlayableClass(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPlayableClassMedia(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPlayableRacesIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPlayableRace(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPowerTypesIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPowerType(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPvPSeasonsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPvPSeason(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPvPRegionIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPvPRegionalSeasonIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPvPRegionalSeason(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPvPLeaderboardsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPvPLeaderboard(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPvPRewardsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginRealmsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginRealm(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginRealmSearch(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginRegionsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginRegion(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginWoWTokenIndex(CN)(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
