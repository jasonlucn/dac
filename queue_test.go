package dac

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := Queue{}
	q.Init()
	q.Enqueue("a")
	q.Enqueue("b")
	fmt.Println(q.Travel())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}
