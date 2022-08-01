package manager

import (
	"Slot_booking/entity"
	"gorm.io/gorm"
	"time"
)

type SlotRepository interface {
	Create(slot []entity.Slot) error
	FindAll(startDate string, endDate string) []entity.Slot
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
	//db.connection.AutoMigrate(&entity.Slot{})
	err := db.connection.Create(&slot).Error
	return err
}

func (db *SlotDB) FindAll(startDate string, endDate string) []entity.Slot {
	var slot []entity.Slot
	db.connection.Where("date>=? and date<=?", startDate, endDate).Find(&slot)
	return slot
}

func (db *SlotDB) Find(slot entity.Slot) (entity.Slot, error) {
	err := db.connection.Where(&slot).Find(&slot).Error
	return slot, err
}
func (db *SlotDB) GetSlots(slotIDs []uint64) ([]entity.Slot, error) {
	var slot []entity.Slot
	todayTime := time.Now()
	startDate := entity.DateForSlot(todayTime)
	endTime := todayTime.Add(6 * 24 * time.Hour)
	endDate := entity.DateForSlot(endTime)
	err := db.connection.Model(&entity.Slot{}).Where("id in (?) and date>=? and date<=?", slotIDs, startDate, endDate).Find(&slot).Error
	return slot, err
}

func (db *SlotDB) GetCount(date string) (int64, error) {
	var count int64
	err := db.connection.Model(&entity.Slot{}).Where("date=?", date).Count(&count).Error
	return count, err
}
