package dac

type ListNode struct {
	pred *ListNode
	succ *ListNode
	Data interface{}
}

type ListI interface {
	Init()
	InsertBefore(node *ListNode, newNode *ListNode) *ListNode
	Remove(node *ListNode) *ListNode
	SelectMax(node *ListNode, n int) *ListNode
	SelectionSort(node *ListNode, n int)
	Travel() []interface{}
	Search(e interface{}, n int, node *ListNode) *ListNode
	InsertAfter(node *ListNode, e interface{}) *ListNode
	InsertionSort(node *ListNode, n int)
	First() *ListNode
	Empty() bool
}

type List struct {
	header  *ListNode
	_size   int
	trailer *ListNode
}

func (l *List) Init() {
	l.header = new(ListNode)
	l.trailer = new(ListNode)
	l.header.succ = l.trailer
	l.header.pred = nil
	l.trailer.succ = nil
	l.trailer.pred = l.header
	l._size = 0
}

func (l *List) InsertAsLast(e interface{}) *ListNode {
	newNode := &ListNode{Data: e}
	predNode := l.trailer.pred
	newNode.pred = predNode
	newNode.succ = l.trailer
	l.trailer.pred = newNode
	predNode.succ = newNode
	l._size++
	return newNode
}

func (l *List) Empty() bool {
	return l._size <= 0
}

func (l *List) InsertBefore(currentNode *ListNode, newNode *ListNode) *ListNode {
	predNode := currentNode.pred
	newNode.pred = predNode
	newNode.succ = currentNode
	currentNode.pred = newNode
	predNode.succ = newNode
	l._size++
	return newNode
}

func (l *List) Remove(currentNode *ListNode) *ListNode {
	predNode := currentNode.pred
	succNode := currentNode.succ
	predNode.succ = succNode
	succNode.pred = predNode
	l._size--
	return currentNode
}

func (l *List) First() *ListNode {
	return l.header.succ
}

//只处理int
func (l *List) SelectMax(node *ListNode, n int) *ListNode {
	max := node
	for cur := node.succ; n > 1; n-- {
		if max.Data.(int) <= cur.Data.(int) {
			max = cur
		}
		cur = cur.succ
	}
	return max
}

func (l *List) SelectionSort(node *ListNode, n int) {
	head := node.pred
	tail := node
	for i := 0; i < n; i++ {
		tail = tail.succ
	}
	for n > 1 {
		p := l.Remove(l.SelectMax(head.succ, n))
		//fmt.Printf("p.data %d n %d alldata %v tail.data %d\n", p.data, n, l.Travel(), tail.data)
		l.InsertBefore(tail, p)
		//fmt.Printf("p.data %d n %d alldata %v tail.data %d\n", p.data, n, l.Travel(), tail.data)
		tail = tail.pred
		n--
	}
}

func (l *List) Travel() []interface{} {
	data := make([]interface{}, l._size)
	cur := l.header
	for i := 0; i < l._size; i++ {
		cur = cur.succ
		data[i] = cur.Data
	}
	return data
}

//有序的列表
func (l *List) Search(e interface{}, n int, p *ListNode) *ListNode {
	for n >= 0 {
		p = p.pred
		if p.Data.(int) <= e.(int) {
			break
		}
		n--
	}
	return p
}

func (l *List) InsertAfter(node *ListNode, e interface{}) *ListNode {
	nextNode := node.succ
	newNode := &ListNode{Data: e}
	node.succ = newNode
	nextNode.pred = newNode
	newNode.pred = node
	newNode.succ = nextNode
	l._size++
	return newNode
}

func (l *List) InsertionSort(p *ListNode, n int) {
	for r := 0; r < n; r++ {
		s := l.Search(p.Data, r, p)
		//fmt.Printf("r %d s.data %d p.data %d\n", r, s.data, p.data)
		l.InsertAfter(s, p.Data)
		//fmt.Printf("r %d s.data %d ipred.data %d isucc.data %d psucc.data %d\n", r, s.data, i.pred.data, i.succ.data, p.succ.data)
		p = p.succ
		l.Remove(p.pred)
		//fmt.Println(l.Travel())
	}
}
