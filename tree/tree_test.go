package tree_test

import (
	"../tree"
	"fmt"
	"testing"
)

func TestBinary(t *testing.T) {
	nodes := [...]int{3, 4, 1}
	bt := tree.New()

	for i, n := range nodes {
		fmt.Println(i, n, bt)
		tree.Insert(bt, n)
	}

	tree.Draw(bt)
}
