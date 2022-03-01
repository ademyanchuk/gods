// Binary Search Tree Implementation
package bstree

type Node struct {
	Left  *Node
	Right *Node
	Val   int
}

type Tree struct {
	Root *Node
}

func (tree *Tree) Insert(value int) {
	if tree.Root == nil {
		tree.Root = NewNode(value)
		return
	}
	tree.Root.Insert(value)
}

func NewNode(value int) *Node {
	return &Node{Left: nil, Right: nil, Val: value}
}

func (n *Node) Insert(value int) {
	if value < n.Val {
		if n.Left == nil {
			n.Left = NewNode(value)
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = NewNode(value)
		} else {
			n.Right.Insert(value)
		}
	}
}
