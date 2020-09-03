package dac

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{5, 3, 2, 1, 4, 6}
	HeapSort(arr)
	if !reflect.DeepEqual(arr, []int{1, 2, 3, 4, 5, 6}) {
		t.Errorf("err")
	}
}
