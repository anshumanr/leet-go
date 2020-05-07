package addtwonum

import (
	"fmt"
	"testing"
)

var result *ListNode

var tests = []struct {
	l1, l2, want []int
}{
	{
		[]int{3},
		[]int{4},
		[]int{7},
	},
	{
		[]int{6},
		[]int{4},
		[]int{0, 1},
	},
	{
		[]int{9, 8},
		[]int{1},
		[]int{0, 9},
	},
	{
		[]int{9, 9},
		[]int{1},
		[]int{0, 0, 1},
	},
	{
		[]int{9, 9, 9, 9, 9, 9},
		[]int{1},
		[]int{0, 0, 0, 0, 0, 0, 1},
	},
	{
		[]int{9, 8, 9, 9, 9, 9},
		[]int{1, 1},
		[]int{0, 0, 0, 0, 0, 0, 1},
	},
	{
		[]int{2, 4, 3},
		[]int{5, 6, 4},
		[]int{7, 0, 8},
	},
	{
		[]int{0, 4, 3},
		[]int{5, 6, 0, 0},
		[]int{5, 0, 4},
	},
	{
		[]int{0, 4, 3},
		[]int{5, 6},
		[]int{5, 0, 4},
	},
	{
		[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		[]int{0, 0, 0, 0, 0},
		[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	},
	{
		[]int{0, 4, 3, 8, 3, 4, 1, 3, 6},
		[]int{5, 6},
		[]int{5, 0, 4, 8, 3, 4, 1, 3, 6},
	},
	{
		[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 9, 9, 9, 9, 9, 9, 9, 1},
		[]int{5, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		[]int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
	},
	{
		[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		[]int{5, 6, 4},
		[]int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	},
	{
		[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{5, 6, 4},
		[]int{6, 6, 4},
	},
	{
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{5, 6, 4},
		[]int{5, 6, 4},
	},
	{
		[]int{5, 6, 4},
		[]int{0, 0, 9, 9, 9, 9, 9, 9, 9, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1},
		[]int{5, 6, 3, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1},
	},
	{
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{},
		[]int{0},
	},
	{
		[]int{0},
		[]int{0},
		[]int{0},
	},
}

func TestAddTwoNumbersV1(t *testing.T) {
	for _, c := range tests {

		l1 := arrToList(c.l1)
		l2 := arrToList(c.l2)

		l := addTwoNumbersV1(l1, l2)

		sum := listToArr(l)

		fmt.Println(*sum, c.want)
	}
}

func TestAddTwoNumbersV2(t *testing.T) {
	for _, c := range tests {

		l1 := arrToList(c.l1)
		l2 := arrToList(c.l2)

		l := addTwoNumbersV2(l1, l2)

		sum := listToArr(l)

		fmt.Println("got:", *sum, "want:", c.want)
	}
}

func TestAddTwoNumbersV3(t *testing.T) {
	for _, c := range tests {

		l1 := arrToList(c.l1)
		l2 := arrToList(c.l2)

		l := addTwoNumbersV3(l1, l2)

		sum := listToArr(l)

		fmt.Println("got:", *sum, "want:", c.want)
	}
}

func TestAddTwoNumbers4ms(t *testing.T) {
	for _, c := range tests {

		l1 := arrToList(c.l1)
		l2 := arrToList(c.l2)

		l := addTwoNumbers4ms(l1, l2)

		sum := listToArr(l)

		fmt.Println("got:", *sum, "want:", c.want)
	}
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	funcs := []struct {
		name string
		f    func(*ListNode, *ListNode) *ListNode
	}{
		{"addTwoNumbersV1", addTwoNumbersV1},
		{"addTwoNumbersV2", addTwoNumbersV2},
		{"addTwoNumbersV3", addTwoNumbersV3},
		{"addTwoNumbers-0ms", addTwoNumbers0ms},
		{"addTwoNumbers-4ms", addTwoNumbers4ms},
	}

	//n := len(tests)
	var res *ListNode
	l1 := arrToList(tests[12].l1)
	l2 := arrToList(tests[12].l2)

	for _, fun := range funcs {
		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(l1, l2)
			}

		})
	}

	result = res
}
