package routes

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type user struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
type usertoken struct {
	Token string `form:"token" json:"token" binding:"required"`
}

var jwtkey = []byte("Token")
var str string

type Claims struct {
	UserId     string
	ExpiryTime int
	jwt.StandardClaims
}

func LoginRoute(e *gin.Engine) {
	e.POST("/login", login)
	e.POST("/getuserinfo", getuserinfo)
	e.POST("/add", add)
}
func add(c *gin.Context) {

	c.JSON(200, gin.H{"code": 300})
}
func login(c *gin.Context) {
	var form user
	err := c.ShouldBindJSON(&form)
	if err != nil {
		c.JSON(400, gin.H{"error": "缺少必要数据"})
		return
	}
	fmt.Println(form)
	tokenString := createToken(&form)
	c.JSON(200, gin.H{"code": 200, "token": tokenString})
	//c.String(http.StatusOK, "/index")
}

func createToken(form *user) string {
	expireTime := time.Now().Add(time.Second * 10)
	claims := &Claims{
		UserId: form.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1", // 签名颁发者
			Subject:   "userToken", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println(token)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println(err)
	}
	str = tokenString
	return tokenString
}
func getuserinfo(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(400, gin.H{"message": "缺少必要参数"})
		return
	}
	strtoken, claims, errtoken := ParseToken(tokenString)
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
