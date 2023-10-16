package handler

import "sync"

type safeInt struct {
	mu *sync.Mutex
	v  int
}

func NewSafeInt(v int) *safeInt {
	return &safeInt{
		mu: &sync.Mutex{},
		v:  v,
	}
}

func (c *safeInt) AddDelta(value int) {
	c.mu.Lock()
	c.v += value
	c.mu.Unlock()
}

func (c *safeInt) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v
}
