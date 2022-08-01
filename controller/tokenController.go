package controller

import (
	"Slot_booking/entity"
	"Slot_booking/service"
	"Slot_booking/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TokenController interface {
	GenerateToken(context *gin.Context) (string, error, int)
}

type tokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	service  service.UserService
}

func NewTokenController(service service.UserService) TokenController {
	return &tokenRequest{
		service: service,
	}
}

func (c *tokenRequest) GenerateToken(ctx *gin.Context) (string, error, int) {
	var userRequest tokenRequest
	err := ctx.ShouldBindJSON(&userRequest)
	if err != nil {
		return "", err, http.StatusBadRequest
	}

	user, err := c.service.FindUsingEmail(entity.User{Email: userRequest.Email})
	if err != nil {
		err = errors.New("wrong email id")
		return "", err, http.StatusInternalServerError
	}

	err = user.CheckPassword(userRequest.Password)
	if err != nil {
		err = errors.New("invalid credentials")
		return "", err, http.StatusUnauthorized
	}
	tokenString, err := utils.GenerateJWT(user)
	if err != nil {
		return "", err, http.StatusInternalServerError
	}
	return tokenString, nil, http.StatusOK
}
