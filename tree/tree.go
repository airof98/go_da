package tree

import (
	_ "fmt"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func Insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}

	if v < t.Value {
		t.Left = Insert(t.Left, v)
	} else {
		t.Right = Insert(t.Right, v)
	}

	return t
}
