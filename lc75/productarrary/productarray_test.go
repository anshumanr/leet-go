package productarrary

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
		[]int{24, 12, 8, 6},
	},
	{
		[]int{-1, 1, 0, -3, 3},
		[]int{0, 0, 9, 0, 0},
	},
	{
		[]int{-1, 3, -8, 12, 2, 9, -5, 3, 1, 2, 10, 1},
		[]int{1555200, -518400, 194400, -129600, -777600, -172800, 311040, -518400, -1555200, -777600, -155520, -1555200},
	},
}

func TestProductExceptSelf(t *testing.T) {
	for i, test := range tests {
		t.Log("test ", i+1)

		got := productExceptSelf(test.nums)
		assert.Equal(t, test.expectedNums, got)

		t.Log("got: ", got, ", expected: ", test.expectedNums)
	}
}

func BenchmarkProductExceptSelf(b *testing.B) {
	funcs := []struct {
		name string
		f    func(nums []int) []int
	}{
		{"mine", productExceptSelf},
		{"leet-fastest", productExceptSelf_leet12ms},
	}

	var res []int
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
