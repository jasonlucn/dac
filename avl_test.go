package dac

import (
	"fmt"
	"testing"
)

func InitAVLTree() *AVL {
	t := &AVL{}
	t.insertAsRoot(4)
	t.Insert(1)
	t.Insert(2)
	t.Insert(3)
	t.Insert(5)
	//t.Insert(6)
	//t.Insert(7)
	return t
}

func TestAVL_Insert(t *testing.T) {
	tree := InitAVLTree()
	fmt.Println(tree)
}
