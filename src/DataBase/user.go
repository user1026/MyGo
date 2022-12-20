package DataBase

import (
	"gin01/src/model/VO"
)

type UserInfo struct {
	VO.UserInfo
	VO.Page
}
type User struct {
	VO.User
}
type UserDao interface {
	SelectByWhere() []UserInfo
	Update() int
	Delete() int
	Insert() int
}

func SelectById(id string) UserInfo {
	var userInfo UserInfo
	Db.Select(&userInfo, "", id)
	return userInfo
}
func (info UserInfo) SelectByWhere() []UserInfo {
	var userinfo = make([]UserInfo, info.Size)

	return userinfo
}
func (info UserInfo) SelectOnLogin(user User) string {
	return "123"
}
func (info UserInfo) Update() error {
	return nil
}
func (info UserInfo) Delete() error {
	_, err := Db.Exec("delete from user where uid=?", info.Uid)
	if err != nil {
		return err
	}
	return nil
}
func (info UserInfo) Insert() (int, error) {
	_, err := Db.Exec("insert into user (uid,userName,uIpCode,createTime,updateTime,passWord,roleId,sex,age,phoneNum) values (?,?,?,?,?,?,?,?,?,?)", info.Uid, info.UserName, info.UipCode, info.CreateTime, info.CreateTime, "123456", info.RoleId, info.Sex, info.Age, info.PhoneNum)
	if err != nil {
		return 0, err
	}
	return 1, nil
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
