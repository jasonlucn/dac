package dac

type BTNode struct {
	parent *BTNode
	key    Vector
	child  Vector
}

type BTreeI interface {
	Search(e interface{}) *BTNode
	Insert(e interface{}) bool
	Remove(e interface{}) bool
	solveOverflow(btNode *BTNode)
	solveUnderflow(btNode *BTNode)
}

type BTree struct {
	_size  int
	_root  *BTNode
	_order int
	_hot   *BTNode
}

func (t *BTree) Search(e interface{}) *BTNode {
	v := t._root
	t._hot = nil
	for v != nil {
		r := v.key.Search(e)
		if r >= 0 && v.key.Get(r) == e {
			return v
		}
		v = v.child.Get(r + 1).(*BTNode)
	}
	return nil
}

func (t *BTree) Insert(e interface{}) bool {
	panic("implement me")
}

func (t *BTree) Remove(e interface{}) bool {
	panic("implement me")
}

func (t *BTree) solveOverflow(btNode *BTNode) {
	panic("implement me")
}

func (*BTree) solveUnderflow(btNode *BTNode) {
	panic("implement me")
}
