package manager

import (
	"Slot_booking/entity"
	"gorm.io/gorm"
)

type SlotRepository interface {
	Create(slot []entity.Slot) error
	FindAll(dates []string) ([]entity.Slot, error)
	Find(slot entity.Slot) (entity.Slot, error)
	GetSlots(slotIDs []uint64) ([]entity.Slot, error)
	GetCount(date string) (int64, error)
}

type SlotDB struct {
	connection *gorm.DB
}

func SlotRepo() SlotRepository {
	return &SlotDB{
		connection: dbClient,
	}
}

func (db *SlotDB) Create(slot []entity.Slot) error {
	err := db.connection.Create(&slot).Error
	return err
}

func (db *SlotDB) FindAll(dates []string) ([]entity.Slot, error) {
	var slots []entity.Slot
	err := db.connection.Where("date in (?)", dates).Order("date").Order("start_time").Find(&slots).Error
	return slots, err
}

func (db *SlotDB) Find(slot entity.Slot) (entity.Slot, error) {
	err := db.connection.Where(&slot).Find(&slot).Error
	return slot, err
}
func (db *SlotDB) GetSlots(slotIDs []uint64) ([]entity.Slot, error) {
	var slot []entity.Slot
	err := db.connection.Model(&entity.Slot{}).Where("id in (?)", slotIDs).Order("date").Order("start_time").Find(&slot).Error
	return slot, err
}

func (db *SlotDB) GetCount(date string) (int64, error) {
	var count int64
	err := db.connection.Model(&entity.Slot{}).Where("date=?", date).Count(&count).Error
	return count, err
}
