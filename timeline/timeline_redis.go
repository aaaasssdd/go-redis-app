package timeline

import (
	"github.com/go-redis/redis"
)

type RedisTimeline struct {
	Redis *redis.Client
}

func (t *RedisTimeline) Push(key, message string) (err error) {
	return t.Redis.LPush(key, message).Err()
}

func (t *RedisTimeline) FetchIndex(key string, start int64, length int64) ([]string, error) {
	return t.Redis.LRange(key, start, start+length-1).Result()
}
