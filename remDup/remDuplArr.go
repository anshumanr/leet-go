package remDup

func removeDuplicates(nums []int) int {
	curIndex := 0

	for i := 1; i < len(nums); i++ {

		if nums[i] != nums[curIndex] {
			curIndex++
			if (i - curIndex) > 0 {
				nums[curIndex] = nums[i]
			}
		}
	}

	return curIndex + 1
}

func removeDuplicates0ms(nums []int) int {
	result := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[result] = nums[i]
			result++
		}
	}
	return result
}
