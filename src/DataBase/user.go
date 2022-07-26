package DataBase

import (
	"fmt"
)

type UserInfo struct {
	Username   string `form:"username" json:"username"`
	Uid        string `form:"uid" json:"uid"`
	Age        string `form:"age" json:"age"`
	Sex        string `form:"sex" json:"sex"`
	UipCode    string `form:"uipcode" json:"uipCode"`
	RoleId     string `form:"roleId" json:"roleId"`
	CreateTime string `form:"createTime" json:"createTime"`
	UpdateTime string `form:"updateTime" json:"updateTime"`
}

type User struct {
	Username string `db:"username" form:"username" json:"username" binding:"required"`
	Password string `db:"password" form:"password" json:"password" binding:"required"`
}

func CheckEffectUser(uid string) bool {
	var user []UserInfo
	err := Db.Select(&user, "select uid,username,roleid from user where uid=?", uid)
	if err != nil {
		fmt.Println(err, "selectErr")
		return false
	}
	if len(user) < 1 {
		return false
	}
	return true
}
func UserSelect(userForm User) UserInfo {
	var user []UserInfo
	err := Db.Select(&user, "select uid,username,age,roleid from user where username=? && password=?", userForm.Username, userForm.Password)
	if err != nil {
		fmt.Println(err, "selectErr")
		return UserInfo{}
	}
	if len(user) < 1 {
		return UserInfo{}
	}
	return user[0]
}
func EditUser(userInfo UserInfo) string {
	res, err := Db.Exec("update user set age=? where username=?", userInfo.Age, userInfo.Username)
	if err != nil {
		fmt.Println(err, "editErr")
		return "出错了"
	}
	rus, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err, "err")
	}
	fmt.Println(rus)
	return "123"
}
