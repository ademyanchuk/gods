package bstree

import (
	"reflect"
	"testing"
)

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
	right_val := 15
	tree.Insert(right_val)
	if tree.Root.Left.Right.Val != right_val {
		t.Errorf("got %q, want %q", tree.Root.Left.Right.Val, right_val)
	}
}

func TestBuild(t *testing.T) {
	values := []int{2, 1, 3}
	tree := Build(values)
	got := []int{tree.Root.Val, tree.Root.Left.Val, tree.Root.Right.Val}
	if !reflect.DeepEqual(values, got) {
		t.Errorf("got %v, want %v", got, values)
	}
}

func TestLevel(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{3, 5, 2, 4, 1}, []int{3, 2, 5, 1, 4}},
	}
	for _, c := range cases {
		tree := Build(c.in)
		got := tree.Level()
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Level(%v)==%v, want: %v", c.in, got, c.want)
		}
	}
}
