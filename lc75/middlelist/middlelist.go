package middlelist

type ListNode struct {
	Val  int
	Next *ListNode
}

//using fast pointer method used in LC solution
func deleteMiddle(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	if fast == nil {
		return nil
	}

	for fast != nil {
		if fast.Next == nil || fast.Next.Next == nil {
			break
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return head
}

func deleteMiddle_BF(head *ListNode) *ListNode {
	t, cnt := head, 0
	for t != nil {
		cnt++
		t = t.Next
	}

	t = head
	var prev *ListNode
	for i := 0; i < cnt/2; i++ {
		prev = t
		t = t.Next
	}

	if prev == nil {
		return prev
	}

	prev.Next = t.Next

	return head
}

func deleteMiddle_leet274ms(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}

	fast := head.Next.Next
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return head
}
