package helper

import (
	"cloud_disk3/core/define"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateToken(id int64, identity string, name string) (token string, err error) {
	//设置token为7天过期
	expireToken := time.Now().Add(time.Hour * 24 * 7).Unix()
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireToken,
		},
	}
	//加密
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	//盐值加密
	signedToken, err := t.SignedString([]byte(define.SECRET_KEY))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
