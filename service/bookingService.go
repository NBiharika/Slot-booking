package service

import (
	"Slot_booking/entity"
	"Slot_booking/manager"
)

type BookingService interface {
	BookSlot(booking entity.Booking) (entity.Booking, error)
	FindAll() []entity.Booking
	CancelBooking(booking entity.Booking) (int64, error)
	GetUserBookings(userID uint64) ([]entity.Booking, error)
}

type bookingService struct {
	bookings manager.BookingRepository
}

func NewService(repo manager.BookingRepository) BookingService {
	return &bookingService{
		bookings: repo,
	}
}

func (service *bookingService) BookSlot(booking entity.Booking) (entity.Booking, error) {
	err := service.bookings.Create(booking)
	return booking, err
}

func (service *bookingService) CancelBooking(booking entity.Booking) (int64, error) {
	return service.bookings.Cancel(booking)
}

func (service *bookingService) FindAll() []entity.Booking {
	return service.bookings.FindAll()
}

func (service *bookingService) GetUserBookings(userID uint64) ([]entity.Booking, error) {
	return service.bookings.GetUserBookings(userID)
}
