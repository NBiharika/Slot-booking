package utils

import (
	"Slot_booking/entity"
	"errors"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	User entity.User `json:"user"`
	jwt.StandardClaims
}

func GenerateJWT(user entity.User) (tokenString string, err error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JWTClaim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(signedToken string) (claims *JWTClaim, err error) {
	signedToken = strings.Split(signedToken, " ")[1]
	claims = &JWTClaim{}
	_, err = jwt.ParseWithClaims(
		signedToken,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)
	if err != nil {
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
	}

	return claims, err
}
