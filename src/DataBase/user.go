package DataBase

type UserInfo struct {
	Username   string `form:"username" json:"username"`
	Uid        string `form:"uid" json:"uid" binding:"require"`
	Age        string `form:"age" json:"age"`
	Sex        string `form:"sex" json:"sex"`
	UipCode    string `form:"uipcode" json:"uipCode"`
	RoleId     string `form:"roleId" json:"roleId"`
	CreateTime string `form:"createTime" json:"createTime"`
	UpdateTime string `form:"updateTime" json:"updateTime"`
}
type Page struct {
	Size string `json:"size" form:"size"`
	Page string `json:"page" form:"page"`
}
type User struct {
	Username string `db:"username" form:"username" json:"username" binding:"required"`
	Password string `db:"password" form:"password" json:"password" binding:"required"`
}
type UserDao interface {
	Select(page Page) []UserInfo
	Update() int
	Delete() int
	Insert() int
}

func (info UserInfo) Select(page Page) []UserInfo {
	var userInfo []UserInfo = make([]UserInfo, 1)

	return userInfo
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
