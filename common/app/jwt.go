package app

import (
	"time"

	"Snai.CMS.Api/common/config"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

func GetJwtSecret() []byte {
	return []byte(config.AppConf.JwtSecret)
}

func GenerateToken(userName string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(config.AppConf.JwtExpire)
	claims := Claims{
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.AppConf.JwtIssuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJwtSecret())
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJwtSecret(), nil
	})

	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok & tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
