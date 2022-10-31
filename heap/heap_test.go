package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

type IntHeap []int

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h IntHeap) Len() int {
	return len(h)
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TestHeap(t *testing.T) {
	intHeap := &IntHeap{8, 1, 4, 3, 5, 2}
	heap.Init(intHeap)
	fmt.Println(intHeap)
	heap.Push(intHeap, 7)
	fmt.Println(intHeap)
	for intHeap.Len() > 0 {
		fmt.Println(heap.Pop(intHeap))
	}
}
