package manager

import (
	"Slot_booking/entity"
	"gorm.io/gorm"
)

type BookingRepository interface {
	Create(booking entity.Booking) error
	FindAll() []entity.Booking
	Cancel(booking entity.Booking) (int64, error)
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

func (db *BookingDB) Create(booking entity.Booking) error {
	//db.connection.AutoMigrate(&entity.Booking{})
	err := db.connection.Model(&entity.Booking{}).Where("user_id=? and slot_id=? and status=?", booking.UserID, booking.SlotID, "cancelled").Find(&booking).Error
	if err == nil {
		booking.Status = "booked"
		db.connection.Model(&entity.Booking{}).Where("user_id=? and slot_id=?", booking.UserID, booking.SlotID).Update("status", booking.Status)
		return nil
	} else {
		err = db.connection.Create(&booking).Error
	}
	return err
}

func (db *BookingDB) Cancel(booking entity.Booking) (int64, error) {
	resp := db.connection.Model(&entity.Booking{}).Where("user_id=? and slot_id=?", booking.UserID, booking.SlotID).Update("status", booking.Status)
	return resp.RowsAffected, resp.Error
}

func (db *BookingDB) GetUserBookings(userID uint64) ([]entity.Booking, error) {
	var booked []entity.Booking
	err := db.connection.Model(&entity.Booking{}).Where("user_id=? and status=?", userID, "booked").Find(&booked).Error
	return booked, err
}

func (db *BookingDB) FindAll() []entity.Booking {
	var booked []entity.Booking
	db.connection.Find(&booked)
	return booked
}
