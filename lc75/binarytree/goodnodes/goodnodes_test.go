package goodnodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result int

var (
	list1 = &TreeNode{
		6,
		&TreeNode{
			2,
			&TreeNode{0, nil, nil},
			&TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}},
		},
		&TreeNode{
			8,
			&TreeNode{7, nil, nil},
			&TreeNode{9, nil, nil},
		},
	}
	list2 = &TreeNode{
		3,
		&TreeNode{
			5,
			&TreeNode{6, nil, nil},
			&TreeNode{2, &TreeNode{7, nil, nil}, &TreeNode{4, nil, nil}},
		},
		&TreeNode{
			1, &TreeNode{0, nil, nil}, &TreeNode{8, nil, nil}},
	}
	list3 = &TreeNode{
		3,
		&TreeNode{1, &TreeNode{3, nil, nil}, nil},
		&TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{5, nil, nil}},
	}
	list4 = &TreeNode{
		3,
		&TreeNode{3, &TreeNode{4, nil, nil}, &TreeNode{2, nil, nil}},
		nil,
	}
	list5 = &TreeNode{
		6,
		&TreeNode{
			2,
			&TreeNode{0, nil, nil},
			&TreeNode{4, &TreeNode{3, nil, &TreeNode{9, nil, nil}}, &TreeNode{5, &TreeNode{6, nil, nil}, nil}},
		},
		&TreeNode{
			8,
			&TreeNode{7, &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}}},
			&TreeNode{9, nil, nil},
		},
	}
)

var tests = []struct {
	root     *TreeNode
	expected int
}{
	{
		list3,
		4,
	},
	{
		list4,
		3,
	},
	{
		list1,
		3,
	},
	{
		list2,
		5,
	},
	{
		list5,
		5,
	},
}

func TestGoodNodes(t *testing.T) {

	for i, test := range tests {
		t.Log("test ", i+1)

		//high = math.MinInt32
		got := goodNodes(test.root)
		assert.Equal(t, test.expected, got)

		t.Log("got: ", got, ", expected: ", test.expected)
	}
}

func BenchmarkGoodNodes(b *testing.B) {
	funcs := []struct {
		name string
		f    func(list *TreeNode) int
	}{
		{"mine", goodNodes},
		{"leet-fastest", goodNodes_leet65ms},
	}

	var res int
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
