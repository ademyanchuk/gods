// Binary Search Tree Implementation
package bstree

// Tree Node Structure
type Node struct {
	Left  *Node
	Right *Node
	Val   int
}

// Tree Structure
type Tree struct {
	Root *Node
}

// Build a new Tree from `values`
func Build(values []int) *Tree {
	tree := &Tree{}
	for _, val := range values {
		tree.Insert(val)
	}
	return tree
}

// Insert according to BS Tree Invariant
func (tree *Tree) Insert(value int) {
	if tree.Root == nil {
		tree.Root = NewNode(value)
		return
	}
	tree.Root.Insert(value)
}

// Node constructor
func NewNode(value int) *Node {
	return &Node{Left: nil, Right: nil, Val: value}
}

// Insert according to BS Tree Invariant
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
