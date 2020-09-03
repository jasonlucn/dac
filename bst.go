package dac

//二叉搜索树
type BSTI interface {
	Search(e interface{}) *BinNode
	Insert(e interface{}) *BinNode
	Remove(e interface{}) bool
	searchIn(v *BinNode, e interface{}, hot *BinNode) *BinNode
	removeAt(x *BinNode, hot *BinNode) *BinNode
	connect34(a, b, c, T0, T1, T2, T3 *BinNode) *BinNode
	rotateAt(v *BinNode) *BinNode
}

type BST struct {
	BinTree
	_hot *BinNode
}

func (t *BST) rotateAt(v *BinNode) *BinNode {
	p := v.parent
	g := p.parent
	if p.isLChild() {
		if v.isLChild() { //zig-zig
			p.parent = g.parent
			return t.connect34(v, p, g, v.lChild, v.rChild, p.rChild, g.rChild)
		} else { //zig-zag
			v.parent = g.parent
			return t.connect34(p, v, g, p.lChild, v.lChild, v.rChild, g.rChild)
		}
	} else {
		if v.isLChild() { //zag-zig
			v.parent = g.parent
			return t.connect34(g, v, p, g.lChild, v.lChild, v.rChild, p.rChild)
		} else { //zag-zag
			p.parent = g.parent
			return t.connect34(g, p, v, g.lChild, p.lChild, v.lChild, v.rChild)
		}
	}
}

func (t *BST) connect34(a, b, c, T0, T1, T2, T3 *BinNode) *BinNode {
	a.lChild = T0
	if T0 != nil {
		T0.parent = a
	}
	a.rChild = T1
	if T1 != nil {
		T1.parent = a
	}
	a.updateHeight()
	
	c.lChild = T2
	if T2 != nil {
		T2.parent = c
	}
	c.rChild = T3
	if T3 != nil {
		T3.parent = c
	}
	c.updateHeight()
	
	b.lChild = a
	a.parent = b
	b.rChild = c
	c.parent = b
	b.updateHeight()
	return b
}

//节点删除
func (t *BST) removeAt(x *BinNode, hot *BinNode) *BinNode {
	w := x
	var succ *BinNode
	if x.lChild == nil { //左子树为空, 使用右子树replace
		succ = x.rChild
		x = x.rChild
	} else if x.rChild == nil { //右子树为空, 使用左子树replace
		succ = x.lChild
		x = x.lChild
	} else { //左子树跟右子树并存
		w = w.succ()
		x.data, w.data = w.data, x.data
		u := w.parent
		if u == x {
			u.rChild = w.rChild
		} else {
			u.lChild = w.rChild
		}
		succ = w.rChild
	}
	t._hot = w.parent
	if succ != nil {
		succ.parent = t._hot
	}
	return succ
}

func (t *BST) searchIn(v *BinNode, e interface{}, hot *BinNode) *BinNode {
	if v == nil || v.data == e {
		return v
	}
	t._hot = v
	if e.(int) < v.data.(int) {
		return t.searchIn(v.lChild, e, t._hot)
	} else {
		return t.searchIn(v.rChild, e, t._hot)
	}
}

func (t *BST) Search(e interface{}) *BinNode {
	return t.searchIn(t._root, e, t._hot)
}

func (t *BST) Insert(e interface{}) *BinNode {
	_ = t.Search(e)
	if e.(int) < t._hot.data.(int) {
		return t.insertAsLC(t._hot, e)
	} else {
		return t.insertAsRC(t._hot, e)
	}
}

func (t *BST) Remove(e interface{}) bool {
	x := t.Search(e)
	if x == nil {
		return false
	}
	t.removeAt(x, t._hot)
	t._size--
	t._hot.updateHeightAbove()
	return true
}

func (t *BST) IsRoot(x *BinNode) bool {
	return x == t._root
}
