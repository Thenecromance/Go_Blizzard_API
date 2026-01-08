package routers

import (
	"net/http"

	"Unofficial_API/app"
	"Unofficial_API/api/wow/DataService/Pet"

	"github.com/gin-gonic/gin"
)

func init() {

	app.Instance().RegisterRoute("GET", "/data/wow/pet/index", ginPetPetsIndex) /* PetsIndex Returns an index of battle pets. */

	app.Instance().RegisterRoute("GET", "/data/wow/pet/:petId", ginPetPet) /* Pet Returns a battle pets by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/media/pet/:petId", ginPetPetMedia) /* PetMedia Returns media for a battle pet by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/pet-ability/index", ginPetPetAbilitiesIndex) /* PetAbilitiesIndex Returns an index of pet abilities. */

	app.Instance().RegisterRoute("GET", "/data/wow/pet-ability/:petAbilityId", ginPetPetAbility) /* PetAbility Returns a pet ability by ID. */

	app.Instance().RegisterRoute("GET", "/data/wow/media/pet-ability/:petAbilityId", ginPetPetAbilityMedia) /* PetAbilityMedia Returns media for a pet ability by ID. */

}




func ginPetPetsIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_Pet.PetsIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Pet.PetsIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginPetPet(c *gin.Context) {
	// binding uri parameters
	var req wow_Pet.PetFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Pet.Pet(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginPetPetMedia(c *gin.Context) {
	// binding uri parameters
	var req wow_Pet.PetMediaFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Pet.PetMedia(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginPetPetAbilitiesIndex(c *gin.Context) {
	// binding uri parameters
	var req wow_Pet.PetAbilitiesIndexFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Pet.PetAbilitiesIndex(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginPetPetAbility(c *gin.Context) {
	// binding uri parameters
	var req wow_Pet.PetAbilityFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Pet.PetAbility(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}



func ginPetPetAbilityMedia(c *gin.Context) {
	// binding uri parameters
	var req wow_Pet.PetAbilityMediaFields
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// binding query parameters
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resp, err := wow_Pet.PetAbilityMedia(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
