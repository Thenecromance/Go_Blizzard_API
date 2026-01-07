package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/DataService/HousingDecor"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/decor/index", ginDecorIndex) //DecorIndex Returns an index of decor.

	app.Instance().RegisterRoute("GET", "/data/wow/decor/:decorId", ginDecor) //Decor Returns a decor by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/search/decor", ginDecorSearch) //DecorSearch Performs a search of decor. The fields below are provided for example. For more detail see the <a href="/documentation/world-of-warcraft/guides/search">Search Guide</a>.

	app.Instance().RegisterRoute("GET", "/data/wow/fixture/index", ginFixtureIndex) //FixtureIndex Returns an index of fixtures.

	app.Instance().RegisterRoute("GET", "/data/wow/fixture/:fixtureId", ginFixture) //Fixture Returns a fixture by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/search/fixture", ginFixtureSearch) //FixtureSearch Performs a search of fixtures. The fields below are provided for example. For more detail see the <a href="/documentation/world-of-warcraft/guides/search">Search Guide</a>.

	app.Instance().RegisterRoute("GET", "/data/wow/fixture-hook/index", ginFixtureHookIndex) //FixtureHookIndex Returns an index of fixture hooks.

	app.Instance().RegisterRoute("GET", "/data/wow/fixture-hook/:fixtureHookId", ginFixtureHook) //FixtureHook Returns a fixture hook by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/search/fixture-hook", ginFixtureHookSearch) //FixtureHookSearch Performs a search of fixture hooks. The fields below are provided for example. For more detail see the <a href="/documentation/world-of-warcraft/guides/search">Search Guide</a>.

	app.Instance().RegisterRoute("GET", "/data/wow/room/index", ginRoomIndex) //RoomIndex Returns an index of rooms.

	app.Instance().RegisterRoute("GET", "/data/wow/room/:roomId", ginRoom) //Room Returns a room by ID.

	app.Instance().RegisterRoute("GET", "/data/wow/search/room", ginRoomSearch) //RoomSearch Performs a search of rooms. The fields below are provided for example. For more detail see the <a href="/documentation/world-of-warcraft/guides/search">Search Guide</a>.

}




func ginDecorIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_HousingDecor.DecorIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_HousingDecor.DecorIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginDecor(c *gin.Context) {
	// binding uri parameters
	var req wow_HousingDecor.DecorFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_HousingDecor.Decor(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginDecorSearch(c *gin.Context) {
	// binding uri parameters
	var req wow_HousingDecor.DecorSearchFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_HousingDecor.DecorSearch(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginFixtureIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_HousingDecor.FixtureIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_HousingDecor.FixtureIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginFixture(c *gin.Context) {
	// binding uri parameters
	var req wow_HousingDecor.FixtureFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_HousingDecor.Fixture(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginFixtureSearch(c *gin.Context) {
	// binding uri parameters
	var req wow_HousingDecor.FixtureSearchFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_HousingDecor.FixtureSearch(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginFixtureHookIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_HousingDecor.FixtureHookIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_HousingDecor.FixtureHookIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginFixtureHook(c *gin.Context) {
	// binding uri parameters
	var req wow_HousingDecor.FixtureHookFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_HousingDecor.FixtureHook(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginFixtureHookSearch(c *gin.Context) {
	// binding uri parameters
	var req wow_HousingDecor.FixtureHookSearchFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_HousingDecor.FixtureHookSearch(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginRoomIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_HousingDecor.RoomIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_HousingDecor.RoomIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginRoom(c *gin.Context) {
	// binding uri parameters
	var req wow_HousingDecor.RoomFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_HousingDecor.Room(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginRoomSearch(c *gin.Context) {
	// binding uri parameters
	var req wow_HousingDecor.RoomSearchFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_HousingDecor.RoomSearch(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
