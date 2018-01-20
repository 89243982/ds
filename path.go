package search

//
// import "fmt"
//
// type Path struct {
// 	g Graph
// 	// 起始点
// 	s       int
// 	visited []bool
// 	//记录路径, from[i]表示查找的路径上i的上一个节点
// 	from []int
// }

//
// func NewPath(g Graph, s int) *Path {
// 	path := &Path{g: g, s: s, visited: make([]bool, g.V()),
// 		from: make([]int, g.V())}
// 	for i, _ := range path.from {
// 		path.from[i] = -1
// 	}
// 	path.s = s
// 	path.DFS(s)
// 	return path
// }
//
// func (path *Path) DFS(v int) {
// 	path.visited[v] = true
// 	it := path.g.Iterator(v)
// 	for w := it.Begin(); !it.End(); w = it.Next() {
// 		if !path.visited[w] {
// 			path.from[w] = v
// 			path.DFS(w)
// 		}
// 	}
// }
//
// // 查询从s点到w点是否有路径
// func (path *Path) HasPath(w int) bool {
// 	return path.visited[w]
// }
//
// func (path *Path) Path(w int) (corresPath []int) {
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
// func (path *Path) ShowPath(w int) {
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
