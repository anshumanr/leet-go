package levelorder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result [][]int

var (
	list1 = &TreeNode{
		3,
		&TreeNode{
			9, nil, nil,
		},
		&TreeNode{
			20, &TreeNode{
				15, nil, nil,
			},
			&TreeNode{
				7, nil, nil,
			},
		},
	}
	list2 = &TreeNode{
		1, nil, nil,
	}
)

var tests = []struct {
	input    *TreeNode
	expected [][]int
}{
	{
		list1,
		[][]int{{3}, {9, 20}, {15, 7}},
	},
	{
		list2,
		[][]int{{1}},
	},
}

func TestLevelOrder(t *testing.T) {

	for i, test := range tests {
		t.Log("test ", i+1)

		got := levelOrder(test.input)
		assert.Equal(t, test.expected, got)

		t.Log("got: ", got, ", expected: ", test.expected)
	}
}

func BenchmarkLevelOrder(b *testing.B) {
	funcs := []struct {
		name string
		f    func(list *TreeNode) [][]int
	}{
		{"mine", levelOrder},
		{"leet-1ms", levelOrder_leet1ms},
	}

	var res [][]int
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
