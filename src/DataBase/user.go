package DataBase

import (
	"fmt"
	"gin01/src/model"
)

type UserInfo struct {
	model.UserInfo
}
type AddUserInfo struct {
	model.UserInfo
	Password string `db:"passWord" form:"password"  gorm:"column:Password"`
}

func (AddUserInfo) TableName() string {
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
	Update() bool
	Delete() bool
	Insert() bool
}

func SelectById(id string) model.UserInfo {
	var userInfo model.UserInfo
	Db.Select(&userInfo, "", id)
	return userInfo
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
func (info AddUserInfo) Insert() bool {
	if result := Db.Create(&info); result.Error == nil {
		fmt.Println("插入成功")
		return true
	}
	return false
}

func (u User) Select() (string, error) {

	return "", nil
}
