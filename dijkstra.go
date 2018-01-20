package search

import (
	"fmt"
)

/**
Dijkstra算法求最短路径
**/
type Dijkstra struct {
	g Graph
	s int
	// distTo[i]存储从起始点s到i的最短路径长度
	distTo []int
	// 标记数组, 在算法运行过程中标记节点i是否已经找到最短路径
	marked []bool
	//from[i]记录最短路径中, 到达i点的边是哪一条
	from []*Edge
}

func NewDijkstra(g Graph, s int) *Dijkstra {
	dka := &Dijkstra{g, s, make([]int, g.V()),
		make([]bool, g.V()), make([]*Edge, g.V())}
	// for i, _ := range dka.distTo {
	// 	dka.distTo[i] = math.MinInt32
	// }
	ipq := NewIndexMinHeap(g.V())
	dka.distTo[s] = 0
	dka.from[s] = NewEdge(s, s, 0)
	ipq.Insert(s, dka.distTo[s])
	dka.marked[s] = true
	for !ipq.IsEmpty() {
		v := ipq.ExtractMinIndex()
		dka.marked[v] = true
		it := dka.g.Iterator(v)
		for e := it.Begin(); !it.End(); e = it.Next() {
			w := e.Other(v)
			if !dka.marked[w] {
				if dka.from[w] == nil || dka.distTo[v]+e.Wt() < dka.distTo[w] {
					dka.distTo[w] = dka.distTo[v] + e.Wt()
					dka.from[w] = e
					if ipq.Contain(w) {
						// fmt.Printf("w:%d\n", w)
						// fmt.Printf("d[w]:%d\n", dka.distTo[w])
						ipq.Change(w, dka.distTo[w])
					} else {
						ipq.Insert(w, dka.distTo[w])
					}
				}
			}
		}
	}
	return dka
}

func (dka *Dijkstra) ShortestPathTo(w int) int {
	return dka.distTo[w]
}

func (dka *Dijkstra) HasPathTo(w int) bool {
	return dka.marked[w]
}
func (dka *Dijkstra) ShortestPath(w int) []Edge {
	stack := make([]*Edge, 0, dka.g.V())
	e := dka.from[w]
	for e.V() != dka.s {
		stack = append(stack, e)
		e = dka.from[e.V()]
	}
	stack = append(stack, e)

	paths := make([]Edge, 0, dka.g.V())
	for len(stack) > 0 {
		e = stack[len(stack)-1]
		paths = append(paths, *e)
		stack = stack[:len(stack)-1]
	}
	return paths
}

func (dka *Dijkstra) ShowPath(w int) {
	paths := dka.ShortestPath(w)
	for i, e := range paths {
		fmt.Printf("%d->", e.V())
		if i == len(paths)-1 {
			fmt.Println(e.W())
		}
	}
}
