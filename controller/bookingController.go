package controller

import (
	"Slot_booking/entity"
	"Slot_booking/service"
	"Slot_booking/utils"
	"errors"
	"strconv"
	"time"
)
import "github.com/gin-gonic/gin"

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
}

func New(service service.BookingService, slotService service.SlotService, userService service.UserService) BookingController {
	return &Controller{
		service:     service,
		slotService: slotService,
		userService: userService,
	}
}

func (c *Controller) FindAll() []entity.Booking {
	return c.service.FindAll()
}

func (c *Controller) BookSlot(ctx *gin.Context) error {
	var booking entity.Booking

	userReq := ctx.Value("user_info")
	jwtData := userReq.(*utils.JWTClaim)
	user, err := c.userService.GetUser(jwtData.User.ID)

	m, err := utils.ReadRequestBody(ctx)

	startTime := m["start_time"].(string)
	date := m["date"].(string)

	slot, err := c.slotService.Find(startTime, date)
	if err != nil {
		return err
	}

	slotTimeH, _ := strconv.Atoi(slot.StartTime[:2])
	slotTimeM, _ := strconv.Atoi(slot.StartTime[3:])
	presentTimeH, _ := strconv.Atoi(entity.PresentTime()[:2])
	presentTimeM, _ := strconv.Atoi(entity.PresentTime()[3:])
	todayDate := entity.DateForSlot(time.Now())
	todayDateMonth, _ := strconv.Atoi(todayDate[5:7])
	dateMonth, _ := strconv.Atoi(date[5:7])
	todayDateDay, _ := strconv.Atoi(todayDate[8:])
	dateDay, _ := strconv.Atoi(date[8:])

	if (todayDateMonth == dateMonth && todayDateDay == dateDay) && (slotTimeH < presentTimeH || (slotTimeH == presentTimeH && slotTimeM < presentTimeM)) {
		err = errors.New("crossed the booking time")
		return err
	}

	booking.UserID = user.ID
	booking.SlotID = slot.ID

	_, err = c.service.BookSlot(booking)
	booking.Status = "booked"
	return err
}
func (c *Controller) CancelBooking(ctx *gin.Context) (string, error) {
	var booking entity.Booking

	userReq := ctx.Value("user_info")
	jwtData := userReq.(*utils.JWTClaim)
	user, err := c.userService.GetUser(jwtData.User.ID)
	booking.UserID = user.ID

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

	slotTimeH, _ := strconv.Atoi(slot.StartTime[:2])
	slotTimeM, _ := strconv.Atoi(slot.StartTime[3:])
	presentTimeH, _ := strconv.Atoi(entity.PresentTimePlus30minutes()[:2])
	presentTimeM, _ := strconv.Atoi(entity.PresentTimePlus30minutes()[3:])
	todayDate := entity.DateForSlot(time.Now())
	todayDateMonth, _ := strconv.Atoi(todayDate[5:7])
	dateMonth, _ := strconv.Atoi(date[5:7])
	todayDateDay, _ := strconv.Atoi(todayDate[8:])
	dateDay, _ := strconv.Atoi(date[8:])

	if (todayDateMonth == dateMonth && todayDateDay == dateDay) && (slotTimeH < presentTimeH || (slotTimeH == presentTimeH && slotTimeM < presentTimeM)) {
		err = errors.New("crossed the cancellation time")
		return "", err
	}
	booking.Status = "cancelled"
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
