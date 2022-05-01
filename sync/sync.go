package sync

import "sync"

type Counter struct {
	lock  sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Increment() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
