package dac

import (
	"fmt"
	"testing"
)

func initVector() *Vector {
	v := &Vector{}
	v.Init(2)
	v.Insert(1)
	v.Insert(3)
	v.Insert(5)
	v.Insert(7)
	v.Insert(9)
	return v
}

func TestVector_Insear(t *testing.T) {
	v := initVector()
	fmt.Println(v.Traverse())
}

func TestVector_Search(t *testing.T) {
	v := initVector()
	t1 := v.Search(1)
	if t1 != 0 {
		t.Errorf("t1 not equal 0")
	}
	t4 := v.Search(4)
	if t4 != 1 {
		t.Errorf("t4 not equal 1")
	}
	t11 := v.Search(11)
	if t11 != 4 {
		t.Errorf("t11 not equal 4")
	}
	t9 := v.Search(9)
	if t9 != 3 {
		t.Errorf("t9 not equal 3")
	}
	fmt.Println(v.Search(-1))
}
