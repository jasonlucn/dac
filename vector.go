package dac

import (
	"fmt"
	"sync"
)

type VectorI interface {
	Get(pos int) interface{}
	Insert(n interface{})
	Traverse() []interface{}
	Disordered() interface{}
	Unify()
	Find(e interface{}, lo, int, hi int) int
	BubbleSort(lo int, hi int)
	MergeSort(lo int, hi int)
	Size() int
	Remove(n int) interface{}
	Init(cap int)
	Empty() bool
	Set(pos int, n interface{})
	Search(e interface{}) int
}

type Vector struct {
	cap   int
	size  int
	elem  []interface{}
	mutex *sync.RWMutex
}

//返回不大于e的秩, 有序
func (v *Vector) Search(e interface{}) int {
	if e.(int) <= v.elem[0].(int) {
		return -1
	}
	for i := 0; i < v.size-1; i++ {
		if v.elem[i].(int) <= e.(int) && e.(int) < v.elem[i+1].(int) {
			return i
		}
	}
	return v.size - 1
}

func (v *Vector) Get(pos int) interface{} {
	return v.elem[pos]
}

func (v *Vector) Insert(n interface{}) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	v.expand()
	v.elem[v.size] = n
	v.size++
}

func (v *Vector) expand() {
	if v.cap < 2*v.size { //expand cap
		newElem := make([]interface{}, 2*v.size)
		for i := 0; i < v.size; i++ {
			newElem[i] = v.elem[i]
		}
		v.elem = newElem
		v.cap = 2 * v.size
	}
}

func (v *Vector) Set(pos int, n interface{}) {
	v.elem[pos] = n
}

func (v *Vector) Traverse() []interface{} {
	return v.elem[0:v.Size()]
}

//only to int
func (v *Vector) Disordered() interface{} {
	n := 0
	for i := 1; i < v.size; i++ {
		if v.elem[i-1].(int) > v.elem[i].(int) {
			n++
		}
	}
	return n
}

//有序
func (v *Vector) Unify() {
	//n := 0
	i := 0
	j := 1
	for {
		if v.elem[i] == v.elem[j] {
			i++
			j++
		} else {
			v.elem[i] = v.elem[j]
			j++
			//最后一个元素copy
			if j == v.size {
				i++
			}
		}
		if j >= v.size {
			break
		}
	}
	fmt.Printf("i %d j %d elem %v\n", i, j, v.elem)
	v.elem = v.elem[0:i]
}

func (v *Vector) Find(e interface{}, lo, int, hi int) int {
	for {
		
		if lo >= hi {
			break
		}
	}
	return 1
}

//有序向量的二分查找
func binSearch(nums []int, target int) int {
	lo := 0
	hi := len(nums)
	for lo < hi {
		mi := (lo + hi) / 2
		if target == nums[mi] {
			return mi
		} else if target < nums[mi] {
			hi = mi
		} else {
			lo = mi + 1
		}
	}
	return -1
}

func (v *Vector) BubbleSort(lo int, hi int) {
	for lo < hi {
		lastSwap := hi
		for i := lo; i < hi; i++ {
			if v.elem[i].(int) > v.elem[i+1].(int) {
				tmp := v.elem[i]
				v.elem[i] = v.elem[i+1]
				v.elem[i+1] = tmp
				lastSwap = i + 1
				fmt.Printf("lastSwap %d hi %d\n", i+1, hi)
			}
		}
		hi = lastSwap - 1
	}
}

func (v *Vector) MergeSort(lo int, hi int) {
	if hi-lo <= 1 {
		return
	}
	mi := (lo + hi) >> 1
	fmt.Printf("MergeSort lo %d mi %d hi %d\n", lo, mi, hi)
	v.MergeSort(lo, mi)
	v.MergeSort(mi, hi)
	//fmt.Printf("Merge lo %d mi %d hi %d\n", mi, lo, hi)
	v.Merge(lo, mi, hi)
	fmt.Printf("Merge lo %d mi %d hi %d elem %v\n", mi, lo, hi, v.elem)
}

func (v *Vector) Merge(lo int, mi int, hi int) {
	if hi-lo <= 1 {
		return
	}
	fmt.Printf("Merge lo %d mi %d hi %d v %v\n", lo, mi, hi, v.elem)
	A := &Vector{size: hi - lo}
	A.elem = make([]interface{}, hi-lo)
	lb := mi - lo
	lc := hi - mi
	B := &Vector{size: lb}
	B.elem = v.elem[lo:mi]
	C := &Vector{size: lc}
	C.elem = v.elem[mi:hi]
	fmt.Printf("b %v lb %d c %v lc %d\n", B.elem, lb, C.elem, lc)
	for i, j, k := 0, 0, 0; j < lb || k < lc; i++ {
		if j < lb && (lc <= k || (B.elem[j].(int) <= C.elem[k].(int))) {
			A.elem[i] = B.elem[j]
			j++
		}
		if k < lc && (lb <= j || (C.elem[j].(int) < B.elem[k].(int))) {
			A.elem[i] = C.elem[k]
			k++
		}
	}
	for i, _ := range A.elem {
		v.elem[i] = A.elem[i]
	}
}

func (v *Vector) Size() int {
	return v.size
}

func (v *Vector) Remove(n int) interface{} {
	ret := v.elem[n]
	for ; n < v.size; n++ {
		v.elem[n] = v.elem[n+1]
	}
	v.size--
	return ret
}

func (v *Vector) Init(cap int) {
	v.mutex = &sync.RWMutex{}
	v.elem = make([]interface{}, cap)
	v.cap = cap
}

func (v *Vector) Empty() bool {
	return v.Size() == 0
}
