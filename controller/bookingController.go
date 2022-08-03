package controller

import (
	"Slot_booking/cache"
	"Slot_booking/entity"
	"Slot_booking/service"
	"Slot_booking/utils"
	"errors"
	"fmt"
	"strconv"
	"time"
)
import "github.com/gin-gonic/gin"

const (
	limitForBookedSlotsOfAUserForADay = 5
	limitOfAllUsersBookingASlot       = 5
)

type BookingController interface {
	FindAll() []entity.Booking
	BookSlot(ctx *gin.Context) error
	CancelBooking(ctx *gin.Context) (string, error)
	GetUserSlot(ctx *gin.Context) ([]entity.Slot, error)
}

type Controller struct {
	service     service.BookingService
	slotService service.SlotService
	userService service.UserService
	userCache   cache.UserCache
}

func New(service service.BookingService, slotService service.SlotService, userService service.UserService, cache cache.UserCache) BookingController {
	return &Controller{
		service:     service,
		slotService: slotService,
		userService: userService,
		userCache:   cache,
	}
}

func (c *Controller) FindAll() []entity.Booking {
	return c.service.FindAll()
}

func (c *Controller) BookSlot(ctx *gin.Context) error {
	var booking entity.Booking

	userReq := ctx.Value("user_info")
	jwtData := userReq.(*utils.JWTClaim)

	key := fmt.Sprintf("user_data_%v", strconv.FormatUint(jwtData.User.ID, 10))
	user, err := c.userCache.GetUser(ctx, key)
	if err != nil {
		user, err = c.userService.GetUser(jwtData.User.ID)
		c.userCache.SetUser(ctx, key, user)
	}

	m, err := utils.ReadRequestBody(ctx)

	startTime := m["start_time"].(string)
	date := m["date"].(string)

	slot, err := c.slotService.Find(startTime, date)
	if err != nil {
		return err
	}

	dateStr := date + " " + slot.StartTime
	loc, _ := time.LoadLocation("Asia/Kolkata")
	slotDate, _ := time.ParseInLocation("2006-01-02 15:04", dateStr, loc)
	if slotDate.Before(time.Now()) {
		err = errors.New("crossed the booking time")
		return err
	}

	booking.UserID = user.ID
	booking.SlotID = slot.ID

	countAllBookedSlotsOfAUserForADay, err := c.service.CountAllBookedSlotsOfAUserForADay(booking, date)
	if err != nil {
		return err
	}
	if int(countAllBookedSlotsOfAUserForADay) >= limitForBookedSlotsOfAUserForADay {
		err = errors.New("a user can only book " + strconv.Itoa(limitForBookedSlotsOfAUserForADay) + " slots")
		return err
	}
	countTotalUsersBookingASlot, err := c.service.CountTotalUsersBookingASlot(booking)
	if err != nil {
		return err
	}
	if int(countTotalUsersBookingASlot) >= limitOfAllUsersBookingASlot {
		err = errors.New("a slot can only be booked by " + strconv.Itoa(limitOfAllUsersBookingASlot) + " users")
		return err
	}
	_, err = c.service.BookSlot(booking)
	booking.Status = "booked"
	return err
}

func (c *Controller) CancelBooking(ctx *gin.Context) (string, error) {
	var booking entity.Booking

	userReq := ctx.Value("user_info")
	jwtData := userReq.(*utils.JWTClaim)

	key := fmt.Sprintf("user_data_%v", strconv.FormatUint(jwtData.User.ID, 10))
	user, err := c.userCache.GetUser(ctx, key)
	if err != nil {
		user, err = c.userService.GetUser(jwtData.User.ID)
		c.userCache.SetUser(ctx, key, user)
	}

	m, err := utils.ReadRequestBody(ctx)

	startTime := m["start_time"].(string)
	date := m["date"].(string)

	slot, err := c.slotService.Find(startTime, date)
	if err != nil {
		return "", err
	}

	_, err = c.userService.GetUser(user.ID)
	if err != nil {
		err = errors.New("not a registered user")
		return "", err
	}

	todayTimePlus30Minutes := time.Now().Add(30 * time.Minute)
	dateStr := date + " " + slot.StartTime
	loc, _ := time.LoadLocation("Asia/Kolkata")
	slotDate, _ := time.ParseInLocation("2006-01-02 15:04", dateStr, loc)

	if slotDate.Before(todayTimePlus30Minutes) {
		err = errors.New("crossed the cancellation time")
		return "", err
	}
	booking.Status = "cancelled"
	booking.UserID = user.ID
	booking.SlotID = slot.ID
	rowsAffected, err := c.service.CancelBooking(booking)
	if err != nil {
		return "", err
	}
	message := "slot has been cancelled"
	if rowsAffected == 0 {
		message = "slot has already been cancelled"
	}
	return message, err
}

func (c *Controller) GetUserSlot(ctx *gin.Context) ([]entity.Slot, error) {
	userReq := ctx.Value("user_info")
	jwtData := userReq.(*utils.JWTClaim)

	user, err := c.userService.GetUser(jwtData.User.ID)
	if err != nil {
		return []entity.Slot{}, err
	}

	var bookedSlots []entity.Booking
	bookedSlots, err = c.service.GetUserBookings(user.ID)
	if err != nil {
		return nil, err
	}

	var slotIDs []uint64
	for _, booking := range bookedSlots {
		slotIDs = append(slotIDs, booking.SlotID)
	}
	slots, err := c.slotService.GetSlots(slotIDs)
	if err != nil {
		return slots, err
	}
	return slots, err
}
