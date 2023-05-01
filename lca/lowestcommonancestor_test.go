package lca

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result *TreeNode

var (
	list1 = &TreeNode{
		6,
		&TreeNode{
			2,
			&TreeNode{
				0, nil, nil,
			},
			&TreeNode{
				4,
				&TreeNode{
					3, nil, nil,
				},
				&TreeNode{
					5, nil, nil,
				},
			},
		},
		&TreeNode{
			8,
			&TreeNode{
				7, nil, nil,
			},
			&TreeNode{
				9, nil, nil,
			},
		},
	}
)

var tests = []struct {
	root     *TreeNode
	p        *TreeNode
	q        *TreeNode
	expected *TreeNode
}{
	{
		list1,
		&TreeNode{2, nil, nil},
		&TreeNode{8, nil, nil},
		&TreeNode{6, nil, nil},
	},
	{
		list1,
		&TreeNode{2, nil, nil},
		&TreeNode{4, nil, nil},
		&TreeNode{2, nil, nil},
	},
}

func TestLCA(t *testing.T) {

	for i, test := range tests {
		t.Log("test ", i+1)

		//high = math.MinInt32
		got := lowestCommonAncestor(test.root, test.p, test.q)
		assert.Equal(t, test.expected.Val, got.Val)

		t.Log("got: ", got, ", expected: ", test.expected)
	}
}

func BenchmarkLCA(b *testing.B) {
	funcs := []struct {
		name string
		f    func(list, p, q *TreeNode) *TreeNode
	}{
		{"mine", lowestCommonAncestor},
		{"leet-fastest", lowestCommonAncestor_leet6ms},
	}

	var res *TreeNode
	for _, fun := range funcs {
		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			//b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(tests[k%2].root, tests[k%2].p, tests[k%2].q)
			}
		})
	}

	result = res
}
