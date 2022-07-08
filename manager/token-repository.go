package manager

import (
	"Slot_booking/entity"
	"gorm.io/gorm"
)

type TokenRepository interface {
	Find(userID uint64) error
}

type UserDB struct {
	connection *gorm.DB
}

func UserRepo() TokenRepository {
	return &UserDB{
		connection: dbClient,
	}
}

func (db *UserDB) Find(user entity.User) error {
	var user entity.User
	user.ID = userID
	err := db.connection.Where("email = ?", request.Email).First(&user).Error
	return err
}
