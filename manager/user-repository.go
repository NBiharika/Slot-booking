package manager

import (
	"Slot_booking/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	Create(user entity.User) (entity.User, error)
	Find(userID uint64) (entity.User, error)
	FindUsingEmail(user entity.User) (entity.User, error)
	FindAll() ([]entity.User, error)
	UpdateToSwitchRoles(email string, role string) (entity.User, error)
	UpdateToSwitchStatus(email string, status string) (entity.User, error)
}

type UserDB struct {
	connection *gorm.DB
}

func UserRepo() UserRepository {
	return &UserDB{
		connection: dbClient,
	}
}

func (db *UserDB) Create(user entity.User) (entity.User, error) {
	err := db.connection.Create(&user).Error
	return user, err
}

func (db *UserDB) Find(userID uint64) (entity.User, error) {
	var user entity.User
	user.ID = userID
	err := db.connection.First(&user).Error

	return user, err
}

func (db *UserDB) FindUsingEmail(user entity.User) (entity.User, error) {
	err := db.connection.Where("email = ?", user.Email).First(&user).Error
	return user, err
}

func (db *UserDB) FindAll() ([]entity.User, error) {
	var users []entity.User
	err := db.connection.Find(&users).Error
	return users, err
}

func (db *UserDB) UpdateToSwitchRoles(email string, role string) (entity.User, error) {
	var user entity.User
	err := db.connection.Model(&user).Clauses(clause.Returning{}).Where("email=? and status=? and role!=?", email, "active", "owner").Update("role", role).Error
	return user, err
}

func (db *UserDB) UpdateToSwitchStatus(email string, status string) (entity.User, error) {
	var user entity.User
	err := db.connection.Model(&user).Clauses(clause.Returning{}).Where("email=? and role!=?", email, "owner").Update("status", status).Error
	return user, err
}
