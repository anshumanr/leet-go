package increasingtriplet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result bool

var tests = []struct {
	nums         []int
	expectedNums bool
}{
	{
		[]int{1, 2},
		false,
	},
	{
		[]int{1, 2, 2, 1, 1, 3},
		true,
	},
	{
		[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
		true,
	},
	{
		[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0, -3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
		true,
	},
}

func TestProductExceptSelf(t *testing.T) {
	for i, test := range tests {
		t.Log("test ", i+1)

		got := uniqueOccurrences(test.nums)
		assert.Equal(t, test.expectedNums, got)

		t.Log("got: ", got, ", expected: ", test.expectedNums)
	}
}

func BenchmarkProductExceptSelf(b *testing.B) {
	funcs := []struct {
		name string
		f    func(nums []int) bool
	}{
		{"mine", uniqueOccurrences},
		{"leet-fastest", uniqueOccurrences_leet0ms},
	}

	var res bool
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
