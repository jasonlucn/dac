package dac

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{6, 5, 4, 3, 2, 1}
	QuickSort(arr, 0, 5)
	if !reflect.DeepEqual(arr, []int{1, 2, 3, 4, 5, 6}) {
		t.Errorf("err")
	}
}
