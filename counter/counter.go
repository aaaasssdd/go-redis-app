package counter

import "github.com/yacen/go-redis-app/idgenerator"

type Counter interface {
	idgenerator.IdGenerator
	Decr(string, ...int64) (int64, error)
}

type CounterWrapper struct {
	idgenerator.IdGeneratorWrapper
	Counter Counter
}

func (c *CounterWrapper) Decr(key string, by ...int64) (int64, error) {
	return c.Counter.Decr(key, by...)
}
