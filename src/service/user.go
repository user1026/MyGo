package service

import (
	"gin01/src/DataBase"
	"github.com/gin-gonic/gin"
)

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
func Login(c *gin.Context) {
	var form DataBase.User
	err := c.ShouldBindJSON(&form)
	if err != nil {
		c.JSON(400, gin.H{"error": "缺少必要数据"})
		return
	}
	user, e := form.Select()
	if e != nil {

		return
	}
	if user != "" {
		c.JSON(200, gin.H{"code": 200})
	}
	//tokenString := createToken(&user)

	//c.String(http.StatusOK, "/index")
}

func Getuserinfo(c *gin.Context) {

}

func AddUser(c *gin.Context) {

}
func EditUser(c *gin.Context) {
	//var form DataBase.UserInfo
	//err := c.ShouldBindJSON(&form)
	//if err != nil {
	//
	//}
	//message := DataBase.EditUser(form)
	//c.JSON(200, gin.H{message: message})
}
func DelUser(c *gin.Context) {

}
func SelectUser(c *gin.Context) {

}
