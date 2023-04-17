package pivotindex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result int

var tests = []struct {
	nums       []int
	runningSum []int
	pivot      int
}{
	{
		[]int{1, 7, 3, 6, 5, 6},
		[]int{1, 8, 11, 17, 22, 28},
		3,
	},
	{
		[]int{2, 1, -1},
		[]int{2, 3, 2},
		0,
	},
	{
		[]int{1, 2, 3},
		[]int{1, 3, 5},
		-1,
	},
}

func TestPivotIndex(t *testing.T) {
	for i, test := range tests {
		t.Log("test ", i+1)

		got := pivotIndex(test.nums)
		assert.Equal(t, got, test.pivot)

		t.Log("got: ", got, ", expected: ", test.pivot)

	}
}

func BenchmarkPivotIndex(b *testing.B) {
	funcs := []struct {
		name string
		f    func(nums []int) int
	}{
		{"mine", pivotIndex},
		{"leet-code", pivotIndexV2},
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
