package service

import (
	"gin01/src/DataBase"
	"gin01/src/model"
	"github.com/gin-gonic/gin"
)

func GetUid(user DataBase.User) (string, error) {
	uid, err := user.Select()
	if err != nil {
		return "", err
	}
	return uid, nil
}
func GetToken(uid string) (string, error) {
	token, err := model.CreateToken(uid)
	if err != nil {
		return "", err
	}
	return token, nil
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
