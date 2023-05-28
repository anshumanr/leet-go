package rightsideview

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result []int

const null = -10000000000

var (
	list1 = []int64{1, 2, 3, null, 5, null, 4}
	list2 = []int64{1, 1, 1, null, 1, null, null, 1, 1, null, 1}
	list3 = []int64{3, 2, null, 11, null, 1, 5, 3, 2, null, 11, 3, null, null, 11, null, 1, 3, 5, 3, 2, null, 11, null, 1, 10, 5, 3, 2, null, 11, 3, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, 3, 2, null, 11, 3, null, 1, 3, 5, 3, 2, null, 11, null, 1, 3, 5, 3, 2, null, 11, null, 1, 10, 5, 3, 2, null, 11, 3, null, 1, 3, 5, 3, 2, null, 11, null, 1, 3, 5, 3, 2, null, 11, null, 1, 10, 5, 3, 2, null, 11, 3, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, 3, 2, null, 11, 3, null, 1, 3, 5, 3, 2, null}
)

var tests = []struct {
	root     *TreeNode
	expected []int
}{
	{
		createTree(list1),
		[]int{1, 3, 4},
	},
	{
		createTree([]int64{1, null, 3}),
		[]int{1, 3},
	},
	{
		createTree([]int64{}),
		[]int{},
	},
	{
		createTree([]int64{1}),
		[]int{1},
	},
	{
		createTree([]int64{1, null, 2, null, 3, null, 4, null, 5}),
		[]int{1, 2, 3, 4, 5},
	},
	{
		createTree([]int64{6, 9, 7, 3, null, 2, 8, 5, 8, 9, 7, 3, 9, 9, 4, 2, 10, null, 5, 4, 3, 10, 10, 9, 4, 1, 2, null, null, 6, 5, null, null, null, null, 9, null, 9, 6, 5, null, 5, null, null, 7, 7, 4, null, 1, null, null, 3, 7, null, 9, null, null, null, null, null, null, null, null, 9, 9, null, null, null, 7, null, null, null, null, null, null, null, null, null, 6, 8, 7, null, null, null, 3, 10, null, null, null, null, null, 1, null, 1, 2}),
		[]int{6, 7, 8, 9, 4, 4, 7, 3, 1, 2},
	},
	{
		createTree(list3),
		[]int{3, 2, 11, 5, 11, 1, 11, 3, 1, 2, 11, 2},
	},
}

func TestRightSideView(t *testing.T) {

	for i, test := range tests {
		t.Log("test ", i+1)

		got := rightSideView(test.root)
		assert.Equal(t, test.expected, got)

		t.Log("got: ", got, ", expected: ", test.expected)
	}
}

func BenchmarkRightSideView(b *testing.B) {
	funcs := []struct {
		name string
		f    func(*TreeNode) []int
	}{
		{"mine", rightSideView},
		{"leet-fastest", rightSideView_leet1ms},
	}

	var res []int
	for _, fun := range funcs {
		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			//b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(tests[k%len(tests)].root)
			}
		})
	}

	result = res
}

func createTree(arr []int64) *TreeNode {
	var (
		head *TreeNode
	)

	l := len(arr)
	if l == 0 {
		return nil
	}
	tm := make(map[int]*TreeNode, l)
	child := 1
	for i := 0; i < l; i++ {
		if arr[i] == null {
			continue
		}

		node, ok := tm[i]
		if !ok {
			node = &TreeNode{Val: int(arr[i])}
			tm[i] = node
		}

		if child < l {
			if arr[child] > null {
				_, ok := tm[child]
				if !ok {
					node.Left = &TreeNode{Val: int(arr[child])}
					tm[child] = node.Left
				}
			}
			child++
		}
		if child < l {
			if arr[child] > null {
				_, ok := tm[child]
				if !ok {
					node.Right = &TreeNode{Val: int(arr[child])}
					tm[child] = node.Right
				}
			}
			child++
		}

		if head == nil {
			head = node
		}
	}

	return head
}
