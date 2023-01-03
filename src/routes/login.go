package routes

import (
	"fmt"
	"gin01/src/DataBase"
	"gin01/src/service"
	"gin01/src/utils"
	"github.com/gin-gonic/gin"
)

func LoginRoute(e *gin.Engine) {
	e.POST("/login", Login)
	//e.POST("/getuserinfo", service.Getuserinfo)
	e.POST("/logout", Logout)
}
func Login(c *gin.Context) {
	var form DataBase.User
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(400, gin.H{"message": "缺少必要参数"})
		return
	}
	uid, uidErr := service.GetUid(form)
	token, tokenErr := service.GetToken(uid)
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
	var form DataBase.User
	err := c.ShouldBindJSON(&form)
	if err != nil {
		return
	}
	uid, err := form.Select()
	if err != nil {
		c.JSON(500, gin.H{"message": "程序出错"})
		return
	}
	if uid == "" {

	}

	c.JSON(200, gin.H{"uid": uid})
}
