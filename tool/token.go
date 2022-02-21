package tool

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type  MyClaims struct {
	Username string
	Type     string //"REFRESH_TOKEN"表示为一个refresh token，"TOKEN"表示为一个token
	Time     time.Time
	jwt.StandardClaims
}

var MySecret = []byte("皇帝爱吃饼干")

const TokenExpireDuration = time.Hour * 24

//颁布令牌
func GetToken (username, tokenType string, expireTime int64) (string ,error){

	c := MyClaims{
		Username: username,
		Type:     tokenType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + expireTime,
		},
	}
	//创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,c)

	return token.SignedString(MySecret)
}

//解析token
func ParseToken(tokenString string) (*MyClaims, error) {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil{
		fmt.Println("扯拐了")
		log.Fatal(err.Error())
		return nil, err
	}
	// 校验token
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		if claims.Type == "REFRESH_TOKEN" {
			errClaims := new(MyClaims)
			errClaims.Type = "ERR"
			return errClaims, nil
		}
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

//解析refresh—token
func ParseRefreshToken(tokenString string) (*MyClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})

	if clams, ok := token.Claims.(*MyClaims); ok && token.Valid {
		if clams.Type == "TOKEN" {
			errClaims := new(MyClaims)
			errClaims.Type = "ERR"
			return errClaims, nil
		}
		return clams, nil
	} else {
		return nil, err
	}
}

func CheckTokenErr(ctx *gin.Context, claims *MyClaims, err error) bool {
	if err == nil && claims.Type == "ERR" {
		Failed(ctx, "PARSE_TOKEN_ERROR")
		return false
	}

	if err != nil {
		fmt.Println("HAHA")
		if err.Error()[:16] == "token is expired" {
			Failed(ctx, "TOKEN_EXPIRED")
			return false
		}

		fmt.Println("getTokenParseTokenErr:", err)
		Failed(ctx, "PARSE_TOKEN_ERROR")
		return false
	}

	return true
}
