package addtwonum

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbersV1(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		sum1 int
		sum2 int
	)

	mult := 1
	for l1 != nil {
		sum1 = sum1 + (l1.Val * mult)
		mult *= 10
		l1 = l1.Next
	}

	mult = 1
	for l2 != nil {
		sum2 = sum2 + l2.Val*mult
		mult *= 10
		l2 = l2.Next
	}

	sum := sum1 + sum2

	list := &ListNode{0, nil}

	iter := list
	for {
		iter.Val = sum % 10
		iter.Next = nil

		sum = sum / 10
		if sum > 0 {
			iter.Next = &ListNode{0, nil}
			iter = iter.Next
		} else {
			break
		}
	}

	return list
}

func addTwoNumbersV2(l1 *ListNode, l2 *ListNode) *ListNode {
	arr1 := listToArr(l1)
	arr2 := listToArr(l2)

	bigA := *arr1
	smallA := *arr2

	if len(*arr1) < len(*arr2) {
		bigA = *arr2
		smallA = *arr1
	}

	var r, s, q int
	ansArr := []int{}
	for i := 0; i < len(bigA); i++ {
		if i > len(smallA)-1 {
			s = bigA[i] + q
			r = s % 10
			q = s / 10
			ansArr = append(ansArr, r)
			continue
		}

		s = smallA[i] + bigA[i] + q
		r = s % 10
		q = s / 10

		ansArr = append(ansArr, r)
	}

	if q != 0 {
		ansArr = append(ansArr, q)
	}

	for j := len(ansArr) - 1; j >= 0; j-- {
		if ansArr[j] != 0 || j == 0 {
			ansArr = ansArr[:j+1]
			break
		}
	}

	return arrToList(ansArr)
}

func addTwoNumbersV3(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, prev, list *ListNode

	var r, s, q int
	lc1 := l1
	lc2 := l2

	for lc1 != nil || lc2 != nil {
		list = &ListNode{}
		if head == nil {
			head = list
		}

		if lc1 != nil && lc2 == nil {
			s = lc1.Val + q
			r = s % 10
			q = s / 10

			lc1 = lc1.Next

		} else if lc1 == nil && lc2 != nil {
			s = lc2.Val + q
			r = s % 10
			q = s / 10

			lc2 = lc2.Next
		} else {

			s = lc1.Val + lc2.Val + q
			r = s % 10
			q = s / 10

			lc1 = lc1.Next
			lc2 = lc2.Next
		}

		list.Val = r
		if prev != nil {
			prev.Next = list
		}
		prev = list
	}

	if q != 0 {
		list = &ListNode{q, nil}
		if prev != nil {
			prev.Next = list
		}
	}

	return head
}

func addTwoNumbers0ms(l1 *ListNode, l2 *ListNode) *ListNode {
	var headNode ListNode
	counter := &headNode
	c := 0
	l1c := l1
	l2c := l2
	for l1c != nil || l2c != nil || c != 0 {
		var a, b int
		if l1c == nil {
			a = 0
		} else {
			a = l1c.Val
			l1c = l1c.Next
		}
		if l2c == nil {
			b = 0
		} else {
			b = l2c.Val
			l2c = l2c.Next
		}
		tmp := c + a + b
		var newNode ListNode
		newNode.Val = tmp % 10
		c = (tmp - tmp%10) / 10
		counter.Next = &newNode
		counter = &newNode
	}
	return headNode.Next
}

func addTwoNumbers4ms(l1 *ListNode, l2 *ListNode) *ListNode {
	rest := 0

	t1 := l1
	t2 := l2
	var prev *ListNode = nil

	for t1 != nil && t2 != nil {
		temp := t1.Val + t2.Val + rest
		rest = 0
		if temp > 9 {
			temp = temp - 10
			rest = 1
		}

		t1.Val = temp

		prev = t1
		t1 = t1.Next
		t2 = t2.Next
	}

	if t2 != nil {
		prev.Next = t2
		t1 = t2
	}

	if rest > 0 {
		for t1 != nil {
			temp := t1.Val + rest
			rest = 0
			if temp > 9 {
				temp = temp - 10
				rest = 1
			}

			t1.Val = temp

			prev = t1
			t1 = t1.Next
		}

		if rest > 0 {
			prev.Next = &ListNode{Val: rest, Next: nil}
		}
	}

	return l1
}

func arrToList(arr []int) *ListNode {
	var list, prev *ListNode

	for _, n := range arr {
		node := &ListNode{n, nil}
		if list == nil {
			list = node
		}

		if prev != nil {
			prev.Next = node
		}
		prev = node
	}

	return list
}

func listToArr(list *ListNode) *[]int {
	var arr []int

	for list != nil {
		arr = append(arr, list.Val)
		list = list.Next
	}

	return &arr
}
