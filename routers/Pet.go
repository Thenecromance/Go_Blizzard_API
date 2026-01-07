package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/pet/index", ginPets_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/pet/:petId", ginPet)
	app.Instance().RegisterRoute("GET", "/data/wow/media/pet/:petId", ginPet_Media)
	app.Instance().RegisterRoute("GET", "/data/wow/pet-ability/index", ginPet_Abilities_Index)
	app.Instance().RegisterRoute("GET", "/data/wow/pet-ability/:petAbilityId", ginPet_Ability)
	app.Instance().RegisterRoute("GET", "/data/wow/media/pet-ability/:petAbilityId", ginPet_Ability_Media)

}

func ginPets_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginPet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginPet_Media(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginPet_Abilities_Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginPet_Ability(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func ginPet_Ability_Media(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
