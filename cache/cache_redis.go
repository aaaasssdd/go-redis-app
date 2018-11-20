package cache

import (
	"github.com/go-redis/redis"
	"time"
)

type RedisCacher struct {
	Redis *redis.Client
}

func (c RedisCacher) Get(key string) (string, error) {
	return c.Redis.Get(key).Result()
}

func (c RedisCacher) Set(key string, value interface{}, expire time.Duration) error {
	return c.Redis.Set(key, value, expire).Err()
}
