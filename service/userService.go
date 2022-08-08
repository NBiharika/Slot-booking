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
	ChangeRoleToUser(email string) error
	ChangeRoleToAdmin(email string) error
	BlockUser(email string) error
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

func (service *userService) ChangeRoleToUser(email string) error {
	return service.user.UpdateToUser(email)
}

func (service *userService) ChangeRoleToAdmin(email string) error {
	return service.user.UpdateToAdmin(email)
}

func (service *userService) BlockUser(email string) error {
	return service.user.UpdateToBlockUser(email)
}
