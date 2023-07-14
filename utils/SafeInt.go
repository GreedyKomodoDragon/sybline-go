package utils

import "sync"

type SafeInt struct {
	mu *sync.Mutex
	v  int
}

func NewSafeInt(v int) *SafeInt {
	return &SafeInt{
		mu: &sync.Mutex{},
		v:  v,
	}
}

func (c *SafeInt) AddDelta(value int) {
	c.mu.Lock()
	c.v += value
	c.mu.Unlock()
}

func (c *SafeInt) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v
}
