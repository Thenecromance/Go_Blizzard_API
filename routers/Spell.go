package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/spell/:spellId", ginSpell)
	app.Instance().RegisterRoute("GET", "/data/wow/media/spell/:spellId", ginSpell_Media)
	app.Instance().RegisterRoute("GET", "/data/wow/search/spell", ginSpell_Search)

}

func ginSpell(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginSpell_Media(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginSpell_Search(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
