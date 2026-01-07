package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/achievement/index", ginAchievements_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/achievement/:achievementId", ginAchievement)
	app.Instance().RegisterRoute("GET", "/data/wow/media/achievement/:achievementId", ginAchievement_Media)
	app.Instance().RegisterRoute("GET", "/data/wow/achievement-category/index", ginAchievement_Categories_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/achievement-category/:achievementCategoryId", ginAchievement_Category)

}

func ginAchievements_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginAchievement(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginAchievement_Media(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginAchievement_Categories_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginAchievement_Category(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
