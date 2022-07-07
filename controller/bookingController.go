package controller

import (
	"Slot_booking/entity"
	"Slot_booking/service"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
)
import "github.com/gin-gonic/gin"

type BookingController interface {
	FindAll() []entity.Booking
	Save(ctx *gin.Context) error
	Cancel(ctx *gin.Context) error
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

func (c *Controller) Save(ctx *gin.Context) error {
	var booking entity.Booking

	jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		// Handle error
		return err
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(jsonData, &m)
	if err != nil {
		return err
	}
	booking.UserID = uint64(m["user_id"].(float64))
	//booking.SlotID = uint64(m["slot_id"].(float64))
	startTime := m["start_time"].(string)
	date := entity.DateForSlot()
	booking.Status = "booked"
	slot, err := c.slotService.Find(startTime, date)
	if err != nil {
		return err
	}
	booking.SlotID = slot.ID

	c.service.Save(booking)
	return nil
}
func (c *Controller) Cancel(ctx *gin.Context) error {
	var booking entity.Booking

	jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		// Handle error
		return err
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(jsonData, &m)
	if err != nil {
		return err
	}
	userID := uint64(m["user_id"].(float64))
	booking.UserID = uint64(m["user_id"].(float64))
	startTime := m["start_time"].(string)
	date := entity.DateForSlot()
	slot, err := c.slotService.Find(startTime, date)
	if err != nil {
		return err
	}

	_, err = c.userService.Find(userID)
	if err != nil {
		err = errors.New("Not a registered user")
		return err
	}

	slotTimeH, _ := strconv.Atoi(slot.StartTime[:2])
	slotTimeM, _ := strconv.Atoi(slot.StartTime[3:])
	presentTimeH, _ := strconv.Atoi(entity.PresentTimePlus30minutes()[:2])
	presentTimeM, _ := strconv.Atoi(entity.PresentTimePlus30minutes()[3:])

	if slotTimeH > presentTimeH {
		booking.Status = "cancelled"
	} else if slotTimeH == presentTimeH && slotTimeM >= presentTimeM {
		booking.Status = "cancelled"
	} else {
		err = errors.New("Crossed the cancellation time")
		return err
	}
	booking.SlotID = slot.ID
	err = c.service.Cancel(booking)
	return err
}

func (c *Controller) GetUserSlot(ctx *gin.Context) ([]entity.Slot, error) {
	UserID, err := strconv.ParseUint(ctx.Query("user_id"), 10, 64)

	if err != nil {
		return []entity.Slot{}, err
	}
	var bookedSlots []entity.Booking

	bookedSlots, err = c.service.GetUserBookings(UserID)

	if err != nil {
		fmt.Println("error", err.Error())
		return nil, err
	}
	var slotIDs []uint64
	for _, booking := range bookedSlots {
		slotIDs = append(slotIDs, booking.SlotID)
	}
	slots, err := c.slotService.GetSlots(slotIDs)
	if err != nil {
		fmt.Println("error", err.Error())
		return slots, err
	}
	return slots, err
}
