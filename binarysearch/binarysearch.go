package binarysearch

func search(nums []int, target int) int {
	lo, high := 0, len(nums)-1
	mid := (lo + high) / 2

	for lo <= high {
		switch {
		case high-lo <= 1:
			switch {
			case target == nums[lo]:
				return lo
			case target == nums[high]:
				return high
			default:
				return -1
			}
		case (high - lo) <= 1:
		case target == nums[mid]:
			return mid
		case target > nums[mid]:
			lo = mid
			mid = (lo + high) / 2
		case target < nums[mid]:
			high = mid
			mid = (lo + high) / 2
		}
	}

	return -1
}

func searchv2(nums []int, target int) int {
	lo, high := 0, len(nums)-1
	mid := (lo + high) / 2

	for lo <= high {
		if target == nums[mid] {
			return mid
		}
		if target > nums[mid] {
			lo = mid + 1
			mid = (lo + high) / 2
		} else {
			high = mid - 1
			mid = (lo + high) / 2
		}
	}

	return -1
}

func search_leet17ms(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
