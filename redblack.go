package dac

type RedBlackI interface {
	Insert(e interface{}) *BinNode
	Remove(e interface{}) bool
	solveDoubleRed(binNode *BinNode)   //双红修正
	solveDoubleBlack(binNode *BinNode) //双黑修正
	updateHeight(binNode *BinNode) int
}

type RedBlack struct {
	BST
}

func (t *RedBlack) Remove(e interface{}) bool {
	x := t.Search(e)
	_ = t.removeAt(x, t._hot)
	return false
}

func (t *RedBlack) Insert(e interface{}) *BinNode {
	x := t.BST.Insert(e)
	if t.IsRoot(x) {
		x.color = B
	} else {
		x.color = R
	}
	t.solveDoubleRed(x)
	if x != nil {
		return x
	} else {
		return t._hot.parent
	}
}

func (t *RedBlack) solveDoubleRed(x *BinNode) {
	p := x.parent
	g := p.parent
	if p.color == B {
		return
	}
	var u *BinNode
	if p.isLChild() {
		u = p.rChild
	} else {
		u = p.lChild
	}
	if u.color == B { //3+4
		if p.isLChild() {
			if x.isLChild() {
				x.color = R
				p.color = B
				g.color = R
			} else {
				p.color = R
				x.color = B
				g.color = R
			}
		} else {
			if x.isLChild() {
				x.color = R
				p.color = B
				g.color = R
			} else {
				p.color = R
				x.color = B
				g.color = R
			}
		}
	} else { //分裂
		if p.isLChild() {
			if x.isLChild() {
				g.color = R
				x.color = B
				u.color = B
			} else {
				p.color = B
				u.color = B
				g.color = R
			}
		}
		t.solveDoubleRed(g)
	}
}

func (t *RedBlack) solveDoubleBlack(binNode *BinNode) {
	panic("implement me")
}

func (t *RedBlack) updateHeight(binNode *BinNode) int {
	panic("implement me")
}
