package VO

type UserInfo struct {
	UserName   string `form:"username" json:"username"`
	Uid        string `form:"uid" json:"uid" `
	Age        string `form:"age" json:"age"`
	Sex        string `form:"sex" json:"sex"`
	UipCode    string `form:"uipcode" json:"uipCode"`
	RoleId     string `form:"roleId" json:"roleId"`
	CreateTime string `form:"createTime" json:"createTime"`
	UpdateTime string `form:"updateTime" json:"updateTime"`
}

type User struct {
	Username string `db:"userName" form:"username"  binding:"required"`
	Password string `db:"passWord" form:"password"  binding:"required"`
}
type Customer struct {
	ipCode   string
	inTime   string
	outTime  string
	openTime string
}
