package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

var n int = 100

type IntS []int

func (p IntS) Len() int           { return len(p) }
func (p IntS) Less(i, j int) bool { return p[i] < p[j] }
func (p IntS) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

var ints []int

func init() {
	ints = make([]int, 0)
	for i := 0; i < n; i++ {
		ints = append(ints, rand.Intn(n))
	}
}

func TestInsertionSort(t *testing.T) {
	d := IntS(ints[0:])
	fmt.Println(d)
	Sort(d, "insertion")
	fmt.Println(ints)
}

func TestQuickSort(t *testing.T) {
	d := IntS(ints[0:])
	fmt.Println(d)
	Sort(d, "quick")
	fmt.Println(ints)
}

func TestBubbleSort(t *testing.T) {
	d := IntS(ints[0:])
	fmt.Println(d)
	Sort(d, "bubble")
	fmt.Println(ints)
}
