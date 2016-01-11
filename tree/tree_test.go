package tree_test

import (
	"../tree"
	"fmt"
	"testing"
)

func TestBinary(t *testing.T) {
	nodes := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var bt *tree.Tree = nil

	for i, n := range nodes {
		fmt.Println(i, n)
		bt = tree.Insert(bt, n)
	}

	fmt.Println(bt)
	fmt.Println(bt.Left)
	fmt.Println(bt.Right)
}
