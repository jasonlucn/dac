package dac

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	g := new(GraphMatrix)
	g.E.Init(11)
	g.V.Init(11)
	
	g.insertV("a") //0
	g.insertV("s") //1
	g.insertV("e") //2
	g.insertV("d") //3
	g.insertV("c") //4
	g.insertV("f") //5
	g.insertV("b") //6
	g.insertV("g") //7
	
	g.insertV("h") //8
	g.insertV("i") //9
	g.insertV("j") //10
	
	g.insertE(1, 0, 1, 0) //a-s
	g.insertE(1, 0, 2, 0) //a-e
	g.insertE(1, 0, 4, 0) //a-c
	
	g.insertE(1, 1, 0, 0) //s-a
	g.insertE(1, 1, 3, 0) //s-d
	g.insertE(1, 1, 4, 0) //s-c
	
	g.insertE(1, 2, 0, 0) //e-a
	g.insertE(1, 2, 5, 0) //e-f
	g.insertE(1, 2, 7, 0) //e-g
	
	g.insertE(1, 3, 1, 0) //d-s
	g.insertE(1, 3, 6, 0) //d-b
	
	g.insertE(1, 4, 1, 0) //c-s
	g.insertE(1, 4, 6, 0) //c-b
	g.insertE(1, 4, 0, 0) //c-a
	
	g.insertE(1, 5, 2, 0) //f-e
	g.insertE(1, 5, 7, 0) //f-g
	
	g.insertE(1, 6, 3, 0) //b-d
	g.insertE(1, 6, 4, 0) //b-c
	g.insertE(1, 6, 7, 0) //b-g
	
	g.insertE(1, 7, 3, 0) //g-e
	g.insertE(1, 7, 3, 0) //g-f
	g.insertE(1, 7, 3, 0) //g-b
	
	//第二个连通域
	//g.insertE(1, 8, 9, 0)  //h-i
	//g.insertE(1, 8, 10, 0) //h-j
	//
	//g.insertE(1, 9, 8, 0) //i-h
	//
	//g.insertE(1, 10, 8, 0) //j-h
	
	//clock := 0
	//g.BFS(1, &clock)
	//g.MulBFS(1)
	
	clock := 0
	g.DFS(1, &clock)
	
	//fmt.Println(1)
	
	//results := g.V.Traverse()
	//for _, result := range results {
	//	fmt.Println(result.(*Vertex).data)
	//}
	////g.insertV("s")
	//
	//edges := g.E.Traverse()
	//for i, edge := range edges {
	//	v := edge.(*Vector)
	//	for j, val := range v.Traverse() {
	//		fmt.Printf("i %d j %d val %v\n", i, j, val.(*Edge).data)
	//	}
	//}
}

func TestDfs(t *testing.T) {
	g := new(GraphMatrix)
	g.E.Init(11)
	g.V.Init(11)
	
	g.insertV("j") //0
	g.insertV("i") //1
	g.insertV("g") //2
	g.insertV("h") //3
	g.insertV("d") //4
	g.insertV("e") //5
	g.insertV("f") //6
	g.insertV("a") //7
	g.insertV("b") //8
	g.insertV("c") //9
	
	g.insertE(1, 0, 2, 0) //j-g
	
	g.insertE(1, 1, 2, 0) //i-g
	g.insertE(1, 1, 4, 0) //i-d
	
	g.insertE(1, 2, 0, 0) //g-j
	g.insertE(1, 2, 1, 0) //g-i
	g.insertE(1, 2, 3, 0) //g-h
	g.insertE(1, 2, 4, 0) //g-d
	g.insertE(1, 2, 6, 0) //g-f
	
	g.insertE(1, 3, 2, 0) //h-g
	g.insertE(1, 3, 6, 0) //h-f
	
	g.insertE(1, 4, 1, 0) //d-i
	g.insertE(1, 4, 7, 0) //d-a
	g.insertE(1, 4, 2, 0) //d-g
	
	g.insertE(1, 5, 7, 0) //e-a
	g.insertE(1, 5, 8, 0) //e-b
	
	g.insertE(1, 6, 2, 0) //f-g
	g.insertE(1, 6, 3, 0) //f-h
	g.insertE(1, 6, 9, 0) //f-c
	
	g.insertE(1, 7, 4, 0) //a-d
	g.insertE(1, 7, 5, 0) //a-e
	g.insertE(1, 7, 8, 0) //a-b
	
	g.insertE(1, 8, 9, 0) //b-c
	g.insertE(1, 8, 7, 0) //b-a
	g.insertE(1, 8, 5, 0) //b-e
	
	g.insertE(1, 8, 6, 0) //c-f
	g.insertE(1, 8, 5, 0) //c-b
	
	
	clock := 0
	g.DFS(7, &clock)
	fmt.Println(g)
}
