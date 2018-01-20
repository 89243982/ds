package search

type rbNode struct {
	key, value  interface{}
	red         bool
	left, right *rbNode
}

type RBTree struct {
	root *rbNode
	size int
	less func(a, b interface{}) bool
}

func NewRBTree(less func(a, b interface{}) bool) *RBTree {
	return &RBTree{less: less}
}

func isRed(node *rbNode) bool {
	return node != nil && node.red
}

func newRBNode(k, v interface{}) *rbNode {
	return &rbNode{key: k, value: v}
}

func rotateLeft(node *rbNode) *rbNode {
	rightChild := node.right
	node.right = rightChild.left
	rightChild.left = node
	rightChild.red = node.red
	node.red = true
	return rightChild
}

func rotateRight(node *rbNode) *rbNode {
	leftChild := node.left
	node.left = leftChild.right
	leftChild.right = node
	leftChild.red = node.red
	node.red = true
	return leftChild
}

func colorFlip(node *rbNode) *rbNode {
	node.red = !node.red
	if node.left != nil {
		node.left.red = !node.left.red
	}
	if node.right != nil {
		node.right.red = !node.right.red
	}
	return node
}

func fixUp(node *rbNode) *rbNode {
	if isRed(node.right) {
		node = rotateLeft(node)
	}
	if isRed(node.left) && isRed(node.left.left) {
		node = rotateRight(node)
	}
	if isRed(node.left) && isRed(node.right) {
		node = colorFlip(node)
	}
	return node
}

func (r *RBTree) Insert(k, v interface{}) bool {
	r.size++
	var ok bool
	r.root, ok = _insert(r.root, k, v, r.less)
	r.root.red = false
	return ok
}

func _insert(root *rbNode, k, v interface{},
	less func(a, b interface{}) bool) (*rbNode, bool) {
	ok := false
	if root == nil {
		return &rbNode{key: k, value: v, red: true}, true
	}
	if less(k, root.key) {
		root.left, ok = _insert(root.left, k, v, less)
	} else if less(root.key, k) {
		root.right, ok = _insert(root.right, k, v, less)
	} else {
		root.value = v
		ok = true
	}
	return fixUp(root), ok
}

func moveRedRight(node *rbNode) *rbNode {
	node = colorFlip(node)
	if node.left != nil && isRed(node.left.left) {
		node = rotateRight(node)
		node = colorFlip(node)
	}
	return node
}
func (r *RBTree) deleteMax() {
	if !isRed(r.root.left) && !isRed(r.root.right) {
		r.root.red = true
	}
	r.root = deleteMax(r.root)
	if !r.IsEmpty() {
		r.root.red = false
	}
}
func deleteMax(node *rbNode) *rbNode {
	if isRed(node.left) {
		node = rotateRight(node)
	}
	if node.right == nil {
		return nil
	}
	if !isRed(node.right) && !isRed(node.right.left) {
		node = moveRedRight(node)
	}
	node.right = deleteMax(node.right)
	return fixUp(node)
}

func moveRedLeft(node *rbNode) *rbNode {
	node = colorFlip(node)
	if isRed(node.right.left) {
		node.right = rotateRight(node.right)
		node = rotateLeft(node)
		node = colorFlip(node)
	}
	return node
}
func (r *RBTree) deleteMin() {
	if !isRed(r.root.left) && !isRed(r.root.right) {
		r.root.red = true
	}
	r.root = deleteMin(r.root)
	if !r.IsEmpty() {
		r.root.red = false
	}
}
func deleteMin(node *rbNode) *rbNode {
	if node.left == nil {
		return nil
	}
	if !isRed(node.left) && !isRed(node.left.left) {
		node = moveRedLeft(node)
	}
	node.left = deleteMin(node.left)
	return fixUp(node)
}

func (r *RBTree) Delete(k interface{}) bool {
	r.size--
	if !isRed(r.root.left) && !isRed(r.root.right) {
		r.root.red = true
	}
	var ok bool
	r.root, ok = _delete(r.root, k, r.less)
	if !r.IsEmpty() {
		r.root.red = false
	}
	return ok
}

func _delete(node *rbNode, k interface{}, less func(a, b interface{}) bool) (*rbNode, bool) {
	ok := false
	if less(k, node.key) {
		if node.left != nil {
			if !isRed(node.left) && !isRed(node.left.left) {
				node = moveRedLeft(node)
			}
			node.left, ok = _delete(node.left, k, less)
		}
	} else {
		if isRed(node.left) {
			node = rotateRight(node)
		}
		if !less(k, node.key) && !less(node.key, k) && node.right == nil {
			return nil, true
		}
		if node.right != nil {
			if !isRed(node.right) && !isRed(node.right.left) {
				node = moveRedRight(node)
			}
			if !less(k, node.key) && !less(node.key, k) {
				smallest := min(node.right)
				node.key = smallest.key
				node.value = smallest.value
				node.right = deleteMin(node.right)
				ok = true
			} else {
				node.right, ok = _delete(node.right, k, less)
			}
		}
	}
	return fixUp(node), ok
}

func min(node *rbNode) *rbNode {
	for node.left != nil {
		node = node.left
	}
	return node
}

func (r *RBTree) IsEmpty() bool {
	return r.size == 0
}
