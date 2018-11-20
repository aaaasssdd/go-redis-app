package idgenerator

import "github.com/go-redis/redis"

type RedisIdGenerator struct {
	Redis *redis.Client
}

func (c RedisIdGenerator) Incr(key string, by ...int64) (int64, error) {
	if by == nil {
		return c.Redis.Incr(key).Result()
	}
	return c.Redis.IncrBy(key, by[0]).Result()
}

func (c RedisIdGenerator) Get(key string) (int64, error) {
	return c.Redis.Get(key).Int64()
}

func (c RedisIdGenerator) Reset(key string) (int64, error) {
	return c.Redis.GetSet(key, 0).Int64()
}
