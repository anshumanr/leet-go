package pivotindex

func pivotIndex(nums []int) int {
	nums = runningSum(nums)

	l := len(nums)
	var (
		ls int
		rs int
	)
	for i, val := range nums {
		if i == 0 {
			ls = 0
		} else {
			ls = nums[i-1]
		}
		rs = nums[l-1] - val

		if ls == rs {
			return i
		}
	}

	return -1
}

func runningSum(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}

	for i := 1; i < l; i++ {
		nums[i] = nums[i-1] + nums[i]
	}

	return nums
}

func pivotIndexV2(nums []int) int {
	preSums := make([]int, len(nums))
	for i, v := range nums {
		preSums[i] = v
		if i > 0 {
			preSums[i] += preSums[i-1]
		}
	}

	lastIdx := len(nums) - 1
	for i := range nums {
		sumLeft := 0
		if i > 0 {
			sumLeft = preSums[i-1]
		}
		sumRight := preSums[lastIdx] - preSums[i]
		if sumLeft == sumRight {
			return i
		}
	}
	return -1
}
