package ring

import (
	"fmt"
	"testing"
)

func TestRing(t *testing.T) {
	r := New(4)
	var (
		r0 *Ring
		//r1 ring.Ring
	)
	r.Link(r0)
	//r.Link(&r1)
	fmt.Println(r)
	for p := r.Next(); p != r; p = p.Next() {
		fmt.Println(p)
	}

	fmt.Println("")
	fmt.Println(r)
	for p := r.Next(); p != r; p = p.Next() {
		fmt.Println(p)
	}
}
