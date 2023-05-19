package oddevenlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	oddT, evenHdr, even := head, head.Next, head.Next

	for even != nil && even.Next != nil {
		oddT.Next = even.Next
		even.Next, oddT.Next.Next = even.Next.Next, evenHdr
		oddT, even = oddT.Next, even.Next
	}
	return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList_leet1ms(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	oddCur := head
	evenHead := head.Next
	evenCur := evenHead

	for oddCur.Next != nil && evenCur.Next != nil {
		oddCur.Next = oddCur.Next.Next
		oddCur = oddCur.Next

		evenCur.Next = evenCur.Next.Next
		evenCur = evenCur.Next
	}

	oddCur.Next = evenHead

	return head

}
