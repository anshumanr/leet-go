package maxtwinsumlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result int

var tests = []struct {
	input    *ListNode
	expected int
}{
	{
		&ListNode{2, &ListNode{1, &ListNode{3, &ListNode{5, nil}}}},
		7,
	},
	{
		&ListNode{2, &ListNode{1, nil}},
		3,
	},
	{
		&ListNode{1, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}}},
		9,
	},
	{
		&ListNode{1, &ListNode{31, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{1, &ListNode{2, &ListNode{23, &ListNode{4, &ListNode{5, &ListNode{15, &ListNode{4, &ListNode{4, &ListNode{53, &ListNode{1, &ListNode{143, &ListNode{1, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{5, &ListNode{4, &ListNode{4, &ListNode{3, &ListNode{1, &ListNode{143, &ListNode{1, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{5, &ListNode{4, &ListNode{4, &ListNode{3, &ListNode{1, &ListNode{143, nil}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}},
		144,
	},
}

func TestPairSum(t *testing.T) {
	for i, test := range tests {
		t.Log("test ", i+1)

		got := pairSum_leet154ms(test.input)
		assert.Equal(t, test.expected, got)

		t.Log("got: ", got, ", expected: ", test.expected)
	}
}

func BenchmarkPairSum(b *testing.B) {
	funcs := []struct {
		name string
		f    func(list *ListNode) int
	}{
		{"mine", pairSum},
		{"leet-fastest", pairSum_leet154ms},
	}

	var res int
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
