package parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result int

var tests = []struct {
	input    string
	output   string
	expected int
}{
	{
		"(()",
		"()",
		2,
	},
	{
		")()())",
		"()()",
		4,
	},
	{
		")()()))()()))()()))()()))()()))()())",
		"()()",
		4,
	},
	{
		"",
		"",
		0,
	},
	{
		"))((()))()(())(",
		"((()))()(())",
		12,
	},
	{
		"()(())",
		"()(())",
		6,
	},
	{
		"()(()",
		"()",
		2,
	},
}

func TestLongestValidParentheses(t *testing.T) {
	for i, test := range tests {
		t.Log("test ", i+1, test.input)

		got := longestValidParentheses(test.input)
		assert.Equal(t, test.expected, got)

		t.Log("got: ", got, ", expected: ", test.expected)
	}
}

func TestLongestValidParentheses2(t *testing.T) {
	for i, test := range tests {
		t.Log("test ", i+1)

		got := longestValidParentheses2(test.input)
		assert.Equal(t, test.expected, got)

		t.Log("got: ", got, ", expected: ", test.expected)
	}
}

func BenchmarkLongestValidParentheses(b *testing.B) {
	funcs := []struct {
		name string
		f    func(s string) int
	}{
		{"mine", longestValidParentheses},
		{"mine2", longestValidParentheses2},
		{"leetcode", longestValidParentheses_leet},
	}

	var res int
	for _, fun := range funcs {

		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(tests[k%2].input)
			}

		})
	}

	result = res
}
