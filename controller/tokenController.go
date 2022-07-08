package controller

import (
	"Slot_booking/entity"
	"Slot_booking/service"
	"Slot_booking/utils"
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

	// check if email exists and password is correct
	if requestErr := c.service.FindUsingEmail(user); requestErr != nil {
		return "", requestErr, http.StatusInternalServerError
	}

	if credentialError := user.CheckPassword(request.Password); credentialError != nil {
		return "", credentialError, http.StatusUnauthorized
	}

	tokenString, err := utils.GenerateJWT(user.Email)
	if err != nil {
		return "", err, http.StatusInternalServerError
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
	return tokenString, nil, http.StatusOK
}
