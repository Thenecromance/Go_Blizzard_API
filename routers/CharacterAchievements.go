package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wowClassic/ProfileService/CharacterAchievements"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/achievements", ginCharacterAchievementsSummary) //CharacterAchievementsSummary Returns a summary of the achievements a character has completed.

	app.Instance().RegisterRoute("GET", "/profile/wow/character/:realmSlug/:characterName/achievements/statistics", ginCharacterAchievementStatistics) //CharacterAchievementStatistics Returns a character's statistics as they pertain to achievements.

}




func ginCharacterAchievementsSummary(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_CharacterAchievements.CharacterAchievementsSummaryFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_CharacterAchievements.CharacterAchievementsSummary(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginCharacterAchievementStatistics(c *gin.Context) {
	// binding uri parameters
	var req wowClassic_CharacterAchievements.CharacterAchievementStatisticsFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wowClassic_CharacterAchievements.CharacterAchievementStatistics(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
