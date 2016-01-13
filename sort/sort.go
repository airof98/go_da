package sort

import (
	_ "fmt"
)

type SortI interface {
	Less(i, j int) bool
	Swap(i, j int)
	Len() int
}

func insertionSort(data SortI, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

func quickSort(d SortI, a, b int) {
	if a >= b {
		return
	}

	i := a
	p := b
	j := b - 1
	if b-a <= 2 {
		insertionSort(d, a, b+1)
		return
	}

	for i < j {
		if !d.Less(i, p) {
			if !d.Less(p, i) {
				i++
			} else {
				d.Swap(i, j)
				d.Swap(j, p)
				p = j
				j = p - 1
			}
		} else {
			i++
		}

		if i == j && !d.Less(j, p) && d.Less(p, j) {
			d.Swap(i, p)
			p--
		}
	}

	quickSort(d, a, p-1)
	quickSort(d, p+1, b)
}

func bubbleSort(data SortI) {
	n := data.Len()
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if !data.Less(j, j+1) {
				data.Swap(j, j+1)
			}
		}
	}
}

func Sort(data SortI, t string) {
	if t == "insertion" {
		insertionSort(data, 0, data.Len())
	} else if t == "quick" {
		quickSort(data, 0, data.Len()-1)
	} else if t == "bubble" {
		bubbleSort(data)
	}
}
