package middlelist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var result *ListNode

var (
	list1 = &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
	list2 = &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	list3 = &ListNode{1, &ListNode{1, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}
)

var tests = []struct {
	input    *ListNode
	expected *ListNode
}{
	{
		list1,
		&ListNode{4, &ListNode{5, nil}},
	},
	{
		list2,
		&ListNode{3, &ListNode{4, nil}},
	},
	{
		list3,
		&ListNode{4, &ListNode{4, &ListNode{5, nil}}},
	},
	{
		&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		&ListNode{3, &ListNode{4, &ListNode{5, nil}}},
	},
	{
		&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
		&ListNode{2, &ListNode{4, nil}},
	},
	{
		&ListNode{1, &ListNode{2, nil}},
		&ListNode{2, nil},
	},
	{
		&ListNode{1, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}}},
		&ListNode{5, &ListNode{6, &ListNode{7, nil}}},
	},
	{
		nil,
		nil,
	},
	{
		&ListNode{0, nil},
		&ListNode{0, nil},
	},
	{
		&ListNode{1, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{5, &ListNode{4, &ListNode{4, &ListNode{3, &ListNode{1, &ListNode{143, nil}}}}}}}}}}}}}}}}},
		&ListNode{3, &ListNode{4, &ListNode{5, &ListNode{5, &ListNode{4, &ListNode{4, &ListNode{3, &ListNode{1, &ListNode{143, nil}}}}}}}}},
	},
}

func TestMiddleNode(t *testing.T) {
	for i, test := range tests {
		t.Log("test ", i+1)

		got := middleNode(test.input)
		assert.Equal(t, test.expected, got)
		for got != nil {
			fmt.Println("test ", i+1, got.Val)
			got = got.Next
		}

		t.Log("got: ", got, ", expected: ", test.expected)
	}
}

func BenchmarkMerge2Lists(b *testing.B) {
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	list3 := &ListNode{5, &ListNode{6, &ListNode{7, nil}}}
	list4 := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}}
	list5 := &ListNode{5, &ListNode{6, &ListNode{7, nil}}}
	list6 := &ListNode{1, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}}}

	lists := []*ListNode{list1, list2, list3, list4, list5, list6}

	funcs := []struct {
		name string
		f    func(list *ListNode) *ListNode
	}{
		{"mine", middleNode},
		{"leet-1ms", middleNode_leet0ms},
	}

	var res *ListNode
	for _, fun := range funcs {
		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(lists[k%6])
			}
		})
	}

	result = res
}
