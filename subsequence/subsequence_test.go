package subsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result bool

var tests = []struct {
	input1   string
	input2   string
	expected bool
}{
	{
		"abc",
		"ahbgdc",
		true,
	},
	{
		"axc",
		"ahbgdc",
		false,
	},
	{
		"acxw",
		"ahbgdcsdfcxvdfewsdf",
		true,
	},
	{
		"axc",
		"axc",
		true,
	},
	{
		"acxw",
		"ahbxgdcsdfcxvdfewsdf",
		true,
	},
	{
		"aaaaaaaa",
		"ahbxgadcsdfacxvadfewasdfaaa",
		true,
	},
	{
		"",
		"",
		true,
	},
	{
		"",
		"lkhsfikojewpoj",
		true,
	},
	{
		"acb",
		"ahbgdc",
		false,
	},
}

func TestIsomorphic(t *testing.T) {
	for i, test := range tests {
		t.Log("test ", i+1)

		got := isSubsequence(test.input1, test.input2)
		assert.Equal(t, test.expected, got)

		t.Log("got: ", got, ", expected: ", test.expected)

	}
}

func BenchmarkSubsequence(b *testing.B) {
	funcs := []struct {
		name string
		f    func(s, t string) bool
	}{
		{"mine", isSubsequence},
		{"leetcode-0ms", isSubsequence_leet0ms},
	}

	var res bool
	for _, fun := range funcs {

		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(tests[k%2].input1, tests[k%2].input2)
			}

		})
	}

	result = res
}
