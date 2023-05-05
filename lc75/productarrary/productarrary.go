package productarrary

func productExceptSelf(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}

	left := make([]int, l)
	right := make([]int, l)
	res := make([]int, l)

	i, j := 0, l-1
	left[i], right[j] = 1, 1
	for ; i < l && j >= 0; i, j = i+1, j-1 {
		if i > 0 {
			left[i] = nums[i-1] * left[i-1]
		}
		if j < l-1 {
			right[j] = nums[j+1] * right[j+1]
		}
	}

	for k := 0; k < l; k++ {
		res[k] = left[k] * right[k]
	}

	return res
}

func productExceptSelf_leet12ms(nums []int) []int {
	pref, suff, res := 1, 1, make([]int, len(nums))

	for i := 0; i < len(res); i++ {
		res[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		res[i] *= pref
		pref *= nums[i]

		res[len(nums)-i-1] *= suff
		suff *= nums[len(nums)-i-1]
	}

	return res
}
