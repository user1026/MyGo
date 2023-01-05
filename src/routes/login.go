package routes

import (
	"fmt"
	"gin01/src/model"
	"gin01/src/utils"
	"github.com/gin-gonic/gin"
)

func LoginRoute(e *gin.Engine) {
	e.POST("/login", Login)
	//e.POST("/getuserinfo", service.Getuserinfo)
	e.POST("/logout", Logout)
}
func Login(c *gin.Context) {
	var form model.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(400, gin.H{"message": "缺少必要参数"})
		return
	}
	uid, uidErr := serv.GetUid(form)
	token, tokenErr := serv.GetToken(uid)
	if uidErr != nil || tokenErr != nil {
		utils.FailWithMsg("登陆失败，稍后重试", c)
		//c.JSON(500, gin.H{"code": 500, "message": "登陆失败，稍后重试"})
		return
	}
	var t = map[string]string{"token": token}
	utils.OKWithData(t, c)
	//c.JSON(200, gin.H{"token": token})
}
func Logout(c *gin.Context) {
	fmt.Println("123")
	c.JSON(200, gin.H{"c": 123})
	fmt.Println("3333")
}
func TestUrl(c *gin.Context) {

}
