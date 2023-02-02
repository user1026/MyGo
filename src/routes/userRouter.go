package routes

import (
	"fmt"
	"gin01/src/model"
	"gin01/src/utils"
	"github.com/gin-gonic/gin"
)

func UserRouter(e *gin.Engine) {
	user := e.Group("/user")
	{
		user.POST("/add", AddUser)
		user.POST("/delete", DeleteUser)
		user.POST("/edit", EditUser)
		user.POST("/query", QueryUser)
		user.POST("/queryById", QueryUserById)
	}
}

func AddUser(c *gin.Context) {
	var userInfo model.UserInfo
	if getUserInfoErr := c.ShouldBindJSON(&userInfo); getUserInfoErr != nil {
		fmt.Println(userInfo)
		c.JSON(400, gin.H{"message": "缺乏必要参数"})
		return
	}
	res := serv.AddUser(userInfo)
	if res != true {
		c.JSON(500, gin.H{"message": "添加失败"})
		return
	} else {
		c.JSON(200, gin.H{"message": "添加成功"})
	}
}
func EditUser(c *gin.Context) {
	var userinfo model.EditUserInfo
	if err := c.ShouldBindJSON(&userinfo); err != nil || userinfo.Uid == "" {
		utils.FailWithMsg("参数错误", c)
		return
	}
	if isEdit := serv.EditUser(userinfo); isEdit != true {
		utils.FailWithMsg("修改失败", c)
		return
	}
	utils.OKWithMsg("修改成功", c)
}
func DeleteUser(c *gin.Context) {
	var userInfo model.UserInfo
	if err := c.ShouldBindJSON(&userInfo); err != nil || userInfo.Uid == "" {
		utils.FailWithMsg("uid获取失败", c)
		return
	}
	if isDel, msg := serv.DelUser(userInfo); isDel == false {
		utils.FailWithMsg(msg, c)
		return
	}
	utils.OKWithMsg("删除成功", c)
}
func QueryUser(c *gin.Context) {

}
func QueryUserById(c *gin.Context) {
	var user model.UserInfo
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.FailWithMsg("未获取到uid", c)
		return
	}
	userinfo, err := serv.SelectUserById(user.Uid)
	if err != nil {
		utils.FailWithMsg("查询失败", c)
		return
	}
	utils.OKWithData(map[string]interface{}{"userinfo": userinfo}, c)
}
