package service

import (
	"Slot_booking/entity"
	"Slot_booking/manager"
)

type SlotService interface {
	Save(booking entity.Slot) (entity.Slot, error)
	FindAll() []entity.Slot
	Find(startTime, date string) (entity.Slot, error)
	GetSlots(slotIDs []uint64) ([]entity.Slot, error)
	//CloseDB()
}

type service struct {
	slots manager.SlotRepository
}

func NewSlotService(repo manager.SlotRepository) SlotService {
	return &service{
		slots: repo,
	}
}

func (service *service) Save(slot entity.Slot) (entity.Slot, error) {
	err := service.slots.Save(slot)
	return slot, err
}

func (service *service) FindAll() []entity.Slot {
	return service.slots.FindAll()
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
