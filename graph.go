package dac

import (
	"fmt"
	"math"
)

//Vertex status
type VStatus int

const (
	UNDISCOVERED VStatus = 1 << iota
	DISCOVERED
	VISITED
)

//Edge status
type EStatus int

const (
	UNDETERMINED EStatus = 1 << iota
	TREE
	CROSS
	FORWARD
	BACKWARD
)

type VertexI interface {
	Init(d interface{})
	Status(s VStatus)
	fistNbr() *Vertex
	nextNbr() *Vertex
	DTime(s int)
}

//顶点
type Vertex struct {
	data      interface{}
	inDegree  int
	outDegree int
	status    VStatus
	dTime     int
	fTime     int
	parent    int
	priority  int
}

func (v *Vertex) fistNbr() *Vertex {
	panic("implement me")
}

func (v *Vertex) nextNbr(i, j int) *Vertex {
	panic("implement me")
}

func (v *Vertex) Status(s VStatus) {
	v.status = s
}

func (v *Vertex) DTime(s int) {
	v.dTime = s
}

type EdgeI interface {
	Status(s EStatus)
}

//边
type Edge struct {
	data   interface{}
	weight int
	status EStatus
}

func (e *Edge) Status(s EStatus) {
	e.status = s
}

func (v *Vertex) Init(d interface{}) {
	v.data = d
	v.inDegree, v.outDegree = 0, 0
	v.status = UNDISCOVERED
	v.dTime, v.fTime, v.parent = -1, -1, -1
	v.priority = math.MaxInt64
}

func (e *Edge) Init(d interface{}, w int) {
	e.data = d
	e.weight = w
	e.status = UNDETERMINED
}

type GraphMatrixI interface {
	dTime(i int) int
	nextNbr(i, j int) int
	firstNbr(i int) int
	edge(i, j int) interface{}
	status(i, j int) EStatus
	weight(i, j int) int
	insertE(e interface{}, i, j, w int) //插入边
	removeE(i, j int) *Edge             //删除边
	insertV(e interface{})              //插入顶点
	removeV(e interface{}) *Vertex      //删除顶点
	BFS(v int, clock *int)              //广度优先遍历
}

type GraphMatrix struct {
	V Vector
	E Vector
	n int //顶点数量GraphMatrix
	e int //边的数量
}

func (g *GraphMatrix) removeV(i int) *Vertex { //删除顶点及其关联的边
	for j := 0; j < g.n; j++ {
		if g.Exist(i, j) { //删除所有的出边
			g.Set(i, j, nil)
			g.V.Get(j).(*Vertex).inDegree--
		}
	}
	g.E.Remove(i)
	g.n--
	for j := 0; j < g.n; j++ {
		if g.Exist(j, i) { // 删除所有的入边
			g.Set(j, i, nil)
			g.V.Get(j).(*Vertex).outDegree--
		}
	}
	v := g.V.Get(i).(*Vertex)
	g.V.Remove(i)
	return v
}

func (g *GraphMatrix) insertV(e interface{}) {
	for j := 0; j < g.n; j++ {
		ej := g.E.Get(j)
		if ej != nil {
			e := ej.(*Vector)
			e.Insert(&Edge{})
		}
	}
	g.n++
	
	vec := &Vector{}
	vec.Init(100)
	for n := 0; n < g.n; n++ {
		vec.Insert(&Edge{})
	}
	g.E.Insert(vec)
	v := new(Vertex)
	v.Init(e)
	g.V.Insert(v)
}

func (g *GraphMatrix) removeE(i, j int) *Edge {
	edge := g.E.Get(i).(*Vector).Get(j).(*Edge)
	g.E.Get(i).(*Vector).Set(j, nil)
	g.e--
	g.V.Get(i).(*Vertex).outDegree--
	g.V.Get(j).(*Vertex).inDegree--
	return edge
}

func (g *GraphMatrix) insertE(e interface{}, i, j, w int) {
	if g.Exist(i, j) {
		return
	}
	edge := new(Edge)
	edge.Init(e, w)
	g.E.Get(i).(*Vector).Set(j, edge)
	g.e++
	g.V.Get(i).(*Vertex).outDegree++
	g.V.Get(j).(*Vertex).inDegree++
}

func (g *GraphMatrix) firstNbr(i int) int {
	return g.nextNbr(i, g.n)
}

func (g *GraphMatrix) nextNbr(i, j int) int {
	for -1 < j {
		j -= 1
		if g.Exist(i, j) {
			return j
		}
	}
	return -1
}

func (g *GraphMatrix) dTime(i int) int {
	return g.V.Get(i).(*Vertex).dTime
}

//判断边是否存在
func (g *GraphMatrix) Exist(i, j int) bool {
	return (0 <= i) && (i < g.n) && (0 <= j) && (j < g.n) && g.E.Get(i).(*Vector).Get(j).(*Edge).data != nil
}

//设置边的值
func (g *GraphMatrix) Set(i, j int, e interface{}) {
	g.E.Get(i).(*Vector).Set(j, e)
}

//多连通域遍历
func (g *GraphMatrix) MulBFS(s int) {
	clock := 0
	v := s
	for v < g.n {
		if g.V.Get(v).(*Vertex).status == UNDISCOVERED {
			g.BFS(v, &clock)
		}
		v++
	}
}

//广度优先遍历
func (g *GraphMatrix) BFS(v int, clock *int) {
	q := new(Queue)
	q.Init()
	g.V.Get(v).(*Vertex).Status(DISCOVERED)
	q.Enqueue(v)
	for !q.Empty() {
		v := q.Dequeue().(int)
		fmt.Println(g.V.Get(v).(*Vertex).data)
		*clock += 1
		g.V.Get(v).(*Vertex).DTime(*clock)
		for u := g.firstNbr(v); -1 < u; u = g.nextNbr(v, u) { //考察v的每一个邻居u
			if g.V.Get(u).(*Vertex).status == UNDISCOVERED { //如果邻居u未被发现
				g.V.Get(u).(*Vertex).status = DISCOVERED
				q.Enqueue(u)
				g.E.Get(v).(*Vector).Get(u).(*Edge).status = TREE
				fmt.Printf("[%d][%d] tree\n", v, u)
				g.V.Get(u).(*Vertex).parent = v
			} else {
				g.E.Get(v).(*Vector).Get(u).(*Edge).status = CROSS
			}
			g.V.Get(v).(*Vertex).Status(VISITED)
		}
	}
}

//深度优先遍历
func (g *GraphMatrix) DFS(v int, clock *int) {
	current := g.V.Get(v).(*Vertex)
	current.status = DISCOVERED
	*clock++
	current.dTime = *clock
	for u := g.firstNbr(v); -1 < u; u = g.nextNbr(v, u) {
		nbr := g.V.Get(u).(*Vertex)
		switch nbr.status {
		case UNDISCOVERED:
			g.E.Get(v).(*Vector).Get(u).(*Edge).status = TREE
			nbr.dTime = *clock
			nbr.parent = v
			//fmt.Println(nbr.data)
			g.DFS(u, clock)
			break
		case DISCOVERED:
			g.E.Get(v).(*Vector).Get(u).(*Edge).status = BACKWARD
		case VISITED:
			if g.V.Get(v).(*Vertex).dTime < g.V.Get(u).(*Vertex).dTime {
				g.E.Get(v).(*Vector).Get(u).(*Edge).status = FORWARD
			} else {
				g.E.Get(v).(*Vector).Get(u).(*Edge).status = CROSS
			}
		}
	}
	current.status = VISITED
	*clock++
	current.fTime = *clock
}
