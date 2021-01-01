package threeSumClosest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result int

func Test3sum1(t *testing.T) {
	input := []int{1, 2, -1, -4}

	target := 1
	want := 2
	got := threeSumClosest(input, target)
	assert.Equal(t, want, got)

	target = -5
	want = -4
	got = threeSumClosest(input, target)
	assert.Equal(t, want, got)
}

func Test3sum2(t *testing.T) {
	input := []int{-1, 0, 2, -1, -4}

	target := -5
	want := -5
	got := threeSumClosest(input, target)
	assert.Equal(t, want, got)
}

func Test3sumLarge0Input(t *testing.T) {
	input := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		if i < 500 {
			input[i] = i % 1000
		} else {
			input[i] = -i % 1000
		}
	}

	target := 399
	want := 399
	got := threeSumClosest(input, target)
	assert.Equal(t, want, got)
}

func Benchmark3sum(b *testing.B) {
	input := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		if i < 500 {
			input[i] = i % 1000
		} else {
			input[i] = -i % 1000
		}
	}

	target := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		if i < 500 {
			target[i] = i % 1000
		} else {
			target[i] = -i % 1000
		}
	}

	funcs := []struct {
		name string
		f    func(nums []int, target int) int
	}{
		{"my3sumClosest", threeSumClosest},
		{"threeSumClosest0ms", threeSumClosest0ms},
	}

	var res int
	for _, fun := range funcs {

		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(input, target[k%1000])
			}

		})
	}

	result = res
}
