package sort

import (
	"fmt"
	"math/rand"
	_ "sort"
	"testing"
)

var n int = 10

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
	fmt.Println(n)
}

func TestInsertionSort(t *testing.T) {
	d := IntS(ints[0:])
	fmt.Println(d)
	Sort(d, "insertion")
	fmt.Println(ints)
}

func TestQuickSort(t *testing.T) {
	d := IntS(ints[0:])
	Sort(d, "quick")
}

func TestBubbleSort(t *testing.T) {
	//ints := []int{3, 4, 7, 5, 6, 2, 8, 1}
	d := IntS(ints[0:])
	fmt.Println(d)
	Sort(d, "bubble")
	fmt.Println(ints)
}
