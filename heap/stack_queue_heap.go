
type Stack struct {
    nodes []int
    c int
}

func (s *Stack) push(a int) {
    s.nodes = append(s.nodes, a)
    s.c++
}

func (s *Stack) pop() int {
    if s.c == 0 {
        return -1
    }
    
    a := s.nodes[s.c-1]
    s.nodes = s.nodes[:s.c-1]
    s.c--
    return a
}

type Queue struct {
    nodes []int
    c int
}

func (q *Queue) push(a int) {
    q.nodes = append([]int{a}, q.nodes...)
    q.c++
}

func (q *Queue) pop() int {
    if q.c == 0 {
        return -1
    }
    
    a := q.nodes[q.c-1]
    q.nodes = q.nodes[:q.c-1]
    q.c--
    return a
}

type Heap struct {
    nodes []int
    c int
    sz int
}

func NewHeap(hs int) *Heap {
    h := &Heap{}
    h.nodes = make([]int, hs)
    h.sz = hs
    return h
}

func (h *Heap) push(a int) {
    h.c++
    h.nodes[h.c] = a
    if h.c == 1 {
        return
    }
    
    p := h.c / 2
    c := h.c
    for p > 0 && h.nodes[p] < h.nodes[c] {       
        h.nodes[p], h.nodes[c] = h.nodes[c], h.nodes[p]
        c = p
        p = p / 2
    }
}

func (h *Heap) pop() int {
    if h.c == 0 {
        return -1
    }
    a := h.nodes[1]
    h.nodes[1] = h.nodes[h.c]
    h.c--
    if h.c == 1 {
        return a
    }
    
    p := 1
    for true {
        l := p * 2
        r := p * 2 + 1 
        x := h.nodes[r]
        xi := r
        if h.nodes[l] > h.nodes[r] {
            x = h.nodes[l]
            xi = l
        }
        
        if x > h.nodes[p] {
            h.nodes[p], h.nodes[xi] = x, h.nodes[p]
            p = l
        } else {
            break
        }
    }
    
    return a
}
