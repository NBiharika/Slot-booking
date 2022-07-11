package entity

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment"`
	FirstName string    `json:"firstName" binding:"required" gorm:"type:varchar(32)"`
	LastName  string    `json:"lastName" binding:"required" gorm:"type:varchar(32)"`
	Email     string    `json:"email" binding:"required,email" gorm:"type:varchar(256);unique"`
	Password  string    `json:"password" binding:"required"`
	CreatedOn time.Time `json:"created_on" gorm:"autoUpdateTime:milli"`
	UpdatedOn time.Time `json:"updated_on" gorm:"autoUpdateTime:nano"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
