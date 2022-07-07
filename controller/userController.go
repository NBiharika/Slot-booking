package controller

import (
	"Slot_booking/entity"
	"Slot_booking/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController interface {
	Find(ctx *gin.Context) (entity.User, error)
	Save(ctx *gin.Context) error
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}
func (c *userController) Find(ctx *gin.Context) (entity.User, error) {
	userID, err := strconv.ParseUint(ctx.Query("user_id"), 10, 64)
	if err != nil {
		return entity.User{}, err
	}
	return c.service.Find(userID)
}

func (c *userController) Save(ctx *gin.Context) error {
	var user entity.User
	err := ctx.BindJSON(&user)
	if err != nil {
		return err
	}
	fmt.Println(user)
	_, err = c.service.Save(user)
	return err
}
