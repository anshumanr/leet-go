package merge2list

import (
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
	input1   *ListNode
	input2   *ListNode
	expected *ListNode
}{
	{
		list1,
		list2,
		list3,
	},
	{
		&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
		&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
		&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
	},
	{
		&ListNode{5, &ListNode{6, &ListNode{7, nil}}},
		&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
		&ListNode{1, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}}},
	},
	{
		nil,
		nil,
		nil,
	},
	{
		nil,
		&ListNode{0, nil},
		&ListNode{0, nil},
	},
}

func TestMerge2Lists(t *testing.T) {

	for i, test := range tests {
		t.Log("test ", i+1)

		got := mergeTwoLists(test.input1, test.input2)
		assert.Equal(t, test.expected, got)

		t.Log("got: ", got, ", expected: ", test.expected)
	}
}

func BenchmarkMerge2Lists(b *testing.B) {
	// list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	// list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	// list3 := &ListNode{5, &ListNode{6, &ListNode{7, nil}}}
	list4 := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}}
	//list5 := &ListNode{5, &ListNode{6, &ListNode{7, nil}}}
	list6 := &ListNode{1, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}}}

	//lists := []*ListNode{list1, list2, list3, list4, list5, list6}

	funcs := []struct {
		name string
		f    func(list1, list2 *ListNode) *ListNode
	}{
		{"mine", mergeTwoLists},
		{"leet-1ms", mergeTwoLists_leet1ms},
	}

	var res *ListNode
	for _, fun := range funcs {
		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			//b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(list4, list6)
			}
		})
	}

	result = res
}
