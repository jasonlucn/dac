package dac

type PQCompleteHeapI interface {
	percolateUp(i int)   //上滤
	percolateDown(i int) //下滤
	parent(i int) int    //获取元素逻辑parent的秩
	lChild(i int) int    //获取元素逻辑lChild的秩
	rChild(i int) int    //获取元素逻辑rChild的秩
	newPQCompleteHeap(i int)
}

//完全二叉堆
type PQCompleteHeap struct {
	Vector
}

func NewPQCompleteHeap(e []interface{}, n int) *PQCompleteHeap {
	p := &PQCompleteHeap{}
	p.Init(2 * n)
	p.elem = e
	p.size = n
	return p
}

func (p *PQCompleteHeap) lChild(i int) int {
	return i>>1 + 1
}

func (p *PQCompleteHeap) rChild(i int) int {
	return i>>1 + 2
}

func (p *PQCompleteHeap) percolateUp(i int) {
	if i == 0 {
		return
	}
	parentRank := p.parent(i)
	if p.elem[parentRank].(int) < p.elem[i].(int) {
		p.elem[parentRank], p.elem[i] = p.elem[i], p.elem[parentRank]
	}
	p.percolateUp(parentRank)
}

func (p *PQCompleteHeap) percolateDown(i int) {
	if i == p.Size() {
		return
	}
	lRank := p.lChild(i)
	rRank := p.rChild(i)
	if p.elem[i].(int) < p.elem[lRank].(int) || p.elem[i].(int) < p.elem[rRank].(int) {
		if p.elem[lRank].(int) < p.elem[rRank].(int) {
			p.elem[i], p.elem[rRank] = p.elem[rRank], p.elem[i]
			p.percolateUp(rRank)
		} else {
			p.elem[i], p.elem[lRank] = p.elem[lRank], p.elem[i]
			p.percolateUp(lRank)
		}
	}
}

func (p *PQCompleteHeap) parent(i int) int {
	return (i - 1) >> 1
}

func (p *PQCompleteHeap) GetMax() interface{} {
	return p.elem[0]
}

func (p *PQCompleteHeap) DelMax() {
	p.elem[0], p.elem[p.Size()-1] = p.elem[p.Size()-1], p.elem[0]
	p.Remove(p.Size() - 1)
	p.percolateDown(0)
}

func (p *PQCompleteHeap) Insert(e interface{}) {
	p.Vector.Insert(e)
	rank := p.Size() - 1
	p.percolateUp(rank)
}
