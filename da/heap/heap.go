package heap

type Interface interface {
	Push(x interface{})
	Pop() interface{}
	Less(i, j int) bool
	Swap(i, j int)
	Len() int
}

func Init(h Interface) {
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

func Push(h Interface, v interface{}) {
	h.Push(v)
	up(h, h.Len()-1)
}

func Pop(h Interface) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

func up(h Interface, j int) {
	for {
		i := (j - 1) / 2
		if i == j || !h.Less(j, i) {
			break
		}

		h.Swap(i, j)
		j = i
	}
}

func down(h Interface, i, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && !h.Less(j1, j2) {
			j = j2
		}

		if !h.Less(j, i) {
			break
		}

		h.Swap(i, j)
		i = j
	}
}
