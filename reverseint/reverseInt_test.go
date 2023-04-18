package reverseint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result int

var tests = []struct {
	input    int
	expected int
}{
	{
		123,
		321,
	},
	{
		-123,
		-321,
	},
	{
		-2147483648,
		0,
	},
	{
		2147483647,
		0,
	},
	{
		2147478412,
		0,
	},
	{
		2147447412,
		2147447412,
	},
	{
		1563847412,
		0,
	},
	{
		-1563847412,
		0,
	},
}

func TestPivotIndex(t *testing.T) {
	for i, test := range tests {
		t.Log("test ", i+1)

		got := reverse(test.input)
		assert.Equal(t, test.expected, got)

		t.Log("got: ", got, ", expected: ", test.expected)

	}
}

func BenchmarkReverse(b *testing.B) {
	funcs := []struct {
		name string
		f    func(n int) int
	}{
		{"mine", reverse},
		{"leetcode-0ms", reverse_leet0ms},
	}

	var res int
	for _, fun := range funcs {

		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(tests[k%2].input)
			}

		})
	}

	result = res
}
