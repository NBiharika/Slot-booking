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

func (c *tokenRequest) GenerateToken(context *gin.Context) (string, error, int) {
	var request tokenRequest
	var user entity.User

	if err := context.ShouldBindJSON(&request); err != nil {
		return "", err, http.StatusBadRequest
	}

	if requestErr := c.service.FindUsingEmail(user); requestErr != nil {
		requestErr = errors.New("wrong email id or password")
		return "", requestErr, http.StatusInternalServerError
	}

	if credentialError := user.CheckPassword(request.Password); credentialError != nil {
		//credentialError = errors.New("")
		return "", credentialError, http.StatusUnauthorized
	}

	tokenString, err := utils.GenerateJWT(user.Email)
	if err != nil {
		return "", err, http.StatusInternalServerError
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
	return tokenString, nil, http.StatusOK
}
