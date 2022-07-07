package controller

import (
	"Slot_booking/entity"
	"Slot_booking/service"
	"github.com/gin-gonic/gin"
)

type SlotController interface {
	FindAll() []entity.Slot
	Save(ctx *gin.Context) error
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
}

func (c *slotController) Save(ctx *gin.Context) error {
	slot := make([]entity.Slot, 24) //make
	for i := 0; i < 24; i++ {
		slot[i].Date = entity.DateForSlot()
		slot[i].StartTime = entity.StartTimeOfSlot(i)
		_, err := c.service.Save(slot[i])
		if err != nil {
			return err
		}
	}
	return nil
}
