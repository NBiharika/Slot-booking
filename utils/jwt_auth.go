package utils

import (
	"Slot_booking/entity"
	"fmt"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	userID uint64
	jwt.StandardClaims
}

func GenerateJWT(user entity.User) (tokenString string, err error) {
	fmt.Println("user", user)
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JWTClaim{
		userID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
func ValidateToken(signedToken string) (err error) {
	signedToken = strings.Split(signedToken, " ")[1]
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(
		signedToken,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	//claims, ok := token.Claims.(*JWTClaim)
	//if !ok {
	//	err = errors.New("couldn't parse claims")
	//}
	//if claims.ExpiresAt < time.Now().Local().Unix() {
	//	err = errors.New("token expired")
	//}
	fmt.Println("claims:", claims)
	return err
}

//func Decode()
