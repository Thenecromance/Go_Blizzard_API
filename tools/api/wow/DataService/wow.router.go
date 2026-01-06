package wow_routers

import "Unofficial_API/app"

func init() {

 app.Instance().RegisterRoute("GET", "/data/wow/achievement/index",ginAchievementsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/achievement/:achievementId",ginAchievement)
 app.Instance().RegisterRoute("GET", "/data/wow/media/achievement/:achievementId",ginAchievementMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/achievement-category/index",ginAchievementCategoriesIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/achievement-category/:achievementCategoryId",ginAchievementCategory)
 app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/:connectedRealmId/auctions",ginAuctions)
 app.Instance().RegisterRoute("GET", "/data/wow/auctions/commodities",ginCommodities)
 app.Instance().RegisterRoute("GET", "/data/wow/azerite-essence/index",ginAzeriteEssencesIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/azerite-essence/:azeriteEssenceId",ginAzeriteEssence)
 app.Instance().RegisterRoute("GET", "/data/wow/search/azerite-essence",ginAzeriteEssenceSearch)
 app.Instance().RegisterRoute("GET", "/data/wow/media/azerite-essence/:azeriteEssenceId",ginAzeriteEssenceMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/index",ginConnectedRealmsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/:connectedRealmId",ginConnectedRealm)
 app.Instance().RegisterRoute("GET", "/data/wow/search/connected-realm",ginConnectedRealmSearch)
 app.Instance().RegisterRoute("GET", "/data/wow/covenant/index",ginCovenantIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/covenant/:covenantId",ginCovenant)
 app.Instance().RegisterRoute("GET", "/data/wow/media/covenant/:covenantId",ginCovenantMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/covenant/soulbind/index",ginSoulbindIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/covenant/soulbind/:soulbindId",ginSoulbind)
 app.Instance().RegisterRoute("GET", "/data/wow/covenant/conduit/index",ginConduitIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/covenant/conduit/:conduitId",ginConduit)
 app.Instance().RegisterRoute("GET", "/data/wow/creature/:creatureId",ginCreature)
 app.Instance().RegisterRoute("GET", "/data/wow/search/creature",ginCreatureSearch)
 app.Instance().RegisterRoute("GET", "/data/wow/media/creature-display/:creatureDisplayId",ginCreatureDisplayMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/creature-family/index",ginCreatureFamiliesIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/creature-family/:creatureFamilyId",ginCreatureFamily)
 app.Instance().RegisterRoute("GET", "/data/wow/media/creature-family/:creatureFamilyId",ginCreatureFamilyMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/creature-type/index",ginCreatureTypesIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/creature-type/:creatureTypeId",ginCreatureType)
 app.Instance().RegisterRoute("GET", "/data/wow/guild-crest/index",ginGuildCrestComponentsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/media/guild-crest/border/:borderId",ginGuildCrestBorderMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/media/guild-crest/emblem/:emblemId",ginGuildCrestEmblemMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/heirloom/index",ginHeirloomIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/heirloom/:heirloomId",ginHeirloom)
 app.Instance().RegisterRoute("GET", "/data/wow/decor/index",ginDecorIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/decor/:decorId",ginDecor)
 app.Instance().RegisterRoute("GET", "/data/wow/search/decor",ginDecorSearch)
 app.Instance().RegisterRoute("GET", "/data/wow/fixture/index",ginFixtureIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/fixture/:fixtureId",ginFixture)
 app.Instance().RegisterRoute("GET", "/data/wow/search/fixture",ginFixtureSearch)
 app.Instance().RegisterRoute("GET", "/data/wow/fixture-hook/index",ginFixtureHookIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/fixture-hook/:fixtureHookId",ginFixtureHook)
 app.Instance().RegisterRoute("GET", "/data/wow/search/fixture-hook",ginFixtureHookSearch)
 app.Instance().RegisterRoute("GET", "/data/wow/room/index",ginRoomIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/room/:roomId",ginRoom)
 app.Instance().RegisterRoute("GET", "/data/wow/search/room",ginRoomSearch)
 app.Instance().RegisterRoute("GET", "/data/wow/item/:itemId",ginItem)
 app.Instance().RegisterRoute("GET", "/data/wow/search/item",ginItemSearch)
 app.Instance().RegisterRoute("GET", "/data/wow/media/item/:itemId",ginItemMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/item-class/index",ginItemClassesIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/item-class/:itemClassId",ginItemClass)
 app.Instance().RegisterRoute("GET", "/data/wow/item-set/index",ginItemSetsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/item-set/:itemSetId",ginItemSet)
 app.Instance().RegisterRoute("GET", "/data/wow/item-class/:itemClassId/item-subclass/:itemSubclassId",ginItemSubclass)
 app.Instance().RegisterRoute("GET", "/data/wow/item-appearance/:appearanceId",ginItemAppearance)
 app.Instance().RegisterRoute("GET", "/data/wow/search/item-appearance",ginItemAppearanceSearch)
 app.Instance().RegisterRoute("GET", "/data/wow/item-appearance/set/index",ginItemAppearanceSetsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/item-appearance/set/:appearanceSetId",ginItemAppearanceSet)
 app.Instance().RegisterRoute("GET", "/data/wow/item-appearance/slot/index",ginItemAppearanceSlotIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/item-appearance/slot/:slotType",ginItemAppearanceSlot)
 app.Instance().RegisterRoute("GET", "/data/wow/journal-expansion/index",ginJournalExpansionsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/journal-expansion/:journalExpansionId",ginJournalExpansion)
 app.Instance().RegisterRoute("GET", "/data/wow/journal-encounter/index",ginJournalEncountersIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/journal-encounter/:journalEncounterId",ginJournalEncounter)
 app.Instance().RegisterRoute("GET", "/data/wow/search/journal-encounter",ginJournalEncounterSearch)
 app.Instance().RegisterRoute("GET", "/data/wow/journal-instance/index",ginJournalInstancesIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/journal-instance/:journalInstanceId",ginJournalInstance)
 app.Instance().RegisterRoute("GET", "/data/wow/media/journal-instance/:journalInstanceId",ginJournalInstanceMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/search/media",ginMediaSearch)
 app.Instance().RegisterRoute("GET", "/data/wow/modified-crafting/index",ginModifiedCraftingIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/modified-crafting/category/index",ginModifiedCraftingCategoryIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/modified-crafting/category/:categoryId",ginModifiedCraftingCategory)
 app.Instance().RegisterRoute("GET", "/data/wow/modified-crafting/reagent-slot-type/index",ginModifiedCraftingReagentSlotTypeIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/modified-crafting/reagent-slot-type/:slotTypeId",ginModifiedCraftingReagentSlotType)
 app.Instance().RegisterRoute("GET", "/data/wow/mount/index",ginMountsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/mount/:mountId",ginMount)
 app.Instance().RegisterRoute("GET", "/data/wow/search/mount",ginMountSearch)
 app.Instance().RegisterRoute("GET", "/data/wow/keystone-affix/index",ginMythicKeystoneAffixesIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/keystone-affix/:keystoneAffixId",ginMythicKeystoneAffix)
 app.Instance().RegisterRoute("GET", "/data/wow/media/keystone-affix/:keystoneAffixId",ginMythicKeystoneAffixMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/index",ginMythicKeystoneIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/dungeon/index",ginMythicKeystoneDungeonsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/dungeon/:dungeonId",ginMythicKeystoneDungeon)
 app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/period/index",ginMythicKeystonePeriodsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/period/:periodId",ginMythicKeystonePeriod)
 app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/season/index",ginMythicKeystoneSeasonsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/mythic-keystone/season/:seasonId",ginMythicKeystoneSeason)
 app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/:connectedRealmId/mythic-leaderboard/index",ginMythicKeystoneLeaderboardsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/connected-realm/:connectedRealmId/mythic-leaderboard/:dungeonId/period/:period",ginMythicKeystoneLeaderboard)
 app.Instance().RegisterRoute("GET", "/data/wow/leaderboard/hall-of-fame/:raid/:faction",ginMythicRaidLeaderboard)
 app.Instance().RegisterRoute("GET", "/data/wow/neighborhood-map/index",ginNeighborhoodMapIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/neighborhood-map/:neighborhoodMapId",ginNeighborhoodMap)
 app.Instance().RegisterRoute("GET", "/data/wow/neighborhood-map/:neighborhoodMapId/neighborhood/:neighborhoodId",ginNeighborhood)
 app.Instance().RegisterRoute("GET", "/data/wow/pet/index",ginPetsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/pet/:petId",ginPet)
 app.Instance().RegisterRoute("GET", "/data/wow/media/pet/:petId",ginPetMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/pet-ability/index",ginPetAbilitiesIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/pet-ability/:petAbilityId",ginPetAbility)
 app.Instance().RegisterRoute("GET", "/data/wow/media/pet-ability/:petAbilityId",ginPetAbilityMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/playable-class/index",ginPlayableClassesIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/playable-class/:classId",ginPlayableClass)
 app.Instance().RegisterRoute("GET", "/data/wow/media/playable-class/:playableClassId",ginPlayableClassMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/playable-class/:classId/pvp-talent-slots",ginPvPTalentSlots)
 app.Instance().RegisterRoute("GET", "/data/wow/playable-race/index",ginPlayableRacesIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/playable-race/:playableRaceId",ginPlayableRace)
 app.Instance().RegisterRoute("GET", "/data/wow/playable-specialization/index",ginPlayableSpecializationsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/playable-specialization/:specId",ginPlayableSpecialization)
 app.Instance().RegisterRoute("GET", "/data/wow/media/playable-specialization/:specId",ginPlayableSpecializationMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/power-type/index",ginPowerTypesIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/power-type/:powerTypeId",ginPowerType)
 app.Instance().RegisterRoute("GET", "/data/wow/profession/index",ginProfessionsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/profession/:professionId",ginProfession)
 app.Instance().RegisterRoute("GET", "/data/wow/media/profession/:professionId",ginProfessionMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/profession/:professionId/skill-tier/:skillTierId",ginProfessionSkillTier)
 app.Instance().RegisterRoute("GET", "/data/wow/recipe/:recipeId",ginRecipe)
 app.Instance().RegisterRoute("GET", "/data/wow/media/recipe/:recipeId",ginRecipeMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/pvp-season/index",ginPvPSeasonsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/pvp-season/:pvpSeasonId",ginPvPSeason)
 app.Instance().RegisterRoute("GET", "/data/wow/pvp-season/:pvpSeasonId/pvp-leaderboard/index",ginPvPLeaderboardsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/pvp-season/:pvpSeasonId/pvp-leaderboard/:pvpBracket",ginPvPLeaderboard)
 app.Instance().RegisterRoute("GET", "/data/wow/pvp-season/:pvpSeasonId/pvp-reward/index",ginPvPRewardsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/pvp-tier/index",ginPvPTiersIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/pvp-tier/:pvpTierId",ginPvPTier)
 app.Instance().RegisterRoute("GET", "/data/wow/media/pvp-tier/:pvpTierId",ginPvPTierMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/quest/index",ginQuestsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/quest/:questId",ginQuest)
 app.Instance().RegisterRoute("GET", "/data/wow/quest/category/index",ginQuestCategoriesIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/quest/category/:questCategoryId",ginQuestCategory)
 app.Instance().RegisterRoute("GET", "/data/wow/quest/area/index",ginQuestAreasIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/quest/area/:questAreaId",ginQuestArea)
 app.Instance().RegisterRoute("GET", "/data/wow/quest/type/index",ginQuestTypesIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/quest/type/:questTypeId",ginQuestType)
 app.Instance().RegisterRoute("GET", "/data/wow/realm/index",ginRealmsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/realm/:realmSlug",ginRealm)
 app.Instance().RegisterRoute("GET", "/data/wow/search/realm",ginRealmSearch)
 app.Instance().RegisterRoute("GET", "/data/wow/region/index",ginRegionsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/region/:regionId",ginRegion)
 app.Instance().RegisterRoute("GET", "/data/wow/reputation-faction/index",ginReputationFactionsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/reputation-faction/:reputationFactionId",ginReputationFaction)
 app.Instance().RegisterRoute("GET", "/data/wow/reputation-tiers/index",ginReputationTiersIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/reputation-tiers/:reputationTiersId",ginReputationTiers)
 app.Instance().RegisterRoute("GET", "/data/wow/spell/:spellId",ginSpell)
 app.Instance().RegisterRoute("GET", "/data/wow/media/spell/:spellId",ginSpellMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/search/spell",ginSpellSearch)
 app.Instance().RegisterRoute("GET", "/data/wow/talent-tree/index",ginTalentTreeIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/talent-tree/:talentTreeId/playable-specialization/:specId",ginTalentTree)
 app.Instance().RegisterRoute("GET", "/data/wow/talent-tree/:talentTreeId",ginTalentTreeNodes)
 app.Instance().RegisterRoute("GET", "/data/wow/talent/index",ginTalentsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/talent/:talentId",ginTalent)
 app.Instance().RegisterRoute("GET", "/data/wow/pvp-talent/index",ginPvPTalentsIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/pvp-talent/:pvpTalentId",ginPvPTalent)
 app.Instance().RegisterRoute("GET", "/data/wow/tech-talent-tree/index",ginTechTalentTreeIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/tech-talent-tree/:techTalentTreeId",ginTechTalentTree)
 app.Instance().RegisterRoute("GET", "/data/wow/tech-talent/index",ginTechTalentIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/tech-talent/:techTalentId",ginTechTalent)
 app.Instance().RegisterRoute("GET", "/data/wow/media/tech-talent/:techTalentId",ginTechTalentMedia)
 app.Instance().RegisterRoute("GET", "/data/wow/title/index",ginTitlesIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/title/:titleId",ginTitle)
 app.Instance().RegisterRoute("GET", "/data/wow/toy/index",ginToyIndex)
 app.Instance().RegisterRoute("GET", "/data/wow/toy/:toyId",ginToy)
 app.Instance().RegisterRoute("GET", "/data/wow/token/index",ginWoWTokenIndex(US,EU,KR,TW))
 app.Instance().RegisterRoute("GET", "/data/wow/token/index",ginWoWTokenIndex(CN))

}

func ginAchievementsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginAchievement(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginAchievementMedia(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginAchievementCategoriesIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginAchievementCategory(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginAuctions(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCommodities(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginAzeriteEssencesIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginAzeriteEssence(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginAzeriteEssenceSearch(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginAzeriteEssenceMedia(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginConnectedRealmsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginConnectedRealm(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginConnectedRealmSearch(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCovenantIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCovenant(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCovenantMedia(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginSoulbindIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginSoulbind(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginConduitIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginConduit(c *gin.Context){
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
func ginCreatureFamiliesIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCreatureFamily(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCreatureFamilyMedia(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCreatureTypesIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginCreatureType(c *gin.Context){
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
func ginHeirloomIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginHeirloom(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginDecorIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginDecor(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginDecorSearch(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginFixtureIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginFixture(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginFixtureSearch(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginFixtureHookIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginFixtureHook(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginFixtureHookSearch(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginRoomIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginRoom(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginRoomSearch(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginItem(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginItemSearch(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginItemMedia(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginItemClassesIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginItemClass(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginItemSetsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginItemSet(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginItemSubclass(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginItemAppearance(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginItemAppearanceSearch(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginItemAppearanceSetsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginItemAppearanceSet(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginItemAppearanceSlotIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginItemAppearanceSlot(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginJournalExpansionsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginJournalExpansion(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginJournalEncountersIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginJournalEncounter(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginJournalEncounterSearch(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginJournalInstancesIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginJournalInstance(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginJournalInstanceMedia(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginMediaSearch(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginModifiedCraftingIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginModifiedCraftingCategoryIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginModifiedCraftingCategory(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginModifiedCraftingReagentSlotTypeIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginModifiedCraftingReagentSlotType(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginMountsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginMount(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginMountSearch(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginMythicKeystoneAffixesIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginMythicKeystoneAffix(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginMythicKeystoneAffixMedia(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginMythicKeystoneIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginMythicKeystoneDungeonsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginMythicKeystoneDungeon(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginMythicKeystonePeriodsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginMythicKeystonePeriod(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginMythicKeystoneSeasonsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginMythicKeystoneSeason(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginMythicKeystoneLeaderboardsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginMythicKeystoneLeaderboard(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginMythicRaidLeaderboard(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginNeighborhoodMapIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginNeighborhoodMap(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginNeighborhood(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPetsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPet(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPetMedia(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPetAbilitiesIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPetAbility(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPetAbilityMedia(c *gin.Context){
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
func ginPvPTalentSlots(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPlayableRacesIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPlayableRace(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPlayableSpecializationsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPlayableSpecialization(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPlayableSpecializationMedia(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPowerTypesIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPowerType(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginProfessionsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginProfession(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginProfessionMedia(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginProfessionSkillTier(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginRecipe(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginRecipeMedia(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPvPSeasonsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPvPSeason(c *gin.Context){
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
func ginPvPTiersIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPvPTier(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPvPTierMedia(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginQuestsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginQuest(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginQuestCategoriesIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginQuestCategory(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginQuestAreasIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginQuestArea(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginQuestTypesIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginQuestType(c *gin.Context){
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
func ginReputationFactionsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginReputationFaction(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginReputationTiersIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginReputationTiers(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginSpell(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginSpellMedia(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginSpellSearch(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginTalentTreeIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginTalentTree(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginTalentTreeNodes(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginTalentsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginTalent(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPvPTalentsIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginPvPTalent(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginTechTalentTreeIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginTechTalentTree(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginTechTalentIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginTechTalent(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginTechTalentMedia(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginTitlesIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginTitle(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginToyIndex(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginToy(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginWoWTokenIndex(US,EU,KR,TW)(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginWoWTokenIndex(CN)(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
