package service

import (
	"Slot_booking/entity"
	"Slot_booking/manager"
)

type UserService interface {
	AddUser(user entity.User) (entity.User, error)
	GetUser(userID uint64) (entity.User, error)
}

type userService struct {
	user manager.UserRepository
}

func NewUserService(repo manager.UserRepository) UserService {
	return &userService{
		user: repo,
	}
}

func (service *userService) AddUser(user entity.User) (entity.User, error) {
	err := service.user.Create(user)
	return user, err
}

func (service *userService) GetUser(userID uint64) (entity.User, error) {
	return service.user.Find(userID)
}

//
