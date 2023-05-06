package increasingtriplet

import "math"

func increasingTriplet(nums []int) bool {
	l := len(nums)
	if l < 3 {
		return false
	}

	min1, min2, min3 := nums[0], (1 << 31), (1 << 31)
	mini, minj, mink := 0, -1, -1
	previ := l

	i := 1
	for i < l {
		if nums[i] < min1 {
			if min2 == 1<<31 {
				min1 = nums[i]
				mini = i
			} else {
				if previ == l {
					previ = i
				}
			}
		} else if nums[i] < min2 && nums[i] > min1 {
			min2 = nums[i]
			minj = i
		} else if nums[i] < min3 && nums[i] > min2 {
			return true
		}

		if i == l-1 && previ < l {
			i = previ
			previ = l
			min1, min2, min3 = nums[i], (1 << 31), (1 << 31)
			mini, minj, mink = i, -1, -1
			previ = l
		}
		i++
	}

	return mini < minj && minj < mink
}

func increasingTriplet_leet(nums []int) bool {
	max := math.MaxInt
	min := math.MaxInt

	if len(nums) < 3 {
		return false
	}

	for _, num := range nums {
		if min >= num {
			min = num
			continue
		} else if max >= num {
			max = num
			continue
		}
		return true
	}
	return false
}
