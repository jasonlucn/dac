package dac

import "sync"

type (
	Stack2 struct {
		top    *node
		length int
		lock   *sync.RWMutex
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new stack
func NewStack() *Stack2 {
	return &Stack2{nil, 0, &sync.RWMutex{}}
}

// Return the number of items in the stack
func (this *Stack2) Len() int {
	return this.length
}

func (this *Stack2) Empty() bool {
	return this.length <= 0
}

// View the top item on the stack
func (this *Stack2) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

// Pop the top item of the stack and return it
func (this *Stack2) Pop() interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.length == 0 {
		return nil
	}
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the stack
func (this *Stack2) Push(value interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	n := &node{value, this.top}
	this.top = n
	this.length++
}

func (this *Stack2) Reverse() []interface{} {
	result := make([]interface{}, this.Len())
	for this.Len() > 0 {
		result = append(result, this.Pop())
	}
	return result
}

type Stack interface {
	Push(e int)
	Pop() (n int)
	Top() (n int)
}

type MyStack struct {
	Vector
}

func (m *MyStack) Push(e int) {
	m.Insert(e)
}

func (m *MyStack) Pop() interface{} {
	return m.Remove(m.Size() - 1)
}

func (m *MyStack) Top() interface{} {
	return m.Get(m.Size() - 1)
}
