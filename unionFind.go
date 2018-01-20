package search

// type UnionFind struct {
// 	id    []int
// 	count int
// }
//
// func NewUionFind(size int) *UnionFind {
// 	uf := &UnionFind{id: make([]int, size), count: size}
// 	for i, _ := range uf.id {
// 		uf.id[i] = i
// 	}
// 	return uf
// }
//
// // 查找过程, 查找元素p所对应的集合编号
// func (uf *UnionFind) Find(p int) int {
// 	return uf.id[p]
// }
//
// // 查看元素p和元素q是否所属一个集合
// // O(1)复杂度
// func (uf *UnionFind) IsConnected(p, q int) bool {
// 	return uf.id[p] == uf.id[q]
// }

// // 合并元素p和元素q所属的集合
// // O(n) 复杂度
// func (uf *UnionFind) Union(p, q int) {
// 	pId := uf.Find(p)
// 	qId := uf.Find(q)
// 	if pId == qId {
// 		return
// 	}
// 	// 合并过程需要遍历一遍所有元素, 将两个元素的所属集合编号合并
// 	for i := 0; i < uf.count; i++ {
// 		if uf.id[i] == pId {
// 			uf.id[i] = qId
// 		}
// 	}
//
// }
// type UnionFind struct {
// 	parent []int
// 	count  int
// }
//
// func NewUionFind(size int) *UnionFind {
// 	uf := &UnionFind{parent: make([]int, size), count: size}
// 	// 初始化, 每一个parent[i]指向自己, 表示每一个元素自己自成一个集合
// 	for i, _ := range uf.parent {
// 		uf.parent[i] = i
// 	}
// 	return uf
// }
//
// // 查找过程, 查找元素p所对应的集合编号
// // O(h)复杂度, h为树的高度
// func (uf *UnionFind) Find(p int) int {
// 	for p != uf.parent[p] {
// 		p = uf.parent[p]
// 	}
// 	return p
// }
//
// // 查看元素p和元素q是否所属一个集合
// // O(h)复杂度, h为树的高度
// func (uf *UnionFind) IsConnected(p, q int) bool {
// 	return uf.Find(p) == uf.Find(q)
// }
//
// // 合并元素p和元素q所属的集合
// // O(h)复杂度, h为树的高度
// func (uf *UnionFind) Union(p, q int) {
// 	qRoot := uf.Find(p)
// 	pRoot := uf.Find(q)
// 	if qRoot == pRoot {
// 		return
// 	}
// 	uf.parent[p] = pRoot
// }
// type UnionFind struct {
// 	parent []int
// 	sz     []int
// 	count  int
// }
//
// func NewUionFind(size int) *UnionFind {
// 	uf := &UnionFind{parent: make([]int, size), sz: make([]int, size), count: size}
// 	for i, _ := range uf.parent {
// 		uf.parent[i] = i
// 		uf.sz[i] = 1
// 	}
// 	return uf
// }
//
// // 查找过程, 查找元素p所对应的集合编号
// // O(h)复杂度, h为树的高度
// func (uf *UnionFind) Find(p int) int {
// 	for p != uf.parent[p] {
// 		p = uf.parent[p]
// 	}
// 	return p
// }
//
// func (uf *UnionFind) IsConnected(p, q int) bool {
// 	return uf.Find(p) == uf.Find(q)
// }
//
// // 合并元素p和元素q所属的集合
// // O(h)复杂度, h为树的高度
// func (uf *UnionFind) Union(p, q int) {
// 	pRoot := uf.Find(p)
// 	qRoot := uf.Find(q)
// 	if pRoot == qRoot {
// 		return
// 	}
// 	// 根据两个元素所在树的元素个数不同判断合并方向
// 	// 将元素个数少的集合合并到元素个数多的集合上
// 	if uf.sz[pRoot] < uf.sz[qRoot] {
// 		uf.parent[pRoot] = qRoot
// 		uf.sz[qRoot] += uf.sz[pRoot]
// 	} else {
// 		uf.parent[qRoot] = pRoot
// 		uf.sz[pRoot] += uf.sz[qRoot]
// 	}
// }

type UnionFind struct {
	parent []int
	rank   []int
}

//
func NewUnionFind(size int) *UnionFind {
	uf := &UnionFind{make([]int, size), make([]int, size)}
	for i, _ := range uf.parent {
		uf.parent[i] = i
		uf.rank[i] = 1
	}
	return uf
}

//
// /**查找p属于哪个组**/
// func (uf *UnionFind) Find(p int) int {
// 	for p != uf.parent[p] {
// 		p = uf.parent[p]
// 	}
// 	return p
// }
//
// func (uf *UnionFind) IsConnected(p, q int) bool {
// 	return uf.Find(p) == uf.Find(q)
// }
//
// func (uf *UnionFind) Union(p, q int) {
// 	// if uf.IsConnected(p, q) {
// 	// 	return
// 	// }
// 	pRoot := uf.Find(p)
// 	qRoot := uf.Find(q)
// 	if pRoot == qRoot {
// 		return
// 	}
// 	if uf.rank[pRoot] < uf.rank[qRoot] {
// 		uf.parent[pRoot] = qRoot
// 	} else if uf.rank[qRoot] < uf.rank[pRoot] {
// 		uf.parent[qRoot] = pRoot
// 	} else {
// 		uf.parent[pRoot] = qRoot
// 		uf.rank[qRoot]++
// 	}
// }

//路径压缩
func (uf *UnionFind) Find(p int) int {
	if p != uf.parent[p] {
		uf.parent[p] = uf.Find(uf.parent[p])
	}
	return uf.parent[p]
}
