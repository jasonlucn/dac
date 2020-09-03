package dac

import (
	"fmt"
	"testing"
)

func TestMyStack_Pop(t *testing.T) {
	s := &MyStack{}
	s.Init(1000)
	s.Insert(3)
	s.Insert(2)
	s.Insert(5)
	fmt.Println(s.Traverse())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Traverse())
	fmt.Println(s.Traverse())
}
