package cache

import (
	"Slot_booking/entity"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type SlotCache interface {
	SetSlot(ctx *gin.Context, key string, value []entity.Slot)
	GetSlot(ctx *gin.Context, key string) ([]entity.Slot, error)
}

func NewRedisCacheSlot(host string, db int, exp time.Duration) SlotCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) SetSlot(ctx *gin.Context, key string, slot []entity.Slot) {
	client := cache.getClient()
	jsonData, err := json.Marshal(slot)
	if err != nil {
		panic(err)
	}

	client.Set(ctx, key, jsonData, OneMonth)
	fmt.Println(client, key, string(jsonData))
}

func (cache *redisCache) GetSlot(ctx *gin.Context, key string) ([]entity.Slot, error) {
	client := cache.getClient()
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return []entity.Slot{}, err
	}

	slots := []entity.Slot{}
	err = json.Unmarshal([]byte(val), &slots)
	if err != nil {
		panic(err)
	}
	return slots, err
}
