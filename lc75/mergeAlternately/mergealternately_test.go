package mergealternately

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result string

var tests = []struct {
	input1   string
	input2   string
	expected string
}{
	{"abc", "pqr", "apbqcr"},
	{"ab", "pqrs", "apbqrs"},
	{"abcd", "pq", "apbqcd"},
}

func TestMergeAlternately(t *testing.T) {

	for i, v := range tests {
		t.Log("test: ", i+1)
		got := mergeAlternately(v.input1, v.input2)
		assert.Equal(t, v.expected, got)
	}
}

func BenchmarkMergeAlternately(b *testing.B) {
	funcs := []struct {
		name string
		f    func(s1, s2 string) string
	}{
		{"mine", mergeAlternately},
		{"leet-fastest", mergeAlternately_leet0ms},
	}

	var res string
	for _, fun := range funcs {
		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(tests[k%len(tests)].input1, tests[k%len(tests)].input2)
			}

		})
	}

	result = res
}
