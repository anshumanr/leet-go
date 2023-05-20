package maxtwinsumlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	var arr []int
	max, cnt := 0, 0

	l := head
	for l != nil {
		arr, l = append(arr, l.Val), l.Next
		cnt++
	}

	i := 0
	for ; i < cnt/2; i++ {
		if max < arr[i]+arr[cnt-1-i] {
			max = arr[i] + arr[cnt-1-i]
		}
	}

	return max
}

func pairSum_leet154ms(head *ListNode) int {
	slow, fast := head, head
	prev := &ListNode{}
	res := 0
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		tmp := slow.Next
		slow.Next = prev
		prev = slow
		slow = tmp
	}
	for slow != nil {
		res = max(res, slow.Val+prev.Val)
		slow = slow.Next
		prev = prev.Next
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
