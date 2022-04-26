package queue

import (
	"fmt"
	"sync"
)

type customQueue struct {
	queue []string
	lock  sync.RWMutex
}

func (c *customQueue) Enqueue(name string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.queue = append(c.queue, name)
}

func (c *customQueue) Dequeue() error {
	if len(c.queue) > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		c.queue = c.queue[1:]
		return nil
	}
	return fmt.Errorf("Pop Error: Queue is empty")
}

func (c *customQueue) Front() (string, error) {
	if len(c.queue) > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		return c.queue[0], nil
	}
	return "", fmt.Errorf("Peep Error: Queue is empty")
}

func (c *customQueue) Size() int {
	return len(c.queue)
}

func (c *customQueue) Empty() bool {
	return len(c.queue) == 0
}

func SliceImplementationExample() {
	customQueue := &customQueue{
		queue: make([]string, 0),
	}

	fmt.Printf("Enqueue: A\n")
	customQueue.Enqueue("A")
	fmt.Printf("Enqueue: B\n")
	customQueue.Enqueue("B")
	fmt.Printf("Len: %d\n", customQueue.Size())

	// If size greater than zero
        for customQueue.Size() > 0 {
		frontVal, _ := customQueue.Front()
		fmt.Printf("Front: %s\n", frontVal)
		fmt.Printf("Dequeue: %s\n", frontVal)
		customQueue.Dequeue()
	}
	fmt.Printf("Len: %d\n", customQueue.Size())
}
