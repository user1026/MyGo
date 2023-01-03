package model

type UserInfo struct {
	UserName   string `form:"userName" json:"userName" gorm:"column:UserName"`
	Password   string `form:"passWord" form:"password"  gorm:"column:PassWord"`
	Uid        string `form:"uid" json:"uid" gorm:"column:Uid"`
	Age        string `form:"age" json:"age" gorm:"column:Age"`
	Sex        string `form:"sex" json:"sex" gorm:"column:Sex"`
	PhoneNum   string `form:"phoneNum" json:"phoneNum" gorm:"column:PhoneNum"`
	UipCode    string `form:"uipCode" json:"uipCode" gorm:"column:UipCode"`
	RoleId     string `form:"roleId" json:"roleId" gorm:"column:RoleId"`
	CreateTime string `form:"createTime" json:"createTime" gorm:"column:CreateTime"`
	UpdateTime string `form:"updateTime" json:"updateTime" gorm:"column:UpdateTime"`
	Token      string `json:"token"  gorm:"column:Token"`
}

type User struct {
	Username string `db:"userName" form:"username"  binding:"required" gorm:"column:Username"`
	Password string `db:"passWord" form:"password"  binding:"required" gorm:"column:Password"`
}
type Customer struct {
	ipCode   string
	inTime   string
	outTime  string
	openTime string
}
