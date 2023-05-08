package movezeros

func moveZeroes(nums []int) {
	l := len(nums)
	if l == 1 {
		return
	}
	i0, jn0 := 0, 0
	for i0 < l && jn0 < l {
		if nums[i0] != 0 {
			i0++
			continue
		}
		if jn0 < i0 {
			jn0 = i0
		}
		if nums[jn0] == 0 {
			jn0++
			continue
		}

		nums[i0], nums[jn0] = nums[jn0], nums[i0]
	}
}

func moveZeroes_leet7ms(nums []int) {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
		} else if count > 0 {
			nums[i-count] = nums[i]
			nums[i] = 0
		}
	}
}
