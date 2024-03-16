package helper

import (
	"cloud_disk3/core/define"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(id int64, identity string, name string) (token string, err error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	token, err = claims.SignedString([]byte(define.SECRET_KEY))
	if err != nil {
		return "", err
	}

	return token, nil
}
