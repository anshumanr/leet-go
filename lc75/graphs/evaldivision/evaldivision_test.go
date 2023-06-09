package evaldivision

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result []float64

var tests = []struct {
	input    [][]string
	values   []float64
	queries  [][]string
	expected []float64
}{
	{
		[][]string{{"a", "b"}, {"b", "c"}},
		[]float64{2.0, 3.0},
		[][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
		[]float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000},
	},
	{
		[][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
		[]float64{1.5, 2.5, 5.0},
		[][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
		[]float64{3.75000, 0.40000, 5.00000, 0.20000},
	},
	{
		[][]string{{"a", "b"}},
		[]float64{0.5},
		[][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}},
		[]float64{0.50000, 2.00000, -1.00000, -1.00000},
	},
	{
		[][]string{{"x1", "x2"}, {"x2", "x3"}, {"x3", "x4"}, {"x4", "x5"}},
		[]float64{3.0, 4.0, 5.0, 6.0},
		[][]string{{"x1", "x5"}, {"x5", "x2"}, {"x2", "x4"}, {"x2", "x2"}, {"x2", "x9"}, {"x9", "x9"}},
		[]float64{360.00000, 0.008333333333333333, 20.00000, 1.00000, -1.00000, -1.00000},
	},
	{
		[][]string{{"x1", "x2"}, {"x2", "x3"}, {"x1", "x4"}, {"x2", "x5"}},
		[]float64{3.0, 0.5, 3.4, 5.6},
		[][]string{{"x2", "x4"}, {"x1", "x5"}, {"x1", "x3"}, {"x5", "x5"}, {"x5", "x1"}, {"x3", "x4"}, {"x4", "x3"}, {"x6", "x6"}, {"x0", "x0"}},
		[]float64{1.1333333333333333, 16.80000, 1.50000, 1.00000, 0.05952380952380952, 2.2666666666666666, 0.4411764705882353, -1.00000, -1.00000},
	},
}

func TestCalcEquation(t *testing.T) {
	for i, test := range tests {
		t.Log("test ", i+1)

		got := calcEquationV2(test.input, test.values, test.queries)
		assert.Equal(t, test.expected, got)

		t.Log("got: ", got, ", expected: ", test.expected)
	}
}

func BenchmarkCalcEquation(b *testing.B) {
	funcs := []struct {
		name string
		f    func([][]string, []float64, [][]string) []float64
	}{
		{"mine", calcEquation},
		{"mine-v2", calcEquationV2},
		{"leet-fastest", calcEquation_leet0ms},
	}

	var res []float64
	for _, fun := range funcs {

		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(tests[k%len(tests)].input, tests[k%len(tests)].values, tests[k%len(tests)].queries)
			}

		})
	}

	result = res
}
