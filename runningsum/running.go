package runningsum

func runningSum(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}

	sums := make([]int, l)

	sums[0] = nums[0]
	for i := 1; i < l; i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	return sums
}

func runningSumV2(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}

	for i := 1; i < l; i++ {
		nums[i] = nums[i-1] + nums[i]
	}

	return nums
}
