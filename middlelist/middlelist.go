package middlelist

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	l := head

	var arr [100]*ListNode
	i := 0
	for l != nil {
		arr[i] = l
		i++
		l = l.Next
	}

	return arr[i/2]
}

func middleNode_leet0ms(head *ListNode) *ListNode {

	arr := []*ListNode{head}

	for {
		if arr[len(arr)-1].Next != nil {
			arr = append(arr, arr[len(arr)-1].Next)
		} else {
			break
		}
	}

	return arr[len(arr)/2]
}
