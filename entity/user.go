package entity

import "time"

type User struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment"`
	FirstName string    `json:"firstName" binding:"required" gorm:"type:varchar(32)"`
	LastName  string    `json:"lastName" binding:"required" gorm:"type:varchar(32)"`
	Email     string    `json:"email" binding:"required,email" gorm:"type:varchar(256);unique"`
	CreatedOn time.Time `json:"created_on" gorm:"autoUpdateTime:milli"`
	UpdatedOn time.Time `json:"updated_on" gorm:"autoUpdateTime:nano"`
}
