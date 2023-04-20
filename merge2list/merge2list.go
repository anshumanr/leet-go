package merge2list

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	l1, l2 := list1, list2

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var prevl1 *ListNode
	for l1 != nil {
		if l2 == nil {
			return list1
		}

		if l1.Val <= l2.Val {
			prevl1 = l1
			l1 = l1.Next
		} else {
			node := &ListNode{Val: l2.Val, Next: l1}
			if prevl1 != nil {
				prevl1.Next = node
			} else {
				list1 = node
			}
			prevl1 = node
			l2 = l2.Next
		}
	}

	if l2 != nil {
		prevl1.Next = l2
	}

	return list1
}

func mergeTwoLists_leet1ms(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 == nil {
		current.Next = list2
	} else {
		current.Next = list1
	}

	return result.Next
}
