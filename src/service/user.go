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

func (u *UserService) GetUid(user model.Login) (string, error) {
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

func (u *UserService) AddUser(user model.UserInfo) bool {
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
func (u *UserService) EditUser(user model.EditUserInfo) bool {
	user.UpdateTime = utils.GetNowTime()
	if res := dataApi.Update(user); res {
		return true
	}
	return false
}
func (u *UserService) DelUser(user model.UserInfo) (bool, string) {
	_, err := dataApi.SelectById(user.Uid)
	if err != nil {
		return false, "数据库无匹配数据"
	}
	res := dataApi.Delete(user.Uid)
	if res == true {
		return true, ""
	}
	return false, "删除失败"
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
