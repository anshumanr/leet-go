package gcdofstrings

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
	{"ABCABC", "ABC", "ABC"},
	{"ABABAB", "ABAB", "AB"},
	{"BABABA", "ABAB", ""},
	{"LEET", "CODE", ""},
	{"TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXX"},
	{"TTTTTTTTTTTTTTTTTTTTTTTTTTTTTTT", "TTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTT", "T"},
}

func TestGcdOfStrings(t *testing.T) {

	for i, v := range tests {
		t.Log("test: ", i+1)
		got := gcdOfStringsV2(v.input1, v.input2)
		assert.Equal(t, v.expected, got)
	}
}

func BenchmarkGcdOfStrings(b *testing.B) {
	funcs := []struct {
		name string
		f    func(s1, s2 string) string
	}{
		{"mine", gcdOfStrings},
		{"mine-v2", gcdOfStringsV2},
		{"leet-fastest", gcdOfStrings_leet0ms},
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
