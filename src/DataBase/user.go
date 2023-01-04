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

type UserApi struct {
}
type UserDao interface {
	SelectByWhere() []model.UserInfo
	SelectById() model.UserInfo
	Update() bool
	Delete() bool
	Insert() bool
}

func (u *UserApi) SelectById(id string) (model.UserInfo, error) {
	var userInfo model.UserInfo
	if res := Db.Table("user").Select("UserName", "UIpCode", "age", "sex", "PhoneNum").Where("uid=?", id).First(&userInfo); res.Error != nil {
		return userInfo, res.Error
	}
	return userInfo, nil
}
func (u *UserApi) SelectByWhere(info UserInfo, size int) []model.UserInfo {
	var userinfo = make([]model.UserInfo, size)

	if info.UserName != "" {

	}
	return userinfo
}
func (u *UserApi) SelectOnLogin(user User) string {
	return "123"
}
func (u *UserApi) Update() bool {
	return false
}
func (u *UserApi) Delete() bool {

	return false
}
func (u *UserApi) Insert(info UserInfo) bool {
	if result := Db.Create(&info); result.Error == nil {
		fmt.Println("插入成功")
		return true
	}
	return false
}

func (u *UserApi) Select(user User) (string, error) {
	var userinfo UserInfo
	res := Db.Table("user").Where("username=? and password=?", user.Username, user.Password).First(&userinfo)
	if res.Error != nil {
		return "", res.Error
	}
	return userinfo.Uid, nil
}
