package threeSum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result [][]int

func Test3sum1(t *testing.T) {
	input := []int{-1, 0, 1, 2, -1, -4}

	want := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	got := threeSumV2(input)
	assert.Equal(t, want, got)
}

func Test3sum2(t *testing.T) {
	input := []int{}

	want := [][]int{}
	got := threeSumV2(input)
	assert.Equal(t, want, got)
}

func Test3sum3(t *testing.T) {
	input := []int{0}

	want := [][]int{}
	got := threeSumV2(input)
	assert.Equal(t, want, got)
}

func Test3sum4(t *testing.T) {
	input := []int{0, 0, 0, 0}

	want := [][]int{{0, 0, 0}}
	got := threeSumV2(input)
	assert.Equal(t, want, got)
}

func Test3sum5(t *testing.T) {
	input := []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}

	want := [][]int{{-4, 0, 4}, {-4, 1, 3}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1}}
	got := threeSumV2(input)
	assert.Equal(t, want, got)
}

func Test3sumLarge0Input(t *testing.T) {
	input := make([]int, 10000)

	want := [][]int{{0, 0, 0}}
	got := threeSumV2(input)
	assert.Equal(t, want, got)
}

func Benchmark3sum(b *testing.B) {
	input := make([]int, 10000)
	funcs := []struct {
		name string
		f    func(nums []int) [][]int
	}{
		{"my3sum", threeSum},
		{"my3sumV2", threeSumV2},
		{"3sumLeetCode-24ms", threeSum24ms},
	}

	var res [][]int
	for _, fun := range funcs {

		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(input)
			}

		})

		//fmt.Println(result, arr[result[0]], arr[result[1]])
	}

	result = res
}
