package placeflowers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result bool

var tests = []struct {
	nums     []int
	limit    int
	expected bool
}{
	{
		[]int{1, 0, 0, 0, 1},
		1,
		true,
	},
	{
		[]int{0, 0, 1, 0, 1},
		1,
		true,
	},
	{
		[]int{1, 0, 0, 0, 1},
		2,
		false,
	},
	{
		[]int{1, 0, 0, 1, 0},
		1,
		false,
	},
	{
		[]int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0},
		5,
		true,
	},
}

func TestCanPlaceFlowers(t *testing.T) {
	for i, v := range tests {
		t.Log("test: ", i+1)
		got := canPlaceFlowers(v.nums, v.limit)
		assert.Equal(t, v.expected, got)
	}
}

func BenchmarkCanPlaceFlowers(b *testing.B) {
	funcs := []struct {
		name string
		f    func([]int, int) bool
	}{
		{"mine", canPlaceFlowers},
		{"leet-fastest", canPlaceFlowers_leet2ms},
	}

	var res bool
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
