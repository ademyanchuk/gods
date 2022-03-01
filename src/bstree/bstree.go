// Binary Search Tree Implementation
package bstree

type Node struct {
	left  *Node
	right *Node
	value int
}

type Tree struct {
	root *Node
}

func (tree *Tree) insert(value int) {
	if tree.root == nil {
		tree.root = NewNode(value)
		return
	}
	tree.root.insert(value)
}

func NewNode(value int) *Node {
	return &Node{left: nil, right: nil, value: value}
}

func (n *Node) insert(value int) {
	if value < n.value {
		if n.left == nil {
			n.left = NewNode(value)
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = NewNode(value)
		} else {
			n.right.insert(value)
		}
	}
}
