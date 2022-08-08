package service

import (
	"Slot_booking/entity"
	"Slot_booking/manager"
)

type UserService interface {
	AddUser(user entity.User) (entity.User, error)
	GetUser(userID uint64) (entity.User, error)
	FindUsingEmail(user entity.User) (entity.User, error)
	GetAllUsers() ([]entity.User, error)
	ChangeRoleToUser(user entity.User) (entity.User, error)
	ChangeRoleToAdmin(user entity.User) (entity.User, error)
	BlockUser(user entity.User) (entity.User, error)
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
	user, err := service.user.Create(user)
	return user, err
}

func (service *userService) GetUser(userID uint64) (entity.User, error) {
	return service.user.Find(userID)
}

func (service *userService) FindUsingEmail(user entity.User) (entity.User, error) {
	return service.user.FindUsingEmail(user)
}

func (service *userService) GetAllUsers() ([]entity.User, error) {
	return service.user.FindAll()
}

func (service *userService) ChangeRoleToUser(user entity.User) (entity.User, error) {
	return service.user.UpdateToUser(user)
}

func (service *userService) ChangeRoleToAdmin(user entity.User) (entity.User, error) {
	return service.user.UpdateToAdmin(user)
}

func (service *userService) BlockUser(user entity.User) (entity.User, error) {
	return service.user.UpdateToBlockUser(user)
}
