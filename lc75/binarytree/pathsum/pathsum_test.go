package pathsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result int

const null = -10000000000

var (
	list1 = []int64{10, 5, -3, 3, 2, null, 11, 3, -2, null, 1}
	list2 = []int64{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1}
	list3 = []int64{3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 3, 5, -3, 3, 2, null, 11, -3, -2, null, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1, 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1, 10, 5, -3, 3, 2, null, 11, 3, -2, null, 1}
)

var tests = []struct {
	root     *TreeNode
	target   int
	expected int
}{
	{
		createTree(list1),
		8,
		3,
	},
	{
		createTree(list2),
		22,
		3,
	},
	{
		createTree([]int64{}),
		0,
		0,
	},
	{
		createTree([]int64{1}),
		1,
		1,
	},
	{
		createTree([]int64{0, 1, 1}),
		1,
		4,
	},
	{
		createTree([]int64{1, null, 2, null, 3, null, 4, null, 5}),
		3,
		2,
	},
	{
		&TreeNode{
			1, nil, &TreeNode{2, nil, &TreeNode{3, nil, &TreeNode{4, nil, &TreeNode{5, nil, nil}}}},
		},
		3,
		2,
	},
	{
		createTree(list3),
		22,
		571,
	},
}

func TestGoodNodes(t *testing.T) {

	for i, test := range tests {
		t.Log("test ", i+1)

		got := pathSum(test.root, test.target)
		assert.Equal(t, test.expected, got)

		t.Log("got: ", got, ", expected: ", test.expected)
	}
}

func BenchmarkGoodNodes(b *testing.B) {
	funcs := []struct {
		name string
		f    func(*TreeNode, int) int
	}{
		{"mine", pathSum},
		{"leet-fastest", pathSum_leet0ms},
	}

	var res int
	for _, fun := range funcs {
		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			//b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(tests[k%len(tests)].root, tests[k%len(tests)].target)
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
