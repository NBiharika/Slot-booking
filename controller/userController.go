package controller

import (
	"Slot_booking/cache"
	"Slot_booking/entity"
	"Slot_booking/service"
	"Slot_booking/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController interface {
	GetUser(ctx *gin.Context) (entity.User, error, int)
	AddUser(ctx *gin.Context) (error, int)
	GetAllUsers() ([]entity.User, error)
	SwitchRoles(ctx *gin.Context) error
	SwitchStatus(ctx *gin.Context) error
}

type userController struct {
	service   service.UserService
	userCache cache.UserCache
}

func NewUserController(service service.UserService, cache cache.UserCache) UserController {
	return &userController{
		service:   service,
		userCache: cache,
	}
}

func (c *userController) GetUser(ctx *gin.Context) (entity.User, error, int) {
	userReq := ctx.Value("user_info")
	jwtData := userReq.(*utils.JWTClaim)

	key := fmt.Sprintf("user_data_%v", jwtData.User.ID)
	user, err := c.userCache.GetUser(ctx, key)
	if err == nil {
		return user, err, http.StatusOK
	}
	user, err = c.service.GetUser(jwtData.User.ID)
	if err != nil {
		err = errors.New("invalid request")
		return entity.User{}, err, http.StatusBadRequest
	}
	if err != nil {
		return entity.User{}, err, http.StatusInternalServerError
	}
	if user.Status == "blocked" {
		err = errors.New("oops, you are blocked")
		return entity.User{}, err, http.StatusBadRequest
	}

	c.userCache.SetUser(ctx, key, user)
	return user, err, http.StatusOK
}

func (c *userController) AddUser(ctx *gin.Context) (error, int) {
	var user entity.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		return err, http.StatusBadRequest
	}
	if err := user.HashPassword(user.Password); err != nil {
		err = errors.New("password could not be created")
		return err, http.StatusInternalServerError
	}

	user, err := c.service.AddUser(user)
	if err != nil {
		err = errors.New("user already exists")
		return err, http.StatusInternalServerError
	}
	key := fmt.Sprintf("user_data_%v", user.ID)
	c.userCache.SetUser(ctx, key, user)
	return nil, http.StatusOK
}

func (c *userController) GetAllUsers() ([]entity.User, error) {
	return c.service.GetAllUsers()
}

func (c *userController) SwitchRoles(ctx *gin.Context) error {
	m, err := utils.ReadRequestBody(ctx)
	if err != nil {
		return err
	}
	email := m["email"].(string)
	role := m["role"].(string)
	user, err := c.service.SwitchRoles(email, role)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("user_data_%v", user.ID)
	err = c.userCache.RemoveCache(ctx, key)
	if err != nil {
		return err
	}
	return nil
}

func (c *userController) SwitchStatus(ctx *gin.Context) error {
	m, err := utils.ReadRequestBody(ctx)
	if err != nil {
		return err
	}
	email := m["email"].(string)
	status := m["status"].(string)
	user, err := c.service.SwitchStatus(email, status)
	if err != nil {
		err = errors.New("can't block the owner")
		return err
	}

	key := fmt.Sprintf("user_data_%v", user.ID)
	err = c.userCache.RemoveCache(ctx, key)
	if err != nil {
		return err
	}
	return nil
}
