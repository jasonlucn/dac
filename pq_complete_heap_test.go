package dac

import (
	"testing"
)

func initPQ() *PQCompleteHeap {
	pq := &PQCompleteHeap{}
	pq.Init(1000)
	pq.Vector.Insert(4)
	pq.Vector.Insert(3)
	pq.Vector.Insert(0)
	pq.Vector.Insert(1)
	pq.Vector.Insert(2)
	return pq
}

func TestPQCompleteHeap_Insert(t *testing.T) {
	pq := initPQ()
	pq.Insert(5)
	resultShould := []int{5, 3, 4, 1, 2, 0}
	for i, value := range pq.Traverse() {
		if value.(int) != resultShould[i] {
			t.Errorf("insert error %v", pq.Traverse())
		}
	}
}

func TestPQCompleteHeap_Down(t *testing.T) {
	pq := initPQ()
	pq.Insert(5)
	pq.DelMax()
	resultShould := []int{4, 3, 0, 1, 2}
	for i, value := range pq.Traverse() {
		if value.(int) != resultShould[i] {
			t.Errorf("remove error %v", pq.Traverse())
		}
	}
}
