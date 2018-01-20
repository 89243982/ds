package search

import "strings"

type Node struct {
	key   string
	value int
	left  *Node
	right *Node
}

type BST struct {
	root  *Node
	count int
}

func (n *Node) GetNodeKey() string {
	return n.key
}
func NewBST() *BST {
	return &BST{nil, 0}
}

func (bst *BST) Insert(key string, value int) {
	bst.count++
	bst.root = insert(bst.root, key, value)
}
func (bst *BST) Size() int {
	return bst.count
}

func insert(root *Node, key string, value int) *Node {
	if root == nil {
		return &Node{key, value, nil, nil}
	}
	if strings.Compare(key, root.key) == 0 {
		root.value = value
	} else if strings.Compare(key, root.key) < 0 {
		root.left = insert(root.left, key, value)
	} else {
		root.right = insert(root.right, key, value)
	}
	return root
}

func (bst *BST) Search(key string) *int {
	return search(bst.root, key)
}

func search(root *Node, key string) *int {
	if root == nil {
		return nil
	}
	if strings.Compare(key, root.key) == 0 {
		return &root.value
	} else if strings.Compare(key, root.key) < 0 {
		return search(root.left, key)
	} else {
		return search(root.right, key)
	}
}

func (bst *BST) PreOrder(fn func(n *Node)) {
	preOrder(bst.root, fn)
}

func preOrder(root *Node, fn func(n *Node)) {
	if root != nil {
		fn(root)
		preOrder(root.left, fn)
		preOrder(root.right, fn)
	}
}
func (bst *BST) InOrder(fn func(n *Node)) {
	inOrder(bst.root, fn)
}

func inOrder(root *Node, fn func(n *Node)) {
	if root != nil {
		inOrder(root.left, fn)
		fn(root)
		inOrder(root.right, fn)
	}
}
func (bst *BST) PostOrder(fn func(n *Node)) {
	postOrder(bst.root, fn)
}
func postOrder(root *Node, fn func(n *Node)) {
	if root != nil {
		postOrder(root.left, fn)
		postOrder(root.right, fn)
		fn(root)
	}
}

func (bst *BST) Empty() bool {
	return bst.count == 0
}
func (bst *BST) LevelOrder(fn func(n *Node)) {
	q := make([](*Node), 0, bst.count)
	q = append(q, bst.root)
	for len(q) != 0 {
		n := q[0]
		q = q[1:]
		fn(n)
		if n.left != nil {
			q = append(q, n.left)
		}
		if n.right != nil {
			q = append(q, n.right)
		}
	}
}

func (bst *BST) Mini() string {
	/**
	count>0
	**/
	m := mini(bst.root)
	return m.key
}
func mini(root *Node) *Node {
	if root.left == nil {
		return root
	}
	return mini(root.left)
}
func (bst *BST) Max() string {
	m := max(bst.root)
	return m.key
}
func max(root *Node) *Node {
	if root.right == nil {
		return root
	}
	return max(root.right)
}
func (bst *BST) RemoveMin() {
	bst.count--
	bst.root = removeMin(bst.root)
}

func removeMin(root *Node) *Node {
	if root.left == nil {
		return root.right
	}
	root.left = removeMin(root.left)
	return root
}

func (bst *BST) RemoveMax() {
	bst.count--
	bst.root = removeMax(bst.root)
}

func removeMax(root *Node) *Node {
	if root.right == nil {
		return root.left
	}
	root.right = removeMax(root.right)
	return root
}

func (bst *BST) Remove(key string) {
	bst.count--
	bst.root = remove(bst.root, key)
}

func remove(root *Node, key string) *Node {
	if root == nil {
		return nil
	}
	if strings.Compare(key, root.key) < 0 {
		root.left = remove(root.left, key)
		return root
	} else if strings.Compare(root.key, key) < 0 {
		root.right = remove(root.right, key)
		return root
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}
		successor := mini(root.right)
		successor.right = removeMin(root.right)
		successor.left = root.left
		return successor
	}
}
