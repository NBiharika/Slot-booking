package controller

import (
	"Slot_booking/cache"
	"Slot_booking/entity"
	"Slot_booking/service"
	"Slot_booking/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type SlotController interface {
	FindAll(ctx *gin.Context, startDate time.Time, endTime time.Time) []entity.Slot
	AddSlot(ctx *gin.Context) error
}

type slotController struct {
	service   service.SlotService
	slotCache cache.SlotCache
}

func NewSlotController(service service.SlotService, cache cache.SlotCache) SlotController {
	return &slotController{
		service:   service,
		slotCache: cache,
	}
}

func (c *slotController) FindAll(ctx *gin.Context, todayTime time.Time, endTime time.Time) []entity.Slot {
	finalSlots := make([]entity.Slot, 0)
	dates := make([]string, 0)

	for date := todayTime; date.After(endTime) == false; date = date.Add(24 * time.Hour) {
		formatDate := date.Format("2006-01-02")

		key := fmt.Sprintf("slots_%v", formatDate)
		slots, err := c.slotCache.GetSlot(ctx, key)
		if err == nil {
			for _, slot := range slots {
				finalSlots = append(finalSlots, slot)
			}
			return finalSlots
		}
		dates = append(dates, formatDate)
	}

	slots, _ := c.service.FindAll(dates)
	for _, slot := range slots {
		finalSlots = append(finalSlots, slot)
	}
	date := slots[0].Date
	key := fmt.Sprintf("slots_%v", date)
	c.slotCache.SetSlot(ctx, key, slots)
	return finalSlots
}

func (c *slotController) AddSlot(ctx *gin.Context) error {
	m, err := utils.ReadRequestBody(ctx)
	date := m["date"].(string)
	//todayTime := time.Now()
	//NextTime := todayTime.Add(6 * 24 * time.Hour)
	//date := NextTime.Format("2006-01-02")
	count, err := c.service.GetCount(date)
	if err != nil {
		return err
	}
	if count == 24 {
		err = errors.New("slots are already added")
		return err
	}

	slot := make([]entity.Slot, 24)

	for i := range [24]int{} {
		slot[i].Date = date
		slot[i].StartTime = entity.StartTimeOfSlot(i)
	}

	slot, err = c.service.AddSlot(slot)
	if err != nil {
		return err
	}
	fmt.Println(date)
	key := fmt.Sprintf("slots_%v", slot[0].Date)
	c.slotCache.SetSlot(ctx, key, slot)
	return nil
}
