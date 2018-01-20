package search

//
// import "fmt"
//
// type ShortestPath struct {
// 	g       Graph
// 	s       int
// 	visited []bool
// 	from    []int
// 	ord     []int
// }
//
// func NewShortestPath(g Graph, s int) *ShortestPath {
// 	path := &ShortestPath{g: g, s: s, visited: make([]bool, g.V()),
// 		from: make([]int, g.V()), ord: make([]int, g.V())}
//
// 	for i, _ := range path.from {
// 		path.from[i] = -1
// 		path.ord[i] = -1
// 	}
//
// 	queue := make([]int, 0, 8)
// 	queue = append(queue, s)
// 	path.visited[s] = true
// 	path.ord[s] = 0
// 	for len(queue) > 0 {
// 		v := queue[0]
// 		queue = queue[1:]
// 		it := path.g.Iterator(v)
// 		for w := it.Begin(); !it.End(); w = it.Next() {
// 			if !path.visited[w] {
// 				queue = append(queue, w)
// 				path.from[w] = v
// 				path.ord[w] = path.ord[v] + 1
// 				path.visited[w] = true
// 			}
// 		}
//
// 	}
//
// 	return path
// }
//
// // 查询从s点到w点是否有路径
// func (path *ShortestPath) HasPath(w int) bool {
// 	return path.visited[w]
// }
//
// func (path *ShortestPath) Path(w int) (corresPath []int) {
// 	stack := make([]int, 0, 8)
// 	corresPath = make([]int, 0, 8)
// 	p := w
// 	for p != -1 {
// 		stack = append(stack, p)
// 		p = path.from[p]
// 	}
//
// 	for len(stack) > 0 {
// 		top := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]
// 		corresPath = append(corresPath, top)
// 	}
// 	return
// }
//
// func (path *ShortestPath) ShowPath(w int) {
// 	pth := path.Path(w)
// 	for i, v := range pth {
// 		fmt.Printf("%d", v)
// 		if i == len(pth)-1 {
// 			fmt.Println()
// 		} else {
// 			fmt.Printf("%s", "->")
// 		}
// 	}
// }
//
// func (path *ShortestPath) Length(w int) int {
// 	return path.ord[w]
// }
