package entity

import "time"

type Booking struct {
	User      User      `json:"-" gorm:"foreignKey:UserID"`
	Slot      Slot      `json:"-" gorm:"foreignKey:SlotID"`
	UserID    uint64    `gorm:"primaryKey;uniqueIndex:idx_name"`
	SlotID    uint64    `gorm:"primaryKey;uniqueIndex:idx_name"`
	Status    string    `json:"status"`
	CreatedOn time.Time `json:"created_on" gorm:"autoUpdateTime:milli"`
	UpdatedOn time.Time `json:"updated_on" gorm:"autoUpdateTime:nano"`
}

func (Booking) TableName() string {
	return "bookings"
}
