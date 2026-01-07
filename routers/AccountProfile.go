package routers

import (
	"Unofficial_API/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

 app.Instance().RegisterRoute("GET", "/profile/user/wow", ginAccount_Profile_Summary)
 app.Instance().RegisterRoute("GET", "/profile/user/wow/protected-character/:realmId-characterId", ginProtected_Character_Profile_Summary)

}

func ginAccount_Profile_Summary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
func ginProtected_Character_Profile_Summary(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{})
}
