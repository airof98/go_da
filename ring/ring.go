package ring

import (
	_ "fmt"
)

type Ring struct {
	next, prev *Ring
	Value      interface{}
}

func (r *Ring) init() *Ring {
	r.prev = r
	r.next = r
	return r
}

func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}

func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}

	return r.next
}

func (r *Ring) Prev() *Ring {
	if r.next == nil {
		return r.init()
	}

	return r.prev
}

func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

func (r *Ring) UnLink(s *Ring) *Ring {
	return nil
}
