package reverselist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	l := head

	var node *ListNode
	for l != nil {
		node = l.Next
		if node != nil {
			l.Next = node.Next
			node.Next = head
			head = node
		} else {
			break
		}
	}

	return head
}

func reverseList_leet0ms(head *ListNode) *ListNode {
	// get slice of values
	v := []int{}
	for head != nil {
		v = append(v, head.Val)
		head = head.Next
	}

	// reverse the slice
	for i := 0; i < len(v)/2; i++ {
		j := len(v) - 1 - i
		v[i], v[j] = v[j], v[i]
	}
	//fmt.Println(v)

	// create a LinkedList to be the result
	headResult := &ListNode{}
	curr := headResult
	for i := 0; i < len(v); i++ {
		curr.Next = &ListNode{
			Val: v[i],
		}
		curr = curr.Next
	}

	return headResult.Next
}

func reverseList_leet1ms(head *ListNode) *ListNode {
	var h *ListNode
	for head != nil {
		c := head
		head = head.Next
		c.Next = h
		h = c
	}
	return h
}
