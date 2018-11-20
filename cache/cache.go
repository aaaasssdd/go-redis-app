package cache

import (
	"time"
)

type Cacher interface {
	Get(string) (string, error)
	Set(string, interface{}, time.Duration) error
}

type CacherWrapper struct {
	Cacher Cacher
}

// expire = 0 means the key has no expire time.
func (c *CacherWrapper) Set(key string, value interface{}, expire time.Duration) error {
	return c.Cacher.Set(key, value, expire)
}

func (c *CacherWrapper) Get(key string) (string, error) {
	return c.Cacher.Get(key)
}
