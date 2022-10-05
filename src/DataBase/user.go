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
	SelectByWhere(page Page) []UserInfo
	Update() int
	Delete() int
	Insert() int
}

func SelectById(id string) UserInfo {
	var userInfo UserInfo

	return userInfo
}
func (info UserInfo) SelectByWhere(page Page) []UserInfo {
	var userinfo = make([]UserInfo, page.Size)

	return userinfo
}
func (info UserInfo) SelectOnLogin(user User) string {
	return "123"
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
func (u User) Select() (string, error) {
	var userinfo []UserInfo
	err := Db.Select(&userinfo, "select uid from user where userName=? and passWord=?", u.Username, u.Password)
	if err != nil {
		return "", err
	}
	if len(userinfo) > 0 {
		return userinfo[0].Uid, nil
	}
	return "", nil
}
