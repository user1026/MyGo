package DataBase

import (
	"fmt"
	"gin01/src/model"
)

type UserInfo struct {
	model.UserInfo
}

func (UserInfo) TableName() string {
	return "user"
}

type UserInfoByPage struct {
	model.UserInfo
	model.Page
}
type User struct {
	model.User
}

func (User) TableName() string {
	return "user"
}

type UserDao interface {
	SelectByWhere() []model.UserInfo
	SelectById() model.UserInfo
	Update() bool
	Delete() bool
	Insert() bool
}

func SelectById(id string) (model.UserInfo, error) {
	var userInfo model.UserInfo
	if res := Db.Table("user").Select("UserName", "UIpCode", "age", "sex", "PhoneNum").Where("uid=?", id).First(&userInfo); res.Error != nil {
		return userInfo, res.Error
	}
	return userInfo, nil
}
func (info UserInfoByPage) SelectByWhere() []model.UserInfo {
	var userinfo = make([]model.UserInfo, info.Size)

	if info.UserName != "" {

	}
	return userinfo
}
func (info UserInfo) SelectOnLogin(user User) string {
	return "123"
}
func (info UserInfo) Update() bool {
	return false
}
func (info UserInfo) Delete() bool {

	return false
}
func (info UserInfo) Insert() bool {
	if result := Db.Create(&info); result.Error == nil {
		fmt.Println("插入成功")
		return true
	}
	return false
}

func (u User) Select() (string, error) {
	var userinfo UserInfo
	res := Db.Table("user").Where("username=? and password=?", u.Username, u.Password).First(&userinfo)
	if res.Error != nil {
		return "", res.Error
	}
	return userinfo.Uid, nil
}
