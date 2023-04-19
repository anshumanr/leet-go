package isomorphic

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
		"egg",
		"add",
		true,
	},
	{
		"foo",
		"bar",
		false,
	},
	{
		"papet",
		"title",
		true,
	},
}

func TestIsomorphic(t *testing.T) {
	for i, test := range tests {
		t.Log("test ", i+1)

		got := isIsomorphic(test.input1, test.input2)
		assert.Equal(t, test.expected, got)

		t.Log("got: ", got, ", expected: ", test.expected)

	}
}

func TestIsomorphicV2(t *testing.T) {
	for i, test := range tests {
		t.Log("test ", i+1)

		got := isIsomorphicV2(test.input1, test.input2)
		assert.Equal(t, test.expected, got)

		t.Log("got: ", got, ", expected: ", test.expected)

	}
}

func BenchmarkIsomorphic(b *testing.B) {
	funcs := []struct {
		name string
		f    func(s, t string) bool
	}{
		{"mine", isIsomorphic},
		{"mine-v2", isIsomorphicV2},
		{"leetcode-1ms", isIsomorphic_leet1ms},
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
