package bstree

import "testing"

func TestInsert(t *testing.T) {
	tree := &Tree{}
	value := 123
	tree.Insert(value)
	if tree.Root.Val != value {
		t.Errorf("got %q, want %q", tree.Root.Val, value)
	}
}

func TestInsertNodes(t *testing.T) {
	tree := &Tree{}
	value := 123
	left_val := 12
	tree.Insert(value)
	tree.Insert(left_val)
	if tree.Root.Val != value {
		t.Errorf("got %q, want %q", tree.Root.Val, value)
	}
	if tree.Root.Left.Val != left_val {
		t.Errorf("got %q, want %q", tree.Root.Left.Val, left_val)
	}
}
