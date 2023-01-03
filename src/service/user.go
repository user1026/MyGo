package service

import (
	"gin01/src/DataBase"
	"gin01/src/middleware"
	"gin01/src/model"
	"gin01/src/utils"
	"strconv"
	"time"
)

func GetUid(user DataBase.User) (string, error) {
	uid, err := user.Select()
	if err != nil {
		return "", err
	}
	return uid, nil
}
func GetToken(uid string) (string, error) {
	token, err := middleware.CreateToken(uid)
	if err != nil {
		return "", err
	}
	return token, nil
}

func AddUser(user DataBase.UserInfo) bool {
	user.CreateTime = utils.GetNowTime()
	user.UpdateTime = utils.GetNowTime()
	user.Password = "123456"
	user.Uid = strconv.FormatInt(time.Now().Unix(), 10)
	res := user.Insert()
	if res == true {
		return true
	}
	return false
}
func EditUser(user DataBase.UserInfo) {

}
func DelUser(user DataBase.UserInfo) bool {
	res := user.Delete()
	if res == true {
		return true
	}
	return false
}
func SelectUser() {

}
func SelectUserById(id string) (model.UserInfo, error) {
	userinfo, err := DataBase.SelectById(id)
	if err != nil {
		return userinfo, err
	}
	return userinfo, nil
}
