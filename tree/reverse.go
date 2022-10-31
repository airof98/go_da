package tree

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	curr := head
	var pl, nr *ListNode
	i := 1
	var rln, rlnp *ListNode
	for curr != nil {
		if i >= left && i <= right {
			rln = &ListNode{Val: curr.Val, Next: rlnp}
			rlnp = rln
		} else if i < left {
			pl = curr
		} else if i > right && nr == nil {
			nr = curr
		}
		curr = curr.Next
		i++
	}

	if rln != nil {
		if pl != nil {
			pl.Next = rln
		}
		if nr != nil {
			curr = rln
			for curr != nil {
				if curr.Next == nil {
					curr.Next = nr
					break
				}
				curr = curr.Next
			}
		}
	}

	if pl == nil {
		return rln
	}

	return head
}

func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	var lr []*ListNode
	curr := head
	var pl, nr *ListNode
	i := 1
	for curr != nil {
		if i >= left && i <= right {
			lr = append([]*ListNode{curr}, lr...)
		} else if i < left {
			pl = curr
		} else if i > right && nr == nil {
			nr = curr
		}
		curr = curr.Next
		i++
	}

	for i, v := range lr {
		if i == 0 {
			if pl != nil {
				pl.Next = v
			}
			if len(lr) > 1 {
				v.Next = lr[i+1]
			}
		} else if i == len(lr)-1 {
			v.Next = nr
		} else {
			v.Next = lr[i+1]
		}
	}
	if pl == nil && len(lr) > 0 {
		return lr[0]
	}

	return head
}
