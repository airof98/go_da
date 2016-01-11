package ring

import (
	_ "../ring"
	"container/ring"
	"fmt"
	"testing"
)

func TestRing(t *testing.T) {
	r := ring.New(4)
	var (
		r0 *ring.Ring
		//r1 ring.Ring
	)
	r.Link(r0)
	//r.Link(&r1)
	fmt.Println(r)
	for p := r.Next(); p != r; p = p.Next() {
		fmt.Println(p)
	}

	r.Unlink(1)
	fmt.Println("")
	fmt.Println(r)
	for p := r.Next(); p != r; p = p.Next() {
		fmt.Println(p)
	}
}
