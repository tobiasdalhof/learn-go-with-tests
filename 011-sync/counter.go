package counter

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

// type Counter struct {
// 	inc   chan int
// 	value int
// }

func NewCounter() *Counter {
	return &Counter{}
}

// func NewCounter() *Counter {
// 	return &Counter{inc: make(chan int, 1)}
// }

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// func (c *Counter) Inc() {
// 	c.inc <- 1
// 	select {
// 	case <-c.inc:
// 		c.value++
// 	}
// }

func (c *Counter) Value() int {
	return c.value
}
