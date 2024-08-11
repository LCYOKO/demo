package goin

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"testing"
	"time"
)

var secret = []byte("xiaomi")
var secretString = "xiaomi"

type UserInfo struct {
	Id   int64
	Name string
	*jwt.RegisteredClaims
}

func TestGenerateToken(t *testing.T) {
	// 创建 Claims
	claims := &UserInfo{
		1,
		"lisi",
		&jwt.RegisteredClaims{

			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 过期时间
			Issuer:    "liuchiyun",                                        // 签发人
		}}

	// 生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(secret)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(signedString)
}

func TestDecode(t *testing.T) {
	jwtToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiTmFtZSI6Imxpc2kiLCJpc3MiOiJsaXVjaGl5dW4iLCJleHAiOjE3MjMyOTk2NDF9.8DdSZMuNHGG48ND31Vq1oLBG06MkvVES-5ulOiHIEO8"
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	var u UserInfo
	token, err := jwt.ParseWithClaims(jwtToken, &u, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return secret, nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u)
	fmt.Println(token)
}
