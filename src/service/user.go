package service

import (
	"gin01/src/DataBase"
	"gin01/src/middleware"
	"gin01/src/model"
	"gin01/src/utils"
	"strconv"
	"time"
)

type UserService struct {
}

var dataApi = DataBase.DataApi

func (u *UserService) GetUid(user DataBase.User) (string, error) {
	uid, err := dataApi.Select(user)
	if err != nil {
		return "", err
	}
	return uid, nil
}
func (u *UserService) GetToken(uid string) (string, error) {
	token, err := middleware.CreateToken(uid)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *UserService) AddUser(user DataBase.UserInfo) bool {
	user.CreateTime = utils.GetNowTime()
	user.UpdateTime = utils.GetNowTime()
	user.Password = "123456"
	user.Uid = strconv.FormatInt(time.Now().Unix(), 10)
	res := dataApi.Insert(user)
	if res == true {
		return true
	}
	return false
}
func (u *UserService) EditUser(user DataBase.UserInfo) {

}
func (u *UserService) DelUser(user DataBase.UserInfo) bool {
	res := dataApi.Delete()
	if res == true {
		return true
	}
	return false
}
func SelectUser() {

}
func (u *UserService) SelectUserById(id string) (model.UserInfo, error) {
	userinfo, err := dataApi.SelectById(id)
	if err != nil {
		return userinfo, err
	}
	return userinfo, nil
}
