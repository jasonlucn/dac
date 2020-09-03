package dac

import (
	"testing"
)

func InitBSTTree() *BST {
	tree := new(BST)
	tree.insertAsRoot(36)
	node27 := tree.insertAsLC(tree._root, 27)
	_ = tree.insertAsLC(node27, 6)
	node58 := tree.insertAsRC(tree._root, 58)
	node53 := tree.insertAsLC(node58, 53)
	_ = tree.insertAsRC(node58, 69)
	
	_ = tree.insertAsLC(node53, 46)
	return tree
}

func TestBST_Search(t *testing.T) {
	bst := InitBSTTree()
	bst.Search(53)
	if bst._hot.data != 58 {
		t.Errorf("search 53 parent should be 58")
	}
	bst.Search(90)
	if bst._hot.data != 64 {
		t.Errorf("search 90 parent should be 64")
	}
}

func TestBST_Insert(t *testing.T) {
	bst := InitBSTTree()
	node40 := bst.Insert(40)
	if node40.parent.data != 46 || node40.parent.lChild.data != 40 {
		t.Errorf("insert 40 error")
	}
	node55 := bst.Insert(55)
	if node55.parent.data != 53 || node55.parent.rChild.data != 55 {
		t.Errorf("insert 55 error")
	}
}

func TestBST_Remove(t *testing.T) {
	bst := InitBSTTree()
	bst.Remove(36)
	//fmt.Println(bst)
	if bst._root.data != 46 {
		t.Errorf("remove 6 error")
	}
}
