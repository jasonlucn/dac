package dac

type AVL struct {
	BST
}

func (t *AVL) Insert(e interface{}) *BinNode {
	//x := t.Insert(e)
	//x := t.Search(e)
	//if x != nil {
	//	return x
	//}
	x := t.BST.Insert(e)
	//x = &BinNode{data: e, parent: t.hot}
	//t._size++
	xx := x
	//检测是否失衡,从 x的父亲向上一层层检测
	for g := x.parent; g != nil; g = g.parent {
		if !g.avlBalanced() {
			g = t.rotateAt(g.tallerChild().tallerChild())
			if g.parent == nil {
				t._root = g
			}
			break
		} else {
			g.updateHeight()
		}
	}
	return xx
}

func (t *AVL) Remove(e interface{}) bool {
	return true
}
