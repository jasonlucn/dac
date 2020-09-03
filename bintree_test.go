package dac

import (
	"fmt"
	"testing"
)

func initBinTree() *BinTree {
	tree := new(BinTree)
	tree._root = &BinNode{data: "a"}
	
	tree.insertAsLC(tree._root, "b")
	nodeC := tree.insertAsRC(tree._root, "c")
	nodeD := tree.insertAsLC(nodeC, "d")
	nodeF := tree.insertAsRC(nodeC, "f")
	tree.insertAsRC(nodeD, "e")
	tree.insertAsLC(nodeF, "g")
	return tree
}

func TestTravelPreI2(t *testing.T) {
	tree := new(BinTree)
	tree._root = &BinNode{data: "a"}
	
	tree.insertAsLC(tree._root, "b")
	nodeC := tree.insertAsRC(tree._root, "c")
	nodeD := tree.insertAsLC(nodeC, "d")
	nodeF := tree.insertAsRC(nodeC, "f")
	tree.insertAsRC(nodeD, "e")
	tree.insertAsLC(nodeF, "g")
	//tree.travelPreI2(tree._root, func(data interface{}) {
	//	fmt.Println(data)
	//})
	
	//tree.travelInOrderI2(tree._root, func(data interface{}) {
	//	fmt.Println(data)
	//})
	
	tree.travelLevel(tree._root, func(data interface{}) {
		fmt.Println(data)
	})
	
}

func TestBinNode_Succ(t *testing.T) {
	tree := initBinTree()
	node := tree._root.succ()
	if node == nil || node.data != "d" {
		t.Errorf("%v succ is error", tree._root.data)
	}
}
