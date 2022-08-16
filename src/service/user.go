package service

import (
	"fmt"

	"gin01/src/DataBase"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

var jwtkey = []byte("Token")
var str string

type Claims struct {
	UserId     string
	ExpiryTime int
	jwt.StandardClaims
}

func TestUrl(c *gin.Context) {
	var form DataBase.User
	err := c.ShouldBindJSON(&form)
	if err != nil {
		return
	}
	uid := form.Select()
	c.JSON(200, gin.H{"uid": uid})
}
func Login(c *gin.Context) {
	var form DataBase.User
	err := c.ShouldBindJSON(&form)
	if err != nil {
		c.JSON(400, gin.H{"error": "缺少必要数据"})
		return
	}
	user := form.Select()
	if user == true {

		return
	}
	//tokenString := createToken(&user)
	c.JSON(200, gin.H{"code": 200})
	//c.String(http.StatusOK, "/index")
}

func createToken(form *DataBase.UserInfo) string {
	expireTime := time.Now().Add(time.Hour * 10)
	claims := Claims{
		UserId: form.Uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1", // 签名颁发者
			Subject:   "userToken", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	// fmt.Println(token)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println(err)
		return " "
	}
	return tokenString
}
func Getuserinfo(c *gin.Context) {
	tokenString := c.GetHeader("Token")
	if tokenString == "" {
		c.JSON(400, gin.H{"message": "缺少必要参数"})
		return
	}
	_, claims, errtoken := ParseToken(tokenString)
	if errtoken != nil {
		return
	}
	c.JSON(200, gin.H{"c": claims.UserId})
}
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}

func AddUser(c *gin.Context) {

}
func EditUser(c *gin.Context) {
	//var form DataBase.UserInfo
	//err := c.ShouldBindJSON(&form)
	//if err != nil {
	//
	//}
	//message := DataBase.EditUser(form)
	//c.JSON(200, gin.H{message: message})
}
func DelUser(c *gin.Context) {

}
func SelectUser(c *gin.Context) {

}
