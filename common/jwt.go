package common

import (
	"essential/model"
	"github.com/golang-jwt/jwt"
	"time"
)

var jwtKey = []byte("weiyi")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

//iss (issuer)：签发人
//exp (expiration time)：过期时间
//sub (subject)：主题
//aud (audience)：受众
//nbf (Not Before)：生效时间
//iat (Issued At)：签发时间
//jti (JWT ID)：编号

func ReleaseToke(user model.Users) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "liuyang",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParesToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
