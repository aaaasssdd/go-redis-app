package accumulator

import "github.com/yacen/go-redis-app/idgenerator"

type Accumulator interface {
	idgenerator.IdGenerator
	Decr(string, ...int64) (int64, error)
}

type AccumulatorWrapper struct {
	idgenerator.IdGeneratorWrapper
	Accumulator Accumulator
}

func (c *AccumulatorWrapper) Decr(key string, by ...int64) (int64, error) {
	return c.Accumulator.Decr(key, by...)
}
