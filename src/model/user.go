package model

type UserInfo struct {
	UserName   string `json:"userName" gorm:"column:UserName"`
	Password   string `form:"password"  gorm:"column:PassWord"`
	Uid        string `json:"uid" gorm:"column:Uid"`
	Age        string ` json:"age" gorm:"column:Age"`
	Sex        string ` json:"sex" gorm:"column:Sex"`
	PhoneNum   string ` json:"phoneNum" gorm:"column:PhoneNum"`
	UipCode    string ` json:"uipCode" gorm:"column:UipCode"`
	RoleId     string ` json:"roleId" gorm:"column:RoleId"`
	CreateTime string ` json:"createTime" gorm:"column:CreateTime"`
	UpdateTime string ` json:"updateTime" gorm:"column:UpdateTime"`
	Token      string `json:"token"  gorm:"column:Token"`
}
type EditUserInfo struct {
	UserName   string `json:"userName" gorm:"column:UserName" binding:"required"`
	Password   string `form:"password"  gorm:"column:PassWord" binding:"required"`
	Uid        string `json:"uid" gorm:"column:Uid"  binding:"required"`
	Age        string ` json:"age" gorm:"column:Age" binding:"required"`
	Sex        string ` json:"sex" gorm:"column:Sex" binding:"required"`
	PhoneNum   string ` json:"phoneNum" gorm:"column:PhoneNum" binding:"required"`
	UipCode    string ` json:"uipCode" gorm:"column:UipCode" binding:"required"`
	RoleId     string ` json:"roleId" gorm:"column:RoleId" binding:"required"`
	UpdateTime string ` json:"updateTime" gorm:"column:UpdateTime" `
}
type Login struct {
	Username string `db:"userName" form:"username"  binding:"required" gorm:"column:Username"`
	Password string `db:"passWord" form:"password"  binding:"required" gorm:"column:Password"`
}
type Customer struct {
	ipCode   string
	inTime   string
	outTime  string
	openTime string
}
