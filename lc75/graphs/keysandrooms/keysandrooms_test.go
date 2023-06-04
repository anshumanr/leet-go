package keysandrooms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result bool

var tests = []struct {
	nums     [][]int
	expected bool
}{
	{
		[][]int{{1}, {2}, {3}, {}},
		true,
	},
	{
		[][]int{{1, 3}, {3, 0, 1}, {2}, {0}},
		false,
	},
}

func TestCanVisitAllRooms(t *testing.T) {
	for i, test := range tests {
		t.Log("test ", i+1)

		got := canVisitAllRooms(test.nums)
		assert.Equal(t, test.expected, got)

		t.Log("got: ", got, ", expected: ", test.expected)
	}
}

func BenchmarkCanVisitAllRooms(b *testing.B) {
	funcs := []struct {
		name string
		f    func(nums [][]int) bool
	}{
		{"mine", canVisitAllRooms},
		{"leet-fastest", canVisitAllRooms_leet0ms},
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
