package controller

import (
	"Slot_booking/entity"
	"Slot_booking/service"
	"Slot_booking/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController interface {
	GetUser(ctx *gin.Context) (entity.User, error, int)
	AddUser(ctx *gin.Context) (error, int)
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
	userReq := ctx.Value("user_info")
	jwtData := userReq.(*utils.JWTClaim)

	user, err := c.service.GetUser(jwtData.User.ID)
	if err != nil {
		err = errors.New("invalid request")
		return entity.User{}, err, http.StatusBadRequest
	}
	if err != nil {
		return entity.User{}, err, http.StatusInternalServerError
	}

	return user, err, http.StatusOK
}

func (c *userController) AddUser(ctx *gin.Context) (error, int) {
	var user entity.User

	user.FirstName = ctx.PostForm("first_name")
	user.LastName = ctx.PostForm("last_name")
	user.Email = ctx.PostForm("email")
	user.Password = ctx.PostForm("password")

	//if err := ctx.ShouldBindJSON(&user); err != nil {
	//	err = errors.New("invalid request")
	//	//d, _ := utils.ReadRequestBody(ctx)
	//
	//	fmt.Println("getsomething", d)
	//	return err, http.StatusBadRequest
	//
	//}
	if err := user.HashPassword(user.Password); err != nil {
		err = errors.New("password could not be created")
		return err, http.StatusInternalServerError
	}
	_, err := c.service.AddUser(user)
	if err != nil {
		err = errors.New("user already exists")
		return err, http.StatusInternalServerError
	}
	return nil, http.StatusOK
}
