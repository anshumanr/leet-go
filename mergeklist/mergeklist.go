package mergeklist

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {

	arr := []int{}

	for _, l := range lists {
		for l != nil {
			arr = append(arr, l.Val)
			l = l.Next
		}
	}

	sort.Ints(arr)

	return createList(arr)
}

func createList(arr []int) *ListNode {
	var (
		head *ListNode
		prev *ListNode
	)
	sort.Ints(arr)

	for _, v := range arr {
		node := &ListNode{Val: v}
		if head == nil {
			head = node
		}
		if prev != nil {
			prev.Next = node
		}
		prev = node
	}

	return head
}

func mergeKLists0ms(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	mid := len(lists) / 2

	left := mergeKLists(lists[0:mid])
	right := mergeKLists(lists[mid:])

	return merge(left, right)
}

func merge(l1, l2 *ListNode) *ListNode {
	c1 := l1
	c2 := l2

	dummy := &ListNode{0, nil}
	cur := dummy
	for c1 != nil || c2 != nil {
		var c *ListNode

		if c1 != nil && c2 != nil {
			if c1.Val < c2.Val {
				c = c1
				c1 = c1.Next
			} else {
				c = c2
				c2 = c2.Next
			}
		} else {
			if c1 == nil {
				c = c2
				c2 = c2.Next
			} else {
				c = c1
				c1 = c1.Next
			}
		}

		cur.Next = c
		cur = cur.Next
	}

	return dummy.Next
}

func mergeKLists5ms(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	nodes := make([]*ListNode, 0, 10^4)
	for i := range lists {
		l := lists[i]
		for l != nil {
			nodes = append(nodes, l)
			l = l.Next
		}
	}

	if len(nodes) == 0 {
		return nil
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Val < nodes[j].Val
	})

	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}

	nodes[len(nodes)-1].Next = nil
	return nodes[0]
}
