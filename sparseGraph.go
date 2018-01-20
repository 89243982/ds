package search

import "fmt"

type SparseGraph struct {
	n, m     int
	directed bool
	g        []([]*Edge)
}

func NewSparseGraph(n int, directed bool) *SparseGraph {
	sg := &SparseGraph{n: n, directed: directed, g: make([]([]*Edge), n)}
	for i, _ := range sg.g {
		sg.g[i] = make([]*Edge, 0, 8)
	}
	return sg
}

func (sg *SparseGraph) V() int {
	return sg.n
}
func (sg *SparseGraph) E() int {
	return sg.m
}
func (sg *SparseGraph) AddEdge(v, w int, weight int) {
	sg.g[v] = append(sg.g[v], &Edge{a: v, b: w, weight: weight})
	if v != w && !sg.directed {
		sg.g[w] = append(sg.g[w], &Edge{a: w, b: v, weight: weight})
	}
	sg.m++
}

func (sg *SparseGraph) HasEdge(v, w int) bool {
	vvs := sg.g[v]
	for _, e := range vvs {
		if e.Other(v) == w {
			return true
		}
	}
	return false
}

func (sg *SparseGraph) Show() {
	for i, vvs := range sg.g {
		fmt.Printf("vertex %d:\t", i)
		for _, v := range vvs {
			// cout<<"( to:"<<g[i][j]->w()<<",wt:"<<g[i][j]->wt()<<")\t";
			fmt.Printf("(to:->%d,wt:%d)\t", v.W(), v.Wt())
		}
		fmt.Println()
	}
}

func (sg *SparseGraph) Iterator(v int) Iterator {
	return NewSparseGraphAdjIterator(sg, v)
}

type SparseGraphAdjIterator struct {
	sg    *SparseGraph
	v     int
	vvs   []*Edge
	index int
}

func NewSparseGraphAdjIterator(sg *SparseGraph, v int) *SparseGraphAdjIterator {
	return &SparseGraphAdjIterator{sg: sg, v: v, vvs: sg.g[v]}
}

func (it *SparseGraphAdjIterator) Begin() *Edge {
	it.index = 0
	if len(it.vvs) > 0 {
		return it.vvs[it.index]
	}
	return nil
}
func (it *SparseGraphAdjIterator) Next() *Edge {
	it.index++
	if it.index < len(it.vvs) {
		return it.vvs[it.index]
	}
	return nil
}

func (it *SparseGraphAdjIterator) End() bool {
	return it.index >= len(it.vvs)
}
