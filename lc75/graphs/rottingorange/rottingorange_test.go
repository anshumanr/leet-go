package rottingorange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result int

var tests = []struct {
	maze     [][]int
	expected int
}{
	{
		[][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
		4,
	},
	{
		[][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
		-1,
	},
	{
		[][]int{{0, 2}},
		0,
	},
	{
		[][]int{{1, 1, 1}, {1, 2, 0}, {0, 1, 1}},
		2,
	},
}

func TestOrangesRotting(t *testing.T) {
	for i, v := range tests {
		t.Log("test: ", i+1)
		got := orangesRotting(v.maze)
		assert.Equal(t, v.expected, got)
	}
}

func BenchmarkOrangesRotting(b *testing.B) {
	funcs := []struct {
		name string
		f    func([][]int) int
	}{
		{"mine", orangesRotting},
		{"leet-fastest", orangesRotting_0ms},
	}

	var res int
	for _, fun := range funcs {

		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(tests[k%len(tests)].maze)
			}

		})
	}

	result = res
}
