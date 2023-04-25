package maxprofit

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
		[]int{17, 1, 5, 3, 16, 4, 7, 5, 5, 3, 6, 9, 17, 10, 25, 3, 16, 4, 7, 2, 5, 3, 6, 11, 17, 11, 5, 3, 16, 4, 7, 2, 5, 32, 6, 1, 17, 1, 5, 3, 16, 4, 7, 1, 45, 3, 0},
		44,
	},
	{
		[]int{7, 1, 5, 3, 6, 4},
		5,
	},
	{
		[]int{17, 1, 5, 3, 16, 4, 7, 1, 5, 3, 6, 1},
		15,
	},
	{
		[]int{7, 6, 4, 3, 1},
		0,
	},
	{
		[]int{1, 2, 3},
		2,
	},
}

func TestMaxProfit(t *testing.T) {
	for i, test := range tests {
		t.Log("test ", i+1)

		got := maxProfit(test.nums)
		assert.Equal(t, got, test.expected)

		t.Log("got: ", got, ", expected: ", test.expected)

	}
}

func BenchmarkMaxProfit(b *testing.B) {
	funcs := []struct {
		name string
		f    func(nums []int) int
	}{
		{"mine", maxProfit},
		{"leet-code99ms", maxProfit_leet99ms},
	}

	var res int
	for _, fun := range funcs {

		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(tests[k%5].nums)
			}

		})
	}

	result = res
}
