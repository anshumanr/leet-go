package mergeklist

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var result *ListNode

func TestMergeKLists(t *testing.T) {
	list1 := &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	list3 := &ListNode{2, &ListNode{6, nil}}

	lists := []*ListNode{list1, list2, list3}

	t.Log("name: ")
	got := mergeKLists(lists)
	for got != nil {
		fmt.Println(got.Val)
		got = got.Next
	}
}

func BenchmarkMergeKLists(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	list1 := createList(r.Perm(500))
	list2 := createList(r.Perm(300))
	list3 := createList([]int{9, 8, 6, 5, 3, 2, 1, 1, 2, 43, 56, 43, 12, 23, 54, 67, 879, 52, 412, 312, 124, 523,
		43, 3, 856, 578, 643, 352, 214, 12, -12, -54, -1289, -236})

	lists := []*ListNode{list1, list2, list3}

	funcs := []struct {
		name string
		f    func(lists []*ListNode) *ListNode
	}{
		{"myV1", mergeKLists},
		{"leet-0ms", mergeKLists0ms},
		{"leet-5ms", mergeKLists5ms},
	}

	var res *ListNode
	for _, fun := range funcs {
		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(lists)
			}
		})
	}

	result = res
}
