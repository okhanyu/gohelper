package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

const TokenExpireDuration = time.Hour * 2

// const TokenExpireDuration = time.Second * 60
// var Secret = []byte("demo")

type MyClaims struct {
	Username string
	jwt.StandardClaims
}

// GenerateToken 生成Token
func GenerateToken(username string, secret []byte, issuer string) (string, error) {
	claims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    issuer,                                     // 签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 进行签名生成对应的token
	return token.SignedString(secret)
}

// ParseToken 解析Token
func ParseToken(tokenString string, secret []byte) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
