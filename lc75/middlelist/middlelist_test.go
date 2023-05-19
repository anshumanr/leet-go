package middlelist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var result *ListNode

var tests = []struct {
	input    *ListNode
	expected *ListNode
}{
	{
		&ListNode{1, &ListNode{3, &ListNode{4, &ListNode{7, &ListNode{1, &ListNode{2, &ListNode{6, nil}}}}}}},
		&ListNode{1, &ListNode{3, &ListNode{4, &ListNode{1, &ListNode{2, &ListNode{6, nil}}}}}},
	},
	{
		&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
		&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
	},
	{
		&ListNode{2, &ListNode{1, nil}},
		&ListNode{2, nil},
	},
	{
		&ListNode{1, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}}},
		&ListNode{1, &ListNode{3, &ListNode{4, &ListNode{6, &ListNode{7, nil}}}}},
	},
	{
		&ListNode{0, nil},
		nil,
	},
	{
		&ListNode{1, &ListNode{31, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{1, &ListNode{2, &ListNode{23, &ListNode{4, &ListNode{5, &ListNode{15, &ListNode{4, &ListNode{4, &ListNode{53, &ListNode{1, &ListNode{143, &ListNode{1, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{5, &ListNode{4, &ListNode{4, &ListNode{3, &ListNode{1, &ListNode{143, &ListNode{1, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{5, &ListNode{4, &ListNode{4, &ListNode{3, &ListNode{1, &ListNode{143, nil}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}},
		&ListNode{1, &ListNode{31, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{1, &ListNode{2, &ListNode{23, &ListNode{4, &ListNode{5, &ListNode{15, &ListNode{4, &ListNode{4, &ListNode{53, &ListNode{1, &ListNode{143, &ListNode{1, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{1, &ListNode{2, &ListNode{4, &ListNode{5, &ListNode{5, &ListNode{4, &ListNode{4, &ListNode{3, &ListNode{1, &ListNode{143, &ListNode{1, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{5, &ListNode{4, &ListNode{4, &ListNode{3, &ListNode{1, &ListNode{143, nil}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}},
	},
}

func TestReverseList(t *testing.T) {
	for i, test := range tests {
		t.Log("test ", i+1)

		got := deleteMiddle(test.input)
		assert.Equal(t, test.expected, got)
		for got != nil {
			fmt.Println("test ", i+1, got.Val)
			got = got.Next
		}

		t.Log("got: ", got, ", expected: ", test.expected)
	}
}

func BenchmarkMerge2Lists(b *testing.B) {
	funcs := []struct {
		name string
		f    func(list *ListNode) *ListNode
	}{
		{"mine", deleteMiddle},
		{"mine-V1", deleteMiddle_BF},
		{"leet-fastest", deleteMiddle_leet274ms},
	}

	var res *ListNode
	for _, fun := range funcs {
		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(tests[k%len(tests)].input)
			}
		})
	}

	result = res
}
