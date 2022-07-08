package controller

import (
	"Slot_booking/entity"
	"Slot_booking/service"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController interface {
	Find(ctx *gin.Context) (entity.User, error, int)
	AddUser(ctx *gin.Context) error
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}
func (c *userController) Find(ctx *gin.Context) (entity.User, error, int) {
	userID, err := strconv.ParseUint(ctx.Query("user_id"), 10, 64)
	if err != nil {
		err = errors.New("invalid request")
		return entity.User{}, err, http.StatusBadRequest
	}

	user, err := c.service.Find(userID)
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
	fmt.Println(user)
	_, err = c.service.AddUser(user)
	return err
}
