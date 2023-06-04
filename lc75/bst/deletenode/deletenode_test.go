package deletenode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result *TreeNode

const null = -10000000000

var tests = []struct {
	root     *TreeNode
	key      int
	expected *TreeNode
}{
	{
		createTree([]int64{5, 3, 6, 2, 4, null, 7}),
		3,
		createTree([]int64{5, 4, 6, 2, null, null, 7}),
	},
	{
		createTree([]int64{5, 3, 6, 2, 4, null, 7}),
		0,
		createTree([]int64{5, 3, 6, 2, 4, null, 7}),
	},
	{
		createTree([]int64{}),
		0,
		nil,
	},
	{
		createTree([]int64{1}),
		1,
		nil,
	},
	{
		createTree([]int64{1, null, 2, null, 3, null, 4, null, 5}),
		3,
		createTree([]int64{1, null, 2, null, 4, null, 5}),
	},
	{
		createTree([]int64{2, 0, 33, null, 1, 25, 40, null, null, 11, 31, 34, 45, 10, 18, 29, 32, null, 36, 43, 46, 4, null, 12, 24, 26, 30, null, null, 35, 39, 42, 44, null, 48, 3, 9, null, 14, 22, null, null, 27, null, null, null, null, 38, null, 41, null, null, null, 47, 49, null, null, 5, null, 13, 15, 21, 23, null, 28, 37, null, null, null, null, null, null, null, null, 8, null, null, null, 17, 19, null, null, null, null, null, null, null, 7, null, 16, null, null, 20, 6}),
		48,
		createTree([]int64{2, 0, 33, null, 1, 25, 40, null, null, 11, 31, 34, 45, 10, 18, 29, 32, null, 36, 43, 46, 4, null, 12, 24, 26, 30, null, null, 35, 39, 42, 44, null, 49, 3, 9, null, 14, 22, null, null, 27, null, null, null, null, 38, null, 41, null, null, null, 47, null, null, null, 5, null, 13, 15, 21, 23, null, 28, 37, null, null, null, null, null, null, 8, null, null, null, 17, 19, null, null, null, null, null, null, null, 7, null, 16, null, null, 20, 6}),
	},
	{
		createTree([]int64{27, 17, 33, 15, 26, 32, 45, 7, 16, 23, null, 31, null, 40, 49, 5, 14, null, null, 21, 25, 29, null, 34, 41, 48, null, 2, 6, 10, null, 20, 22, 24, null, 28, 30, null, 39, null, 42, 46, null, 1, 4, null, null, 9, 11, 18, null, null, null, null, null, null, null, null, null, 37, null, null, 44, null, 47, 0, null, 3, null, 8, null, null, 13, null, 19, 35, 38, 43, null, null, null, null, null, null, null, null, null, 12, null, null, null, null, 36}),
		49,
		createTree([]int64{27, 17, 33, 15, 26, 32, 45, 7, 16, 23, null, 31, null, 40, 48, 5, 14, null, null, 21, 25, 29, null, 34, 41, 46, null, 2, 6, 10, null, 20, 22, 24, null, 28, 30, null, 39, null, 42, null, 47, 1, 4, null, null, 9, 11, 18, null, null, null, null, null, null, null, null, null, 37, null, null, 44, null, null, 0, null, 3, null, 8, null, null, 13, null, 19, 35, 38, 43, null, null, null, null, null, null, null, 12, null, null, null, null, 36}),
	},
}

func TestDeleteNode(t *testing.T) {

	for i, test := range tests {
		t.Log("test ", i+1)

		got := deleteNode(test.root, test.key)
		assert.Equal(t, test.expected, got)

		t.Log("got: ", got, ", expected: ", test.expected)
	}
}

func BenchmarkDeleteNode(b *testing.B) {
	funcs := []struct {
		name string
		f    func(*TreeNode, int) *TreeNode
	}{
		{"mine", deleteNode},
		{"leet-fastest", deleteNode_leet0ms},
	}

	var res *TreeNode
	for _, fun := range funcs {
		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			//b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(tests[k%len(tests)].root, tests[k%len(tests)].key)
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
