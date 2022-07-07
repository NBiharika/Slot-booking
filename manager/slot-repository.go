package manager

import (
	"Slot_booking/entity"
	"gorm.io/gorm"
)

type SlotRepository interface {
	Save(slot entity.Slot) error
	CreateSlot(string) error
	FindAll() []entity.Slot
	Find(slot entity.Slot) (entity.Slot, error)
	GetSlots(slotIDs []uint64) ([]entity.Slot, error)
}

type SlotDB struct {
	connection *gorm.DB
}

func SlotRepo() SlotRepository {
	return &SlotDB{
		connection: dbClient,
	}
}

func (db *SlotDB) CreateSlot(startTime string) error {
	var slot entity.Slot
	slot.Date = entity.DateForSlot()
	slot.StartTime = startTime
	response := dbClient.Create(&slot)
	if response.Error != nil {
		return response.Error
	}
	return nil
}

func (db *SlotDB) Save(slot entity.Slot) error {
	err := db.connection.Create(&slot).Error
	//db.connection.AutoMigrate(&entity.Slot{})
	return err
}

func (db *SlotDB) FindAll() []entity.Slot {
	var slot []entity.Slot
	db.connection.Find(&slot)
	return slot
}

func (db *SlotDB) Find(slot entity.Slot) (entity.Slot, error) {
	err := db.connection.Where(&slot).Find(&slot).Error
	return slot, err
}
func (db *SlotDB) GetSlots(slotIDs []uint64) ([]entity.Slot, error) {
	var slot []entity.Slot

	err := db.connection.Model(&entity.Slot{}).Debug().Where("id in (?)", slotIDs).Find(&slot).Error
	return slot, err
}
