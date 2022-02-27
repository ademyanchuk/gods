// Binary Search Tree Implementation
package bstree

type node struct {
	left  *node
	right *node
	value int
}

type Tree struct {
	root *node
}

func (tree *Tree) insert(value int) {
	if tree.root == nil {
		tree.root = newNode(value)
		return
	}
	tree.root.insert(value)
}

func newNode(value int) *node {
	return &node{left: nil, right: nil, value: value}
}

func (n *node) insert(value int) {
	if value < n.value {
		if n.left == nil {
			n.left = newNode(value)
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = newNode(value)
		} else {
			n.right.insert(value)
		}
	}
}
