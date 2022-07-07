package service

import (
	"Slot_booking/entity"
	"Slot_booking/manager"
)

type UserService interface {
	Save(user entity.User) (entity.User, error)
	Find(userID uint64) (entity.User, error)
}

type userService struct {
	user manager.UserRepository
}

func NewUserService(repo manager.UserRepository) UserService {
	return &userService{
		user: repo,
	}
}

func (service *userService) Save(user entity.User) (entity.User, error) {
	err := service.user.Save(user)
	return user, err
}

func (service *userService) Find(userID uint64) (entity.User, error) {
	return service.user.Find(userID)
}
