package VO

type UserInfo struct {
	Username   string `form:"username" json:"username"`
	Uid        string `form:"uid" json:"uid" `
	Age        string `form:"age" json:"age"`
	Sex        string `form:"sex" json:"sex"`
	UipCode    string `form:"uipcode" json:"uipCode"`
	RoleId     string `form:"roleId" json:"roleId"`
	CreateTime string `form:"createTime" json:"createTime"`
	UpdateTime string `form:"updateTime" json:"updateTime"`
}

type User struct {
	Username string `db:"username" form:"username" json:"username" binding:"required"`
	Password string `db:"password" form:"password" json:"password" binding:"required"`
}
