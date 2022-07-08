package service

import (
	"Slot_booking/entity"
	"Slot_booking/manager"
)

type TokenService interface {
	GenerateToken(user entity.User) (entity.User, error)
}

type tokenService struct {
	user manager.UserRepository
}

func NewTokenService(repo manager.UserRepository) UserService {
	return &tokenService{
		token: repo,
	}
}

func (service *tokenService) GenerateToken(user entity.User) (entity.User, error) {
	return service.token.Find(user)
}
