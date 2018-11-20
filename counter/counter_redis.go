package counter

import (
	"github.com/yacen/go-redis-app/idgenerator"
)

type RedisCounter struct {
	idgenerator.RedisIdGenerator
}

func (c RedisCounter) Decr(key string, by ...int64) (int64, error) {
	if by == nil {
		return c.Redis.Decr(key).Result()
	}
	return c.Redis.DecrBy(key, by[0]).Result()
}
