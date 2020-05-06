package two-sum

import (
	"math/rand"
	"testing"
)

var result []int

func BenchmarkTwoSum(b *testing.B) {
	n := 50
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(n)
	}

	funcs := []struct {
		name string
		f    func([]int, int) []int
	}{
		{"twoSum-0ms", twoSumMapInt0ms},
		{"twoSum-4ms", twoSumMapInt4ms},
		{"twoSum-8ms", twoSumMapInt8ms},
		{"twoSum-12ms", twoSumMapInt12ms},
		{"twoSum-20ms", twoSumMapInt20ms},
		{"twoSum-32ms", twoSumMapInt32ms},
		{"twoSumAttempt1", twoSumAttempt1},
		{"twoSumAttempt2", twoSumAttempt2},
		{"twoSumAttempt3", twoSumAttempt3},
	}

	var res []int
	for _, fun := range funcs {

		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(arr, 47)
			}

		})

		//fmt.Println(result, arr[result[0]], arr[result[1]])
	}

	result = res
}

// func BenchmarkTwoSumS(b *testing.B) {
// 	n := 100
// 	arr := make([]int, n)

// 	for i := 0; i < n; i++ {
// 		arr[i] = rand.Intn(100)
// 	}

// 	var res []int
// 	for k := 0; k < b.N; k++ {
// 		res = twoSumS(arr, 6)
// 	}

// 	result := res
// 	fmt.Println(result, arr[result[0]], arr[result[1]])
// }
