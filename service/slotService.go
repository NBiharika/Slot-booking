package service

import (
	"Slot_booking/entity"
	"Slot_booking/manager"
)

type SlotService interface {
	AddSlot(slot []entity.Slot) ([]entity.Slot, error)
	FindAll(startDate string, endDate string) []entity.Slot
	Find(startTime, date string) (entity.Slot, error)
	GetSlots(slotIDs []uint64) ([]entity.Slot, error)
	GetCount(date string) (int64, error)
}

type service struct {
	slots manager.SlotRepository
}

func NewSlotService(repo manager.SlotRepository) SlotService {
	return &service{
		slots: repo,
	}
}

func (service *service) AddSlot(slot []entity.Slot) ([]entity.Slot, error) {
	err := service.slots.Create(slot)
	return slot, err
}

func (service *service) FindAll(startDate string, endDate string) []entity.Slot {
	return service.slots.FindAll(startDate, endDate)
}

func (service *service) Find(startTime, date string) (entity.Slot, error) {
	var slot entity.Slot
	slot.StartTime = startTime
	slot.Date = date
	return service.slots.Find(slot)
}

func (service *service) GetSlots(slotIDs []uint64) ([]entity.Slot, error) {
	return service.slots.GetSlots(slotIDs)
}

func (service *service) GetCount(date string) (int64, error) {
	return service.slots.GetCount(date)
}
