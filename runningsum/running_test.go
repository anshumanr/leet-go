package runningsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result []int

var tests = []struct {
	nums         []int
	expectedNums []int
}{
	{
		[]int{1, 2, 3, 4},
		[]int{1, 3, 6, 10},
	},
	{
		[]int{-1},
		[]int{-1},
	},
	{
		[]int{3, 1, 2, 10, 1},
		[]int{3, 4, 6, 16, 17},
	},
	{
		[]int{-1, 3, -8, 12, 2, 9, -5, 3, 1, 2, 10, 1},
		[]int{-1, 2, -6, 6, 8, 17, 12, 15, 16, 18, 28, 29},
	},
}

func TestRunningSum(t *testing.T) {
	for i, test := range tests {
		t.Log("test ", i+1)

		got := runningSum(test.nums)
		assert.Equal(t, len(test.nums), len(got))

		t.Log("got: ", got, ", expected: ", test.nums)

		for i, val := range test.expectedNums {
			assert.Equal(t, val, got[i])
		}
	}
}

func TestRunningSumV2(t *testing.T) {
	for i, test := range tests {
		t.Log("test ", i+1)

		got := runningSumV2(test.nums)
		assert.Equal(t, len(test.nums), len(got))

		t.Log("got: ", got, ", expected: ", test.nums)

		for i, val := range test.expectedNums {
			assert.Equal(t, val, test.expectedNums[i])
		}
	}
}

func BenchmarkRunningSums(b *testing.B) {
	funcs := []struct {
		name string
		f    func(nums []int) []int
	}{
		{"mine", runningSum},
		{"leet-0ms", runningSumV2},
	}

	var res []int
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
