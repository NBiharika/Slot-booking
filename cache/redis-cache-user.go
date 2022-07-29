package cache

import (
	"Slot_booking/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"time"
)

type UserCache interface {
	SetUser(ctx *gin.Context, key string, value entity.User)
	GetUser(ctx *gin.Context, key string) (entity.User, error)
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
	fmt.Println("CheckKey", key)
	jsonData, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	client.Set(ctx, key, jsonData, OneMonth)
	fmt.Println(client, string(jsonData))
}

func (cache *redisCache) GetUser(ctx *gin.Context, key string) (entity.User, error) {
	client := cache.getClient()

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return entity.User{}, err
	}

	users := entity.User{}
	err = json.Unmarshal([]byte(val), &users)
	if err != nil {
		panic(err)
	}
	fmt.Println("1111111111", users)
	return users, err
}
