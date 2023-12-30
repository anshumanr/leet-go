package largestKelem

import "container/heap"

func findKthLargest(nums []int, k int) int {
	l := len(nums)
	q := make([]int, k)

	for i := range q {
		q[i] = -100000
	}

	i := 0
	for ; i < l; i++ {
		if nums[i] <= q[k-1] {
			continue
		}

		q[k-1] = nums[i]
		for j := k - 1; j > 0; j-- {
			if q[j] > q[j-1] {
				q[j], q[j-1] = q[j-1], q[j]
			}
		}

	}
	return q[k-1]
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargestV2(nums []int, k int) int {
	l := len(nums)
	h := &IntHeap{}
	heap.Init(h)

	for _, v := range nums {
		heap.Push(h, v)
	}

	for i := 0; i < l-k; i++ {
		heap.Pop(h)
	}

	a := heap.Pop(h).(int)
	return a
}

func findKthLargest_leet63ms(nums []int, k int) int {
	count := make([]int, 20001)
	for _, num := range nums {
		count[num+10000]++
	}
	running := 0
	for i := len(count) - 1; i >= 0; i-- {
		if count[i] > 0 {
			running += count[i]
			if running >= k {
				return i - 10000
			}
		}
	}
	return -1
}
