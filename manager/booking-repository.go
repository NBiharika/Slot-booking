package manager

import (
	"Slot_booking/entity"
	"fmt"
	"gorm.io/gorm"
)

type BookingRepository interface {
	Save(booking entity.Booking) error
	FindAll() []entity.Booking
	Cancel(booking entity.Booking) error
	GetUserBookings(userID uint64) ([]entity.Booking, error)
}

type BookingDB struct {
	connection *gorm.DB
}

func BookingRepo() BookingRepository {
	return &BookingDB{
		connection: dbClient,
	}
}

func (db *BookingDB) Save(booking entity.Booking) error {
	//db.connection.AutoMigrate(&entity.Booking{})
	err := db.connection.Create(&booking).Error
	return err
}

func (db *BookingDB) Cancel(booking entity.Booking) error {
	err := db.connection.Model(&entity.Booking{}).Where("user_id=? and slot_id=?", booking.UserID, booking.SlotID).Update("status", booking.Status).Error
	return err
}

func (db *BookingDB) GetUserBookings(userID uint64) ([]entity.Booking, error) {
	var booked []entity.Booking
	err := db.connection.Model(&entity.Booking{}).Debug().Where("user_id=? and status=?", userID, "booked").Find(&booked).Error
	fmt.Println("check:", err)
	return booked, err
}

func (db *BookingDB) FindAll() []entity.Booking {
	var booked []entity.Booking
	db.connection.Find(&booked)
	return booked
}
