package DataBase

import (
	"fmt"
	"gin01/src/model"
)

type UserApi struct {
}
type UserDao interface {
	SelectByWhere() []model.UserInfo
	SelectById() (model.UserInfo, error)
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
func (u *UserApi) SelectByWhere(info model.UserInfo, size int) []model.UserInfo {
	var userinfo = make([]model.UserInfo, size)

	if info.UserName != "" {

	}
	return userinfo
}
func (u *UserApi) SelectOnLogin(user model.Login) string {
	return "123"
}
func (u *UserApi) Update(userinfo model.EditUserInfo) bool {
	if res := Db.Table("user").Where("uid=?", userinfo.Uid).Save(&userinfo); res.Error == nil {
		return true
	}
	return false
}
func (u *UserApi) Delete(id string) bool {
	var userinfo model.UserInfo
	res := Db.Table("user").Where("uid=?", id).Delete(&userinfo)
	if res.Error == nil {
		return true
	}
	return false
}
func (u *UserApi) DeleteByWhere() bool {
	return false
}
func (u *UserApi) Insert(info model.UserInfo) bool {
	if result := Db.Create(&info); result.Error == nil {
		fmt.Println("插入成功")
		return true
	}
	return false
}

func (u *UserApi) Select(user model.Login) (string, error) {
	var userinfo model.UserInfo
	res := Db.Table("user").Where("username=? and password=?", user.Username, user.Password).First(&userinfo)
	if res.Error != nil {
		return "", res.Error
	}
	return userinfo.Uid, nil
}
