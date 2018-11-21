package onlinecounter

import (
	"github.com/go-redis/redis"
)

type RedisOnlineCounter struct {
	Redis *redis.Client
}

func (r RedisOnlineCounter) Online(timeKey string, id int64) error {
	return r.Redis.SetBit(timeKey, id, 1).Err()
}

func (r RedisOnlineCounter) Count(timeKey string) (int64, error) {
	return r.Redis.BitCount(timeKey, &redis.BitCount{0, -1}).Result()
}
