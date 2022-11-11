package remDup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result int

var tests = []struct {
	nums         []int
	expectedNums []int
	expectedLen  int
}{
	{
		[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		[]int{0, 1, 2, 3, 4},
		5,
	},
	{
		[]int{1, 1, 2},
		[]int{1, 2},
		2,
	},
	{
		[]int{-1, 1, 2},
		[]int{-1, 1, 2},
		3,
	},
	{
		[]int{-1},
		[]int{-1},
		1,
	},
}

func TestRemoveDuplicates(t *testing.T) {
	for i, test := range tests {
		t.Log("test ", i+1)

		got := removeDuplicates(test.nums)
		assert.Equal(t, test.expectedLen, got)

		for i, val := range test.expectedNums {
			assert.Equal(t, val, test.expectedNums[i])
		}
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {

	funcs := []struct {
		name string
		f    func(nums []int) int
	}{
		{"mine", removeDuplicates},
		{"leet-0ms", removeDuplicates0ms},
	}

	var res int
	for _, fun := range funcs {

		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(tests[k%2].nums)
			}

		})
	}

	result = res
}
