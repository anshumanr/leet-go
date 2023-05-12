package highestaltitude

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result int

var tests = []struct {
	nums     []int
	expected int
}{
	{
		[]int{-5, 1, 5, 0, -7},
		1,
	},
	{
		[]int{-4, -3, -2, -1, 4, 3, 2},
		0,
	},
	{
		[]int{19, 18, 17, 16, 15, -14, 13, 12, 11, 10, 1, 2, 3, -4, 5, 6, 7, 8, 9, 10, -19, 18, 17, 16, 15, 14, -13, 12, 11, -10, 1, 2, 3, 4, 5, 6, -7, 8, 9, -10, 19, 18, -17, 16, 15, 14, 13, 12, -11, 10, 1, 2, -3, 4, 5, 6, 7, -8, 9, 10, -19, 18, -17, 16, 15, -14, 13, 12, 11, -10, 1, 2, 3},
		403,
	},
	{
		[]int{-19, 18, -17, 16, 15, -14, 13, 12, -11, 10, 1, 2, 3, -4, 5, 6, 7, 8, 9, 10, -19, 18, -17, -16, 15, 14, -13, 12, 11, -10, 1, 2, 3, 4, 5, 6, -7, 8, 9, -10, 19, 18, -17, 16, -15, 14, 13, 12, -11, 10, 1, 2, -3, 4, 5, 6, 7, -8, 9, 10, -19, 18, -17, 16, 15, -14, 13, 12, 11, -10, 1, 2, 3, -17, 16, -15, 14, 13, 12, -11, 10, 1, 2, -3},
		234,
	},
	{
		[]int{-1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1},
		0,
	},
}

func TestMaxArea(t *testing.T) {
	for i, v := range tests {
		t.Log("test: ", i+1)
		got := largestAltitude(v.nums)
		assert.Equal(t, v.expected, got)
	}
}

func BenchmarkMaxArea(b *testing.B) {
	funcs := []struct {
		name string
		f    func([]int) int
	}{
		{"mine", largestAltitude},
		{"leet-fastest", largestAltitude_leet0ms},
	}

	var res int
	for _, fun := range funcs {

		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(tests[k%len(tests)].nums)
			}

		})
	}

	result = res
}
