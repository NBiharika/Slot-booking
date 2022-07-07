package service

import (
	"Slot_booking/entity"
	"Slot_booking/manager"
)

type BookingService interface {
	Save(booking entity.Booking) (entity.Booking, error)
	FindAll() []entity.Booking
	Cancel(booking entity.Booking) error
	GetUserBookings(userID uint64) ([]entity.Booking, error)
}

type bookingservice struct {
	bookings manager.BookingRepository
}

func NewService(repo manager.BookingRepository) BookingService {
	return &bookingservice{
		bookings: repo,
	}
}

func (service *bookingservice) Save(booking entity.Booking) (entity.Booking, error) {
	err := service.bookings.Save(booking)
	return booking, err
}

func (service *bookingservice) Cancel(booking entity.Booking) error {
	err := service.bookings.Cancel(booking)
	return err
}

func (service *bookingservice) FindAll() []entity.Booking {
	return service.bookings.FindAll()
}
func (service *bookingservice) GetUserBookings(userID uint64) ([]entity.Booking, error) {
	return service.bookings.GetUserBookings(userID)
}
