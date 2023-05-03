package kidscandies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result []bool

var tests = []struct {
	nums     []int
	limit    int
	expected []bool
}{
	{
		[]int{2, 3, 5, 1, 3},
		3,
		[]bool{true, true, true, false, true},
	},
	{
		[]int{4, 2, 1, 1, 2},
		1,
		[]bool{true, false, false, false, false},
	},
	{
		[]int{12, 1, 12},
		10,
		[]bool{true, false, true},
	},
}

func TestKidsWithCandies(t *testing.T) {
	for i, v := range tests {
		t.Log("test: ", i+1)
		got := kidsWithCandies(v.nums, v.limit)
		assert.Equal(t, v.expected, got)
	}
}

func BenchmarkKidsWithCandies(b *testing.B) {
	funcs := []struct {
		name string
		f    func([]int, int) []bool
	}{
		{"mine", kidsWithCandies},
		{"leet-fastest", kidsWithCandies_leet1ms},
	}

	var res []bool
	for _, fun := range funcs {

		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(tests[k%len(tests)].nums, tests[k%len(tests)].limit)
			}

		})
	}

	result = res
}
