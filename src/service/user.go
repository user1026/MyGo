package service

import (
	"gin01/src/DataBase"
	"gin01/src/middleware"
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

func AddUser(user DataBase.UserInfo) (bool, error) {
	user.CreateTime = utils.GetNowTime()
	user.Uid = strconv.FormatInt(time.Now().Unix(), 10)
	_, err := user.Insert()
	if err != nil {
		return false, err
	}
	return true, nil
}
func EditUser(user DataBase.UserInfo) {

}
func DelUser(user DataBase.UserInfo) error {
	err := user.Delete()
	if err != nil {
		return err
	}
	return nil
}
func SelectUser() {

}
