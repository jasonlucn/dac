package dac

//二叉树
type BinNodeI interface {
	size() int
	insertAsLC(e interface{}) *BinNode
	insertAsRC(e interface{}) *BinNode
	succ() *BinNode //中序遍历的后继
	succIn() *BinNode
	stature() int      //节点高度
	balFac() int       //平衡因子
	avlBalanced() bool //是否avl平衡
	isLChild() bool
	tallerChild() *BinNode
	isRChild() bool
	updateHeightAbove()
	updateHeight() int
	attachAsLChild(v *BinNode) *BinNode
	attachAsRChild(v *BinNode) *BinNode
}

const (
	R Color = 0 >> iota
	B
)

type Color int

type BinNode struct {
	data   interface{}
	parent *BinNode
	lChild *BinNode
	rChild *BinNode
	height int
	color  Color //红黑树使用
}

func (n *BinNode) attachAsLChild(v *BinNode) *BinNode {
	n.lChild = v
	v.parent = n
	return v
}

func (n *BinNode) attachAsRChild(v *BinNode) *BinNode {
	n.rChild = v
	v.parent = n
	return v
}

func (n *BinNode) tallerChild() *BinNode {
	if n.lChild != nil {
		return n.lChild
	} else if n.rChild != nil {
		return n.rChild
	} else {
		return nil
	}
}

func (n *BinNode) isLChild() bool {
	return n.parent.lChild == n
}

func (n *BinNode) isRChild() bool {
	return n.parent.rChild == n
}

func (n *BinNode) balFac() int {
	return n.lChild.stature() - n.rChild.stature()
}

func (n *BinNode) avlBalanced() bool {
	return (-2 < n.balFac()) && (n.balFac() < 2)
}

func (n *BinNode) succIn() *BinNode {
	if n == nil {
		return nil
	}
	if n.lChild == nil {
		return n
	} else {
		return n.lChild.succIn()
	}
}

func (n *BinNode) succ() *BinNode {
	node := n.rChild
	succ := node.succIn()
	return succ
}

func (n *BinNode) insertAsLC(e interface{}) *BinNode {
	n.lChild = &BinNode{data: e, parent: n}
	return n.lChild
}

func (n *BinNode) insertAsRC(e interface{}) *BinNode {
	n.rChild = &BinNode{data: e, parent: n}
	return n.rChild
}

func (n *BinNode) size() int {
	size := 1
	if n.lChild != nil {
		size += n.lChild.size()
	}
	if n.rChild != nil {
		size += n.rChild.size()
	}
	return size
}

func (n *BinNode) stature() int {
	if n == nil {
		return -1
	}
	return n.height
}

func (n *BinNode) updateHeight() int {
	n.height = 1 + max(n.lChild.stature(), n.rChild.stature())
	return n.height
}

func (n *BinNode) updateHeightAbove() {
	for n != nil {
		n.updateHeight()
		n = n.parent
	}
}

type BinTreeI interface {
	size() int
	empty() bool
	root() *BinNode
	insertAsRoot(e interface{}) *BinNode
	visitAlongLeftBranch(n *BinNode, visit func(data interface{}), s *Stack2)
	travelPreI2(n *BinNode, visit func(data interface{}))
	travelLevel(n *BinNode, visit func(data interface{}))
	rebuild(preOrders []interface{}, inOrders interface{})
}

type BinTree struct {
	_size int
	_root *BinNode
}

func (t *BinTree) size() int {
	return t._size
}

func (t *BinTree) empty() bool {
	return t._root == nil
}

func (t *BinTree) root() *BinNode {
	return t._root
}

func (t *BinTree) insertAsLC(n *BinNode, e interface{}) *BinNode {
	t._size++
	n.insertAsLC(e)
	n.updateHeightAbove()
	return n.lChild
}

func (t *BinTree) insertAsRC(n *BinNode, e interface{}) *BinNode {
	t._size++
	n.insertAsRC(e)
	n.updateHeightAbove()
	return n.rChild
}

func (t *BinTree) visitAlongLeftBranch(n *BinNode, visit func(data interface{}), s *Stack2) {
	for n != nil {
		visit(n.data)
		//fmt.Printf("lChild %v rChild %v n.data %v\n", n.lChild, n.rChild, n.data)
		s.Push(n.rChild)
		n = n.lChild
	}
}

func (t *BinTree) visitLeftBranchInOrder(n *BinNode, visit func(data interface{}), rs *Stack2) {
	for n != nil {
		rs.Push(n)
		n = n.lChild
	}
}

//前序遍历
func (t *BinTree) travelPreI2(n *BinNode, visit func(data interface{})) {
	s := NewStack()
	for {
		t.visitAlongLeftBranch(n, visit, s)
		if s.Empty() {
			break
		}
		n = s.Pop().(*BinNode)
	}
}

//中序遍历
func (t *BinTree) travelInOrderI2(n *BinNode, visit func(data interface{})) {
	rs := NewStack()
	for {
		t.visitLeftBranchInOrder(n, visit, rs)
		if rs.Empty() {
			break
		}
		n = rs.Pop().(*BinNode)
		visit(n.data)
		n = n.rChild
	}
}

//层次遍历, 广度优先遍历
func (t *BinTree) travelLevel(n *BinNode, visit func(data interface{})) {
	q := Queue{}
	q.Init()
	q.Enqueue(n)
	for {
		if q.Empty() {
			break
		}
		current := q.Dequeue().(*BinNode)
		visit(current.data)
		if current.lChild != nil {
			q.Enqueue(current.lChild)
		}
		if current.rChild != nil {
			q.Enqueue(current.rChild)
		}
	}
}

func (t *BinTree) insertAsRoot(e interface{}) *BinNode {
	t._root = &BinNode{data: e}
	return t._root
}

func (t *BinTree) rebuild(preOrders []interface{}, inOrders interface{}) {
	t.insertAsRoot(preOrders[0])
}

func getLeftRight(preOrders []interface{}, inOrders []interface{}) (root interface{}, preLeft, inLeft, preRight, inRight []interface{}) {
	root = preOrders[0]
	rootInPos := searchPos(inOrders, root)
	inLeft = inOrders[0:rootInPos]
	inRight = inOrders[rootInPos:]
	
	return root, nil, nil, nil, nil
}

func searchPos(searches []interface{}, e interface{}) (pos int) {
	pos = -1
	for i := 0; i < len(searches); i++ {
		if searches[i] == e {
			return i
		}
	}
	return pos
}
