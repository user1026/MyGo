package DataBase

import "gin01/src/model/VO"

type UserInfo struct {
	VO.UserInfo
}
type Page struct {
	VO.Page
}
type User struct {
	VO.User
}
type UserDao interface {
	SelectById(id string) UserInfo
	SelectByWhere(page Page) []UserInfo
	Update() int
	Delete() int
	Insert() int
}

func (info UserInfo) SelectById(id string) UserInfo {
	var userInfo UserInfo

	return userInfo
}
func (info UserInfo) SelectByWhere(page Page) []UserInfo {
	var userinfo = make([]UserInfo, page.Size)
	return userinfo
}
func (info UserInfo) Update() int {
	return 1
}
func (info UserInfo) Delete() int {
	return 1
}
func (info UserInfo) Insert() int {
	return 1
}
func (u User) Select() bool {
	return true
}
