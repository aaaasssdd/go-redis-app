package idgenerator

import "fmt"

type IdGenerator interface {
	Incr(string, ...int64) (int64, error)
	Get(string) (int64, error)
	Reset(string) (int64, error)
}

type IdGeneratorWrapper struct {
	IdGenerator IdGenerator
}

func (c *IdGeneratorWrapper) Incr(key string, by ...int64) (int64, error) {
	return c.IdGenerator.Incr(key, by...)
}

func (c *IdGeneratorWrapper) Get(key string) (int64, error) {
	return c.IdGenerator.Get(key)
}

func (c *IdGeneratorWrapper) Reset(key string) (int64, error) {
	fmt.Println("jax1", c)
	fmt.Println("jax2", c.IdGenerator)
	return c.IdGenerator.Reset(key)
}
