package narypreorder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result []int

var (
	list1 = &Node{
		1, []*Node{
			{3, []*Node{
				{5, nil},
				{6, nil},
			}},
			{2, nil},
			{4, nil},
		},
	}
	list2 = &Node{
		1, []*Node{
			{2, nil},
			{3, []*Node{
				{6, nil},
				{7, []*Node{
					{11, []*Node{
						{14, nil},
					}},
				}},
			}},
			{4, []*Node{
				{8, []*Node{
					{12, nil},
				}},
			}},
			{5, []*Node{
				{9, []*Node{
					{13, nil},
				}},
				{10, nil},
			}},
		},
	}
)

var tests = []struct {
	input    *Node
	expected []int
}{
	{
		list1,
		[]int{1, 3, 5, 6, 2, 4},
	},
	{
		list2,
		[]int{1, 2, 3, 6, 7, 11, 14, 4, 8, 12, 5, 9, 13, 10},
	},
}

func TestNaryPreorder(t *testing.T) {

	for i, test := range tests {
		t.Log("test ", i+1)

		got := preorder_iter(test.input)
		assert.Equal(t, test.expected, got)

		t.Log("got: ", got, ", expected: ", test.expected)
	}
}

func BenchmarkNaryPreorder(b *testing.B) {
	funcs := []struct {
		name string
		f    func(list *Node) []int
	}{
		{"mine", preorder},
		{"mine-iter", preorder_iter},
		{"leet-0ms", preorder_leet0ms},
	}

	var res []int
	for _, fun := range funcs {
		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			//b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(tests[k%2].input)
			}
		})
	}

	result = res
}
