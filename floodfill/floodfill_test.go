package floodfill

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result [][]int

var tests = []struct {
	image         [][]int
	sr, sc, color int
	expected      [][]int
}{
	{
		[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
		1, 1, 2,
		[][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
	},
	{
		[][]int{{0, 0, 0}, {0, 0, 0}},
		0, 0, 0,
		[][]int{{0, 0, 0}, {0, 0, 0}},
	},
	{
		[][]int{{0, 0, 0}, {0, 0, 0}},
		1, 0, 2,
		[][]int{{2, 2, 2}, {2, 2, 2}},
	},
	{
		[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
		1, 2, 2,
		[][]int{{1, 1, 1}, {1, 1, 2}, {1, 0, 1}},
	},
}

func TestFloodfill(t *testing.T) {
	for i, test := range tests {
		t.Log("test ", i+1)

		got := floodFill(test.image, test.sr, test.sc, test.color)
		assert.Equal(t, len(test.image), len(got))

		t.Log("got: ", got, ", expected: ", test.expected)

		for i, val := range test.expected {
			for k, v := range val {
				assert.Equal(t, v, got[i][k])
			}
		}
	}
}

func BenchmarkFloodfill(b *testing.B) {
	funcs := []struct {
		name string
		f    func([][]int, int, int, int) [][]int
	}{
		{"mine", floodFill},
		{"minev2", floodFillv2},
		{"leet-fastest", floodFill_leet0ms},
	}

	var res [][]int
	for _, fun := range funcs {

		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(tests[k%3].image, tests[k%3].sr, tests[k%3].sc, tests[k%3].color)
			}

		})
	}

	result = res
}

func floodFill_leet0ms(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}
	dfs(image, sr, sc, image[sr][sc], color)
	return image
}

func dfs(image [][]int, i int, j int, startColor, newColor int) {
	if i < 0 || i >= len(image) || j < 0 || j >= len(image[0]) {
		return
	}

	if image[i][j] != startColor {
		return
	}

	image[i][j] = newColor
	dfs(image, i+1, j, startColor, newColor)
	dfs(image, i-1, j, startColor, newColor)
	dfs(image, i, j+1, startColor, newColor)
	dfs(image, i, j-1, startColor, newColor)
}
