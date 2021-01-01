package threeSumClosest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	var cdiff int
	sort.Ints(nums)

	l := len(nums)

	if l < 3 {
		return -999
	}

	csum := nums[0] + nums[1] + nums[2]
	if csum-target > 0 {
		cdiff = csum - target
	} else {
		cdiff = target - csum
	}

	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := l - 1
		for j < k {
			diff := 0
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				//fmt.Println(nums[i], nums[j], nums[k])
				return sum
			}

			if sum-target < 0 {
				diff = target - sum
				j++
			} else {
				diff = sum - target
				k--
			}

			if cdiff > diff {
				csum = sum
				cdiff = diff
			}
		}
	}

	return csum
}

func threeSumClosest0ms(nums []int, target int) int {
	sort.Ints(nums)
	res := math.MaxInt16
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right] + nums[i]
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				return target
			}
			absDiff := Abs(sum - target)
			asbRes := Abs(res - target)
			if absDiff < asbRes {
				res = sum
			}
		}
	}
	return res
}
func Abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
