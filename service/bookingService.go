package service

import (
	"Slot_booking/entity"
	"Slot_booking/manager"
)

type BookingService interface {
	CountAllBookedSlotsOfAUser(booking entity.Booking) (int64, error)
	CountTotalUsersBookingASlot(booking entity.Booking) (int64, error)
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

func (service *bookingService) CountAllBookedSlotsOfAUser(booking entity.Booking) (int64, error) {
	return service.bookings.CountAllBookedSlotsOfAUser(booking)
}

func (service *bookingService) CountTotalUsersBookingASlot(booking entity.Booking) (int64, error) {
	return service.bookings.CountTotalUsersBookingASlot(booking)
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
