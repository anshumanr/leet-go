package exitmaze

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result int

var tests = []struct {
	maze     [][]byte
	entrance []int
	expected int
}{
	{
		[][]byte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}},
		[]int{1, 2},
		1,
	},
	{
		[][]byte{{'+', '+', '+'}, {'.', '.', '.'}, {'+', '+', '+'}},
		[]int{1, 0},
		2,
	},
	{
		[][]byte{{'.', '+'}},
		[]int{0, 0},
		-1,
	},
	{
		[][]byte{{'+', '.', '+', '+', '+', '+', '+'}, {'+', '.', '+', '.', '.', '.', '+'}, {'+', '.', '+', '.', '+', '.', '+'}, {'+', '.', '.', '.', '+', '.', '+'}, {'+', '+', '+', '+', '+', '+', '.'}},
		[]int{0, 1},
		-1,
	},
}

func TestNearestExit(t *testing.T) {
	for i, v := range tests {
		t.Log("test: ", i+1)
		got := nearestExit(v.maze, v.entrance)
		assert.Equal(t, v.expected, got)
	}
}

func BenchmarkNearestExit(b *testing.B) {
	funcs := []struct {
		name string
		f    func([][]byte, []int) int
	}{
		{"mine", nearestExit},
		{"mine-v2", nearestExitV2},
		{"leet-fastest", nearestExit_leet6ms},
	}

	var res int
	for _, fun := range funcs {

		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(tests[k%len(tests)].maze, tests[k%len(tests)].entrance)
			}

		})
	}

	result = res
}
