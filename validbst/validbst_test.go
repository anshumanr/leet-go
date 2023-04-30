package validbst

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var result bool

var (
	list1 = &TreeNode{
		2,
		&TreeNode{
			1, nil, nil,
		},
		&TreeNode{
			3, nil, nil,
		},
	}
	list2 = &TreeNode{
		5,
		&TreeNode{
			1, nil, nil,
		},
		&TreeNode{
			4,
			&TreeNode{
				3, nil, nil,
			},
			&TreeNode{
				6, nil, nil,
			},
		},
	}
	list3 = &TreeNode{
		5,
		&TreeNode{
			4, nil, nil,
		},
		&TreeNode{
			6,
			&TreeNode{
				3, nil, nil,
			},
			&TreeNode{
				7, nil, nil,
			},
		},
	}
	list4 = &TreeNode{
		32,
		&TreeNode{
			26,
			&TreeNode{
				19,
				nil,
				&TreeNode{27, nil, nil},
			},
			nil,
		},
		&TreeNode{
			47,
			nil,
			&TreeNode{
				56, nil, nil,
			},
		},
	}
	list5 = &TreeNode{
		2, nil, nil,
	}
	list6 = &TreeNode{
		3,
		&TreeNode{
			1,
			&TreeNode{0, nil, nil},
			&TreeNode{2, nil, nil},
		},
		&TreeNode{
			5,
			&TreeNode{4, nil, nil},
			&TreeNode{6, nil, nil},
		},
	}
	list7 = &TreeNode{
		3,
		&TreeNode{
			1,
			&TreeNode{0, nil, nil},
			&TreeNode{
				2,
				nil,
				&TreeNode{3, nil, nil},
			},
		},
		&TreeNode{
			5,
			&TreeNode{4, nil, nil},
			&TreeNode{6, nil, nil},
		},
	}
)

var tests = []struct {
	input    *TreeNode
	expected bool
}{
	{
		list4,
		false,
	},
	{
		list3,
		false,
	},
	{
		list5,
		true,
	},
	{
		list1,
		true,
	},
	{
		list2,
		false,
	},
	{
		&TreeNode{0, nil, nil},
		true,
	},
	{
		list6,
		true,
	},
	{
		list7,
		false,
	},
	{
		&TreeNode{math.MinInt32, nil, nil},
		true,
	},
	{
		&TreeNode{
			math.MinInt32,
			&TreeNode{math.MinInt32, nil, nil},
			nil,
		},
		false,
	},
	{
		&TreeNode{
			math.MinInt32,
			nil,
			&TreeNode{math.MinInt32, nil, nil},
		},
		false,
	},
}

func TestValidBST(t *testing.T) {

	for i, test := range tests {
		t.Log("test ", i+1)

		//high = math.MinInt32
		got := isValidBST(test.input)
		assert.Equal(t, test.expected, got)

		t.Log("got: ", got, ", expected: ", test.expected)
	}
}

func BenchmarkLevelOrder(b *testing.B) {
	funcs := []struct {
		name string
		f    func(list *TreeNode) bool
	}{
		{"mine", isValidBST},
		{"leet-0ms", isValidBST_leet0ms},
	}

	var res bool
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
