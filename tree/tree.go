package tree

import (
	"fmt"
)

type Tree struct {
	Left   *Tree
	Value  int
	Right  *Tree
	Parent *Tree
	Color  byte
}

func New() *Tree {
	return &Tree{nil, 0, nil, nil, 2}
}

func Insert(t *Tree, v int) {
	if t.Color == 2 {
		t.Color = 1
		t.Value = v
		return
	}

	if v < t.Value {
		insertLeft(t, v)
	} else {
		insertRight(t, v)
	}
}

func insertLeft(t *Tree, v int) {
	if t.Left == nil {
		t.Left = &Tree{nil, v, nil, t, 0}
		return
	}

	if v < t.Value {
		insertLeft(t.Left, v)
	} else {
		insertRight(t.Right, v)
	}
}

func insertRight(t *Tree, v int) {
	if t.Right == nil {
		t.Right = &Tree{nil, v, nil, t, 0}
		return
	}

	if v < t.Value {
		insertLeft(t.Left, v)
	} else {
		insertRight(t.Right, v)
	}
}

func RightRotate(p *Tree, r *Tree) {
	p.Parent = r.Parent
	p.Right = r
	r.Left = p.Right
}

func LeftRotate(p *Tree, r *Tree) {
	p.Parent = r.Parent
	p.Right = p.Left
	p.Left = r
}

func Draw(t *Tree) {
	if t == nil {
		return
	}

	fmt.Println(t)
	fmt.Print(t.Left)
	fmt.Print("         ")
	fmt.Print(t.Right)
	fmt.Println("")

	Draw(t.Left)
	Draw(t.Right)
}
