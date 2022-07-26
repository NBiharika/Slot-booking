package service

import (
	"Slot_booking/entity"
	"Slot_booking/manager"
)

type BookingService interface {
	CountSlotsForAUser(booking entity.Booking) (int64, error)
	CountUsersForASlot(booking entity.Booking) (int64, error)
	BookSlot(booking entity.Booking) (int64, error)
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

func (service *bookingService) CountSlotsForAUser(booking entity.Booking) (int64, error) {
	return service.bookings.CountSlotsForAUser(booking)
}

func (service *bookingService) CountUsersForASlot(booking entity.Booking) (int64, error) {
	return service.bookings.CountUsersForASlot(booking)
}

func (service *bookingService) BookSlot(booking entity.Booking) (int64, error) {
	return service.bookings.Create(booking)
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
