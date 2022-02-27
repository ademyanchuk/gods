package bstree

import "testing"

func TestInsert(t *testing.T) {
	tree := &Tree{}
	value := 123
	tree.insert(value)
	if tree.root.value != value {
		t.Errorf("got %q, want %q", tree.root.value, value)
	}
}

func TestInsertNodes(t *testing.T) {
	tree := &Tree{}
	value := 123
	left_val := 12
	tree.insert(value)
	tree.insert(left_val)
	if tree.root.value != value {
		t.Errorf("got %q, want %q", tree.root.value, value)
	}
	if tree.root.left.value != left_val {
		t.Errorf("got %q, want %q", tree.root.left.value, left_val)
	}
}
