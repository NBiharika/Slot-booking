package cache

import (
	"Slot_booking/entity"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"time"
)

type UserCache interface {
	SetUser(ctx *gin.Context, key string, value entity.User)
	GetUser(ctx *gin.Context, key string) (entity.User, error)
	RemoveCache(ctx *gin.Context, key string) error
}

func NewRedisCache(host string, db int, exp time.Duration) UserCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) SetUser(ctx *gin.Context, key string, user entity.User) {
	client := cache.getClient()
	jsonData, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	client.Set(ctx, key, jsonData, OneMonth)
}

func (cache *redisCache) GetUser(ctx *gin.Context, key string) (entity.User, error) {
	client := cache.getClient()

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return entity.User{}, err
	}

	users := entity.User{}
	err = json.Unmarshal([]byte(val), &users)
	if err != nil {
		panic(err)
	}
	return users, err
}

func (cache *redisCache) RemoveCache(ctx *gin.Context, key string) error {
	client := cache.getClient()
	return client.Del(ctx, key).Err()
}
