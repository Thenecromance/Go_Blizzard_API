package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/DataService/Achievement"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/achievement/index", ginAchievementsIndex) //AchievementsIndex Returns an index of achievements.

	app.Instance().RegisterRoute("GET", "/data/wow/achievement/:achievementId", ginAchievement) //Achievement Returns an achievement by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/media/achievement/:achievementId", ginAchievementMedia) //AchievementMedia Returns media for an achievement by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/achievement-category/index", ginAchievementCategoriesIndex) //AchievementCategoriesIndex Returns an index of achievement categories.

	app.Instance().RegisterRoute("GET", "/data/wow/achievement-category/:achievementCategoryId", ginAchievementCategory) //AchievementCategory Returns an achievement category by ID.

}




func ginAchievementsIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_Achievement.AchievementsIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Achievement.AchievementsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginAchievement(c *gin.Context) {
	// binding uri parameters
	var req wow_Achievement.AchievementFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Achievement.Achievement(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginAchievementMedia(c *gin.Context) {
	// binding uri parameters
	var req wow_Achievement.AchievementMediaFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Achievement.AchievementMedia(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginAchievementCategoriesIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_Achievement.AchievementCategoriesIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Achievement.AchievementCategoriesIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginAchievementCategory(c *gin.Context) {
	// binding uri parameters
	var req wow_Achievement.AchievementCategoryFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Achievement.AchievementCategory(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
