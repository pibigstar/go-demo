package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

//  使用jwt 生成token 与使用

const (
	// 加密的key值
	secretKey = "pibigstar"
	// token有效期
	TokenClaimEXP = "exp"
	// token使用的范围
	TokenClaimScope = "web"
	TokenClaimAdmin = "admin"

	// 将用户uid存放到token中
	TokenClaimUID = "uid"
)

// 生成token
func GenJwtToken(claims jwt.MapClaims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}

// 检查token是否有效
func CheckJwtToken(tokenString string) bool {
	if tokenString == "" {
		return false
	}
	if err := CheckJwtTokenExpected(tokenString); err != nil {
		return false
	}
	return true
}

// 检查token是否过期
func CheckJwtTokenExpected(tokenString string) error {
	token, err := ParseJwtToken(tokenString)
	if err != nil {
		return err
	}
	return token.Claims.Valid()
}

// 解析token
func ParseJwtToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, errors.New("unexpected token claims")
		}
		return []byte(secretKey), nil
	})

	return token, err
}

// 从token中拿到uid
func GetUIDFromToken(tokenString string) (value interface{}, found bool) {
	token, err := ParseJwtToken(tokenString)
	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if v, ok := claims[TokenClaimUID]; ok {
			return v, true
		}
	}

	return nil, false
}
