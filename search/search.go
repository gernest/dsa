package search

func Sequential(src []int, item int) bool {
	for _, x := range src {
		if item == x {
			return true
		}
	}
	return false
}

func Binary(src []int, item int) bool {
	if len(src) == 0 {
		return false
	}
	first := 0
	last := len(src) - 1
	mid := (first + last) / 2
	midItem := src[mid]
	if midItem == item {
		return true
	}
	if midItem < item {
		return Binary(src[:mid], item)
	}
	return Binary(src[mid+1:], item)
}

type BinaryTree struct {
	root *TreeNode
	size int
}

func (b *BinaryTree) Length() int {
	return b.size
}

func (b *BinaryTree) Put(key int, val interface{}) {
	if b.root == nil {
		b.root = NewTreeNode(key, val)
	}
	if b.root != nil {
		b.put(key, val, b.root)
	}
	b.size = b.size + 1
}
func (b *BinaryTree) put(key int, val interface{}, current *TreeNode) {
	if key < current.Key {
		if current.left != nil {
			b.put(key, val, current.left)
			return
		}
		current.left = NewTreeNode(key, val, current)
		return
	}
	if current.right != nil {
		b.put(key, val, current.right)
		return
	}
	current.right = NewTreeNode(key, val, current)
}

type TreeNode struct {
	Key    int
	Val    interface{}
	parent *TreeNode
	left   *TreeNode
	right  *TreeNode
}

func (t *TreeNode) Collect() []int {
	var rst []int
	if t.left != nil {
		rst = append(rst, t.left.Collect()...)
	}
	rst = append(rst, t.Key)
	if t.right != nil {
		rst = append(rst, t.right.Collect()...)
	}
	return rst
}

func NewTreeNode(key int, val interface{}, parent ...*TreeNode) *TreeNode {
	node := &TreeNode{Key: key, Val: val}
	if len(parent) > 0 {
		node.parent = parent[0]
	}
	return node
}
