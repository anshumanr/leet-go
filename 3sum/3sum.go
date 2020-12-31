package threeSum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	l := len(nums)

	if l == 0 || nums[0] > 0 || nums[l-1] < 0 {
		return [][]int{}
	}

	for i := 0; i < l-2; i++ {
		j := i + 1
		k := l - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				z := []int{nums[i], nums[j], nums[k]}
				if isUnique(res, z) {
					res = append(res, z)
				}
				j++
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}

	return res
}

func isUnique(sum [][]int, c []int) bool {
	l := len(sum)

	for i := 0; i < l; i++ {
		if isEqual(sum[i], c) {
			return false
		}
	}

	return true
}

func isEqual(a, b []int) bool {
	l1 := len(a)
	l2 := len(b)

	if l1 != l2 {
		return false
	}

	for i := 0; i < l1; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func threeSumV2(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	l := len(nums)

	if l == 0 || nums[0] > 0 || nums[l-1] < 0 {
		return [][]int{}
	}

	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := l - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++

				for j < k && (nums[j] == nums[j-1]) {
					j++
				}
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}

	return res
}

func threeSum24ms(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if i >= len(nums)-2 {
			break
		}

		l, r := i+1, len(nums)-1

		for l < r {
			threeSum := nums[i] + nums[l] + nums[r]
			if threeSum < 0 {
				l++
			} else if threeSum > 0 {
				r--
			} else {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}
	return result
}
