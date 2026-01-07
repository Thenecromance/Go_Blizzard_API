package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/DataService/Quest"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/quest/index", ginQuestsIndex) //QuestsIndex Returns the parent index for quests.

	app.Instance().RegisterRoute("GET", "/data/wow/quest/:questId", ginQuest) //Quest Returns a quest by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/quest/category/index", ginQuestCategoriesIndex) //QuestCategoriesIndex Returns an index of quest categories (such as quests for a specific class, profession, or storyline).

	app.Instance().RegisterRoute("GET", "/data/wow/quest/category/:questCategoryId", ginQuestCategory) //QuestCategory Returns a quest category by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/quest/area/index", ginQuestAreasIndex) //QuestAreasIndex Returns an index of quest areas.

	app.Instance().RegisterRoute("GET", "/data/wow/quest/area/:questAreaId", ginQuestArea) //QuestArea Returns a quest area by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/quest/type/index", ginQuestTypesIndex) //QuestTypesIndex Returns an index of quest types (such as PvP quests, raid quests, or account quests).

	app.Instance().RegisterRoute("GET", "/data/wow/quest/type/:questTypeId", ginQuestType) //QuestType Returns a quest type by ID.

}




func ginQuestsIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_Quest.QuestsIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Quest.QuestsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginQuest(c *gin.Context) {
	// binding uri parameters
	var req wow_Quest.QuestFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Quest.Quest(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginQuestCategoriesIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_Quest.QuestCategoriesIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Quest.QuestCategoriesIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginQuestCategory(c *gin.Context) {
	// binding uri parameters
	var req wow_Quest.QuestCategoryFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Quest.QuestCategory(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginQuestAreasIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_Quest.QuestAreasIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Quest.QuestAreasIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginQuestArea(c *gin.Context) {
	// binding uri parameters
	var req wow_Quest.QuestAreaFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Quest.QuestArea(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginQuestTypesIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_Quest.QuestTypesIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Quest.QuestTypesIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginQuestType(c *gin.Context) {
	// binding uri parameters
	var req wow_Quest.QuestTypeFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Quest.QuestType(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
