package cache

import (
	"github.com/go-redis/redis/v9"
	"time"
)

const (
	oneDay   = time.Hour * 24
	OneMonth = 30 * oneDay
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}
