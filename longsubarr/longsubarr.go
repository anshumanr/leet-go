package longsubarr

import (
	"container/list"
)

func longestSubarray(nums []int, limit int) int {
	res := 1
	n := len(nums)
	mn := list.New()
	mx := list.New()
	mn.PushBack(0)
	mx.PushBack(0)
	l := 0
	r := 1
	for r < n {
		cur := nums[r]
		for mn.Len() > 0 && nums[mn.Back().Value.(int)] >= cur {
			mn.Remove(mn.Back())
		}
		for mx.Len() > 0 && nums[mx.Back().Value.(int)] <= cur {
			mx.Remove(mx.Back())
		}
		mn.PushBack(r)
		mx.PushBack(r)

		for nums[mx.Front().Value.(int)]-nums[mn.Front().Value.(int)] > limit {
			l++
			for mx.Front().Value.(int) < l {
				mx.Remove(mx.Front())
			}
			for mn.Front().Value.(int) < l {
				mn.Remove(mn.Front())
			}
		}

		res = max(res, r-l+1)
		r++
	}
	return res
}

func longestSubarrayV1(nums []int, limit int) int {
	if len(nums) > 100000 {
		return 0
	}
	max := 0

	head := 0
	for head < len(nums) {
		i := head
	loop:
		for ; i < len(nums); i++ {
			subarr := nums[head : i+1]
			//fmt.Println(subarr, head)

			tmax := true
			for k := 0; k < len(subarr); k++ {
				for l := 0; l < len(subarr); l++ {
					if abs(subarr[l]-subarr[k]) > limit {
						tmax = false
						break loop
					}
				}
			}

			if tmax && len(subarr) > max {
				max = len(subarr)
			}
		}
		if i == len(nums) {
			break
		}

		head++
	}

	return max
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(res, n int) int {
	if n > res {
		return n
	}

	return res
}

func min(res, n int) int {
	if n < res {
		return n
	}

	return res
}

func longestSubarrayV3(nums []int, limit int) int {
	if len(nums) > 100000 {
		return 0
	}
	res := 1

	left := 0
	right := 1
	lmax := left
	lmin := left

	for right < len(nums) {
		cur := nums[right]

		for abs(nums[lmax]-cur) > limit {
			if nums[lmax] >= cur {
				left = lmax + 1
			} else {
				left = right
			}

			lmax = left
			lmin = left
			for i := left; i <= right; i++ {
				if nums[i] >= nums[lmax] {
					lmax = i
				}
				if nums[i] <= nums[lmin] {
					lmin = i
				}
			}
		}

		for abs(nums[lmin]-cur) > limit {
			if nums[lmin] >= cur {
				left = right
			} else {
				left = lmin + 1
			}

			lmin = left
			lmax = left
			for i := left; i <= right; i++ {
				if nums[i] >= nums[lmax] {
					lmax = i
				}
				if nums[i] <= nums[lmin] {
					lmin = i
				}
			}
		}

		if cur >= nums[lmax] {
			lmax = right
		}
		if cur <= nums[lmin] {
			lmin = right
		}

		res = max(res, right-left+1)
		right++
	}

	return res
}

func longestSubarrayV2(nums []int, limit int) int {
	if len(nums) > 100000 {
		return 0
	}
	res := 0

	left := 0
	right := 1
	lmax := left
	lmin := left
	for right < len(nums) {

		//fmt.Println(subarr, limit, left, right)

		for k := left; k <= right; k++ {
			if abs(nums[lmax]-nums[k]) > limit || abs(nums[lmin]-nums[k]) > limit {
				left = min(left+1, k)
				lmax = left
				lmin = left
				break
			}

			if nums[k] >= nums[lmax] {
				lmax = k
			}
			if nums[k] <= nums[lmin] {
				lmin = k
			}

			//fmt.Println(lmax, lmin, k)
		}

		res = max(res, right-left+1)
		right++
	}

	return res
}
