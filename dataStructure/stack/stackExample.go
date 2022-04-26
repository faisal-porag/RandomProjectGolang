package stack

import (
    "container/list"
    "fmt"
)

/*
A stack will have below operations:

- Push
- Pop
- Front
- Size
- Empty

*/

type customStack struct {
    stack *list.List
}

func (c *customStack) Push(value string) {
    c.stack.PushFront(value)
}

func (c *customStack) Pop() error {
    if c.stack.Len() > 0 {
        ele := c.stack.Front()
        c.stack.Remove(ele)
    }
    return fmt.Errorf("Pop Error: Stack is empty")
}

func (c *customStack) Front() (string, error) {
    if c.stack.Len() > 0 {
        if val, ok := c.stack.Front().Value.(string); ok {
            return val, nil
        }
        return "", fmt.Errorf("Peep Error: Stack Datatype is incorrect")
    }
    return "", fmt.Errorf("Peep Error: Stack is empty")
}

func (c *customStack) Size() int {
    return c.stack.Len()
}

func (c *customStack) Empty() bool {
    return c.stack.Len() == 0
}

func StackExample() {
    customStack := &customStack{
        stack: list.New(),
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
