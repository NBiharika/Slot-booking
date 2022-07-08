package controller

import (
	"Slot_booking/entity"
	"Slot_booking/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController interface {
	GetUser(ctx *gin.Context) (entity.User, error, int)
	AddUser(ctx *gin.Context) error
	RegisterUser(ctx *gin.Context) (entity.User, error, int)
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (c *userController) GetUser(ctx *gin.Context) (entity.User, error, int) {
	userID, err := strconv.ParseUint(ctx.Query("user_id"), 10, 64)
	if err != nil {
		err = errors.New("invalid request")
		return entity.User{}, err, http.StatusBadRequest
	}

	user, err := c.service.GetUser(userID)
	if err != nil {
		return entity.User{}, err, http.StatusInternalServerError
	}

	return user, err, http.StatusOK
}

func (c *userController) AddUser(ctx *gin.Context) error {
	var user entity.User
	err := ctx.BindJSON(&user)
	if err != nil {
		return err
	}
	_, err = c.service.AddUser(user)
	return err
}

func (c *userController) RegisterUser(ctx *gin.Context) (entity.User, error, int) {
	var user entity.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return entity.User{}, err, http.StatusBadRequest
	}
	if err := user.HashPassword(user.Password); err != nil {
		return user, err, http.StatusInternalServerError
	}
	user, err := c.service.AddUser(user)
	if err != nil {
		return user, err, http.StatusInternalServerError
	}
	return user, nil, http.StatusOK
}
