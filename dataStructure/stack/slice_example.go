package stack

import (
    "fmt"
    "sync"
)

type customStack struct {
    stack []string
    lock  sync.RWMutex
}

func (c *customStack) Push(name string) {
    c.lock.Lock()
    defer c.lock.Unlock()
    c.stack = append(c.stack, name)
}

func (c *customStack) Pop() error {
    len := len(c.stack)
    if len > 0 {
        c.lock.Lock()
        defer c.lock.Unlock()
        c.stack = c.stack[:len-1]
        return nil
    }
    return fmt.Errorf("Pop Error: Stack is empty")
}

func (c *customStack) Front() (string, error) {
    len := len(c.stack)
    if len > 0 {
        c.lock.Lock()
        defer c.lock.Unlock()
        return c.stack[len-1], nil
    }
    return "", fmt.Errorf("Peep Error: Stack is empty")
}

func (c *customStack) Size() int {
    return len(c.stack)
}

func (c *customStack) Empty() bool {
    return len(c.stack) == 0
}

func SliceInStackExample() {
    customStack := &customStack{
        stack: make([]string, 0),
    }
    fmt.Printf("Push: A\n")
    customStack.Push("A")
    fmt.Printf("Push: B\n")
    customStack.Push("B")
    fmt.Printf("Size: %d\n", customStack.Size())
    for customStack.Size() > 0 {
        frontVal, _ := customStack.Front()
        fmt.Printf("Front: %s\n", frontVal)
        fmt.Printf("Pop: %s\n", frontVal)
        customStack.Pop()
    }
    fmt.Printf("Size: %d\n", customStack.Size())
}
