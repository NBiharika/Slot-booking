package controller

import (
	"Slot_booking/entity"
	"Slot_booking/service"
	"errors"
	"github.com/gin-gonic/gin"
)

type SlotController interface {
	FindAll() []entity.Slot
	AddSlot(ctx *gin.Context) error
}

type slotController struct {
	service service.SlotService
}

func NewSlotController(service service.SlotService) SlotController {
	return &slotController{
		service: service,
	}
}

func (c *slotController) FindAll() []entity.Slot {
	return c.service.FindAll()
	//argument - startdate,enddate - string
}

func (c *slotController) AddSlot(ctx *gin.Context) error {

	date := entity.DateForSlot()
	count, err := c.service.GetCount(date)
	if err != nil {
		return err
	}
	if count == 24 {
		err = errors.New("today's slots are already added")
		return err
	}

	slot := make([]entity.Slot, 24)

	for i := 0; i < 24; i++ {
		slot[i].Date = entity.DateForSlot()
		slot[i].StartTime = entity.StartTimeOfSlot(i)
	}

	_, err = c.service.AddSlot(slot)
	if err != nil {
		return err
	}
	return nil
}
